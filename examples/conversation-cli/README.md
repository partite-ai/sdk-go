# conversation-cli

A small command-line example for the Partite AI Go SDK. It:

1. starts a conversation on a version of a mesh by sending a first message to an intent, optionally uploading files with it,
2. streams that request's events (thinking, messages, the final response) as they arrive,
3. sends a follow-up message in the same conversation and streams its request's events too.

## Run

From the root of this repository:

```sh
export PARTITE_API_KEY=...
# Optional: point at a different server (defaults to https://api.partite.ai)
# export PARTITE_BASE_URL=http://localhost:8080

go run ./examples/conversation-cli <mesh-id>@<version> <intent> '<input-json>' ['<follow-up-input-json>'] [--attach <path>]...
```

For example:

```sh
go run ./examples/conversation-cli my-mesh-id@prod ask '{"question":"hello"}' '{"question":"and again"}'
```

`<version>` picks the mesh version to run on: a number is a version number (`my-mesh@3`),
`draft` is the mesh's draft (`my-mesh@draft`), and anything else is a version label
(`my-mesh@prod`). A mesh ID without a version is rejected, as the API never picks the draft by
default.

The input arguments must be JSON objects that match the intent's input schema.
The follow-up input defaults to the first input if omitted, unless `--attach` is used
(see below).

## Attachments

`--attach <path>` uploads a file with the first request. Repeat it to attach several
files; it may appear anywhere on the command line. Each file is sent as a separate
`attachments` part, named after the file's base name and with a Content-Type derived
from its extension (`application/octet-stream` if unknown). The follow-up request has
no attachments.

The input message refers to an attachment as `attachment://<filename>`:

```sh
go run ./examples/conversation-cli my-mesh-id@prod ask \
  '{"question":"summarize attachment://report.pdf"}' \
  --attach ./report.pdf
```

The API rejects an input message that references an attachment that wasn't uploaded.

`attachment://` references only resolve against files uploaded in the same request, and
the follow-up request carries no attachments. So when `--attach` is given,
`<follow-up-input-json>` is required: the first input can't be reused for the follow-up.

## Output and exit codes

Each event is printed on its own line, e.g. `[thinking] ...`, `[message] ...`,
`[response] {...}`. The example stops reading a request's event stream as soon as it
sees a terminal event (`response`, `error`, `cancelled`, `server_error` or `task`).

| Exit code | Meaning |
|-----------|---------|
| 0 | Both requests finished with a response |
| 1 | An API call failed, or a request ended with an error, cancellation or a task to complete |
| 2 | Bad arguments (wrong count, a mesh without a version, invalid JSON, an `--attach` path that isn't a readable file, or `--attach` without a follow-up input) |
