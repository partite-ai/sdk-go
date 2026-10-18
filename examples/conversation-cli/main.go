// Command conversation-cli starts a conversation with a Partite AI mesh,
// streams the events of the first request, then sends a follow-up request in
// the same conversation and streams its events too.
//
// Usage:
//
//	conversation-cli <mesh-id>@<version> <intent> <input-json> [<follow-up-input-json>] [--attach <path>]...
//
// <version> picks the mesh version: a number is a version number, "draft" is
// the mesh's draft, and anything else is a version label, e.g. my-mesh@3,
// my-mesh@draft or my-mesh@prod.
//
// Files given with --attach are uploaded with the first request. The input
// message refers to an attachment by its file name using the
// attachment://<filename> syntax, e.g. {"question":"summarize attachment://report.pdf"}.
//
// The API key is read from PARTITE_API_KEY.
// Set PARTITE_BASE_URL to talk to a server other than https://api.partite.ai.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	partiteai "github.com/partite-ai/sdk-go"
	"github.com/partite-ai/sdk-go/client"
	"github.com/partite-ai/sdk-go/option"
)

const usage = `usage: conversation-cli <mesh-id>@<version> <intent> <input-json> [<follow-up-input-json>] [--attach <path>]...

  <version> is a version number, draft, or a version label, e.g. my-mesh@3,
  my-mesh@draft or my-mesh@prod.
  <input-json> and <follow-up-input-json> must be JSON objects, e.g. '{"question":"hello"}'.
  The follow-up defaults to <input-json>, except when --attach is used: then it is
  required, because attachment:// references only resolve against files uploaded
  in the same request.

  --attach <path>  Upload a file with the first request (repeatable, allowed anywhere).
                   Refer to it in the input as attachment://<filename>, e.g.
                   '{"question":"summarize attachment://report.pdf"}'.`

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(argv []string) int {
	// 1. Parse the command line. The standard flag package stops at the first
	// positional argument, so --attach is picked out by hand to allow it anywhere.
	args, attachPaths, err := splitArgs(argv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n%s\n", err, usage)
		return 2
	}
	if len(args) < 3 || len(args) > 4 {
		fmt.Fprintln(os.Stderr, usage)
		return 2
	}
	mesh := parseMesh(args[0])
	if mesh == nil {
		fmt.Fprintf(os.Stderr, "invalid mesh: %s (expected <mesh-id>@<version>, e.g. my-mesh@3, my-mesh@draft or my-mesh@prod)\n\n%s\n", args[0], usage)
		return 2
	}
	intent := args[1]

	input, err := parseJSONObject(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid <input-json>: %v\n\n%s\n", err, usage)
		return 2
	}
	// attachment://<filename> references only resolve against files uploaded in
	// the same request, and the follow-up carries no attachments, so the first
	// input can't be reused for it when files are attached.
	if len(attachPaths) > 0 && len(args) < 4 {
		fmt.Fprintf(os.Stderr, "<follow-up-input-json> is required with --attach: attachment:// references only resolve against files in the same request, so <input-json> can't be reused for the follow-up\n\n%s\n", usage)
		return 2
	}
	followUp := input
	if len(args) == 4 {
		followUp, err = parseJSONObject(args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid <follow-up-input-json>: %v\n\n%s\n", err, usage)
			return 2
		}
	}

	// Open each attachment. The SDK uploads any io.Reader; wrapping it in a
	// FileParam lets us choose the part's filename and Content-Type. The
	// filename is what the input message refers to as attachment://<filename>.
	var attachments []io.Reader
	for _, path := range attachPaths {
		file, size, err := openAttachment(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid --attach %q: %v\n\n%s\n", path, err, usage)
			return 2
		}
		defer file.Close()

		name := filepath.Base(path)
		contentType := mime.TypeByExtension(filepath.Ext(name))
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		fmt.Printf("attachment: %s (%d bytes)\n", name, size)
		attachments = append(attachments, partiteai.NewFileParam(file, name, contentType))
	}

	// 2. Create the client. With no explicit key, NewClient reads
	// PARTITE_API_KEY from the environment and sends it as a bearer token;
	// you can also pass it with option.WithAPIKey.
	var opts []option.RequestOption
	if baseURL := os.Getenv("PARTITE_BASE_URL"); baseURL != "" {
		opts = append(opts, option.WithBaseURL(baseURL))
	}
	c := client.NewClient(opts...)
	ctx := context.Background()

	// 3. Create a conversation on the mesh version, starting it with a first
	// request, and upload the attachments (one multipart part per file).
	created, err := c.Conversations.CreateConversation(ctx, &partiteai.CreateConversationRequest{
		Params: &partiteai.CreateConversationParams{
			Mesh:   mesh,
			Intent: partiteai.String(intent),
			Input:  input,
		},
		Attachments: attachments,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: creating conversation: %v\n", err)
		return 1
	}
	fmt.Printf("conversation: %s\n", created.ID)
	// A conversation created with an input always has the request that input started.
	if created.Request == nil {
		fmt.Fprintln(os.Stderr, "error: the conversation was created without a request")
		return 1
	}
	fmt.Printf("request: %s\n", created.Request.ID)

	// 4. Stream the first request's events until it finishes.
	ok, err := streamRequestEvents(ctx, c, created.ID, created.Request.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: streaming events: %v\n", err)
		return 1
	}
	// 5. Don't send a follow-up if the first request failed.
	if !ok {
		return 1
	}

	// 6. Send a follow-up message to the same conversation (no attachments),
	// which starts another request.
	next, err := c.Conversations.CreateRequest(ctx, &partiteai.CreateRequestRequest{
		ConversationID: created.ID,
		Params: &partiteai.CreateRequestParams{
			Intent: partiteai.String(intent),
			Input:  followUp,
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: creating follow-up request: %v\n", err)
		return 1
	}
	fmt.Printf("request: %s\n", next.ID)

	// 7. Stream the follow-up request's events.
	ok, err = streamRequestEvents(ctx, c, created.ID, next.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: streaming events: %v\n", err)
		return 1
	}
	if !ok {
		return 1
	}
	return 0
}

// streamRequestEvents prints each event of a request as it arrives. It stops
// at the first terminal event (response, error, cancelled, server_error or task)
// and closes the stream rather than waiting for the server to close it.
// It reports true if the request ended with a response.
func streamRequestEvents(ctx context.Context, c *client.Client, conversation, request string) (bool, error) {
	stream, err := c.Conversations.StreamRequestEvents(ctx, &partiteai.StreamRequestEventsRequest{
		ConversationID: conversation,
		RequestID:      request,
	})
	if err != nil {
		return false, err
	}
	defer stream.Close()

	for {
		event, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return false, errors.New("event stream ended before the request finished")
		}
		if err != nil {
			return false, err
		}

		// RequestStreamEvent is a union: Type says which field is populated.
		switch event.Type {
		case "started":
			// The first event of every stream; it names the conversation and the request.
		case "message":
			fmt.Printf("[message] %s\n", event.Message.Data)
		case "thinking":
			fmt.Printf("[thinking] %s\n", event.Thinking.Data)
		case "task_response":
			fmt.Printf("[task_response] %s\n", toJSON(event.TaskResponse.Data))
		case "response":
			fmt.Printf("[response] %s\n", toJSON(event.Response.Data))
			return true, nil
		case "error":
			fmt.Printf("[error] %s\n", event.Error.Data)
			return false, nil
		case "cancelled":
			fmt.Printf("[cancelled] %s\n", event.Cancelled.Data)
			return false, nil
		case "server_error":
			fmt.Printf("[server_error] %s\n", event.ServerError.Error)
			return false, nil
		case "task":
			fmt.Printf("[task] %s\n", toJSON(event.Task.Data))
			fmt.Println("request needs a task completed, which this example does not handle")
			return false, nil
		default:
			// Ignore event types this example doesn't know about.
		}
	}
}

// versionNumber matches a mesh version number.
var versionNumber = regexp.MustCompile(`^[0-9]+$`)

// parseMesh parses "<mesh-id>@<version>" into the mesh to start the
// conversation on, or returns nil. The API takes exactly one of a version
// number, a version label and the draft, so a mesh ID on its own is rejected
// rather than defaulting to the draft.
func parseMesh(s string) *partiteai.ConversationMesh {
	at := strings.LastIndex(s, "@")
	if at < 0 {
		return nil
	}
	id, version := s[:at], s[at+1:]
	if id == "" || version == "" {
		return nil
	}
	mesh := &partiteai.ConversationMesh{ID: id}
	switch {
	case versionNumber.MatchString(version):
		n, err := strconv.Atoi(version)
		if err != nil {
			return nil
		}
		mesh.Version = partiteai.Int(n)
	case version == "draft":
		mesh.Draft = partiteai.Bool(true)
	default:
		mesh.VersionLabel = partiteai.String(version)
	}
	return mesh
}

// splitArgs separates "--attach <path>" / "--attach=<path>" options from the
// positional arguments. They may appear anywhere on the command line.
func splitArgs(argv []string) (positional, attach []string, err error) {
	for i := 0; i < len(argv); i++ {
		arg := argv[i]
		switch {
		case arg == "--attach":
			if i+1 >= len(argv) {
				return nil, nil, errors.New("--attach needs a file path")
			}
			i++
			attach = append(attach, argv[i])
		case strings.HasPrefix(arg, "--attach="):
			attach = append(attach, strings.TrimPrefix(arg, "--attach="))
		case strings.HasPrefix(arg, "--"):
			return nil, nil, fmt.Errorf("unknown option %s", arg)
		default:
			positional = append(positional, arg)
		}
	}
	return positional, attach, nil
}

// openAttachment opens path for upload and returns the open file and its size.
// The caller must close the file.
func openAttachment(path string) (*os.File, int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, 0, err
	}
	if !info.Mode().IsRegular() {
		file.Close()
		return nil, 0, errors.New("not a regular file")
	}
	return file, info.Size(), nil
}

// parseJSONObject parses s and requires it to be a JSON object.
func parseJSONObject(s string) (map[string]any, error) {
	var obj map[string]any
	if err := json.Unmarshal([]byte(s), &obj); err != nil || obj == nil {
		return nil, errors.New("must be a JSON object")
	}
	return obj, nil
}

// toJSON renders v as compact JSON for printing.
func toJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}
