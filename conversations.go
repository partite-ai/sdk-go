// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/partite-ai/sdk-go/internal"
	io "io"
	big "math/big"
	time "time"
)

var (
	attachAssetsToConversationRequestFieldConversationID = big.NewInt(1 << 0)
)

type AttachAssetsToConversationRequest struct {
	// The ID of an existing conversation
	ConversationID string      `json:"-" url:"-"`
	Attachments    []io.Reader `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (a *AttachAssetsToConversationRequest) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AttachAssetsToConversationRequest) SetConversationID(conversationID string) {
	a.ConversationID = conversationID
	a.require(attachAssetsToConversationRequestFieldConversationID)
}

func (a *AttachAssetsToConversationRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler AttachAssetsToConversationRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*a = AttachAssetsToConversationRequest(body)
	return nil
}

func (a *AttachAssetsToConversationRequest) MarshalJSON() ([]byte, error) {
	type embed AttachAssetsToConversationRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	cancelRequestRequestFieldConversationID = big.NewInt(1 << 0)
	cancelRequestRequestFieldRequestID      = big.NewInt(1 << 1)
)

type CancelRequestRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`
	// The ID of an existing request
	RequestID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CancelRequestRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelRequestRequest) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(cancelRequestRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelRequestRequest) SetRequestID(requestID string) {
	c.RequestID = requestID
	c.require(cancelRequestRequestFieldRequestID)
}

var (
	completeTaskRequestFieldConversationID = big.NewInt(1 << 0)
	completeTaskRequestFieldRequestID      = big.NewInt(1 << 1)
	completeTaskRequestFieldTaskID         = big.NewInt(1 << 2)
)

type CompleteTaskRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`
	// The ID of an existing request
	RequestID string `json:"-" url:"-"`
	// The ID of the task to complete
	TaskID string         `json:"-" url:"-"`
	Body   map[string]any `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CompleteTaskRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CompleteTaskRequest) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(completeTaskRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CompleteTaskRequest) SetRequestID(requestID string) {
	c.RequestID = requestID
	c.require(completeTaskRequestFieldRequestID)
}

// SetTaskID sets the TaskID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CompleteTaskRequest) SetTaskID(taskID string) {
	c.TaskID = taskID
	c.require(completeTaskRequestFieldTaskID)
}

func (c *CompleteTaskRequest) UnmarshalJSON(data []byte) error {
	var body map[string]any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.Body = body
	return nil
}

func (c *CompleteTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

var (
	createConversationRequestRequestFieldConversationID = big.NewInt(1 << 0)
	createConversationRequestRequestFieldIntentName     = big.NewInt(1 << 1)
)

type CreateConversationRequestRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`
	// The name of the Intent defined on the target mesh which will be used for this request
	IntentName  string      `json:"-" url:"-"`
	Attachments []io.Reader `json:"-" url:"-"`
	// The input message to provide to the mesh. The structure of this object must conform to the input message type schema defined on the Intent. Within the input message, any asset references may use the syntax `attachment://<filename>` format to refer to assets attached in the same multipart/form-data body.
	InputMessage map[string]any `json:"input_message" url:"-"`
	// Request options
	RequestOptions *RequestCreateOptions `json:"request_options,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateConversationRequestRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationRequestRequest) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(createConversationRequestRequestFieldConversationID)
}

// SetIntentName sets the IntentName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationRequestRequest) SetIntentName(intentName string) {
	c.IntentName = intentName
	c.require(createConversationRequestRequestFieldIntentName)
}

func (c *CreateConversationRequestRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConversationRequestRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateConversationRequestRequest(body)
	return nil
}

func (c *CreateConversationRequestRequest) MarshalJSON() ([]byte, error) {
	type embed CreateConversationRequestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createConversationWithInitialRequestRequestFieldIntentName     = big.NewInt(1 << 0)
	createConversationWithInitialRequestRequestFieldMeshID         = big.NewInt(1 << 1)
	createConversationWithInitialRequestRequestFieldVersionOrdinal = big.NewInt(1 << 2)
	createConversationWithInitialRequestRequestFieldVersionLabel   = big.NewInt(1 << 3)
)

type CreateConversationWithInitialRequestRequest struct {
	// The name of the Intent defined on the target mesh which will be used to start the conversation
	IntentName string `json:"-" url:"-"`
	// The ID of the Mesh for which to create a conversation
	MeshID string `json:"-" url:"mesh_id"`
	// The version number of a mesh version for which to create the conversation
	VersionOrdinal *int `json:"-" url:"version_ordinal,omitempty"`
	// The version label of a mesh version for which to create the conversation
	VersionLabel *string     `json:"-" url:"version_label,omitempty"`
	Attachments  []io.Reader `json:"-" url:"-"`
	// The input message to provide to the mesh. The structure of this object must conform to the input message type schema defined on the Intent. Within the input message, any asset references may use the syntax `attachment://<filename>` format to refer to assets attached in the same multipart/form-data body.
	InputMessage map[string]any `json:"input_message" url:"-"`
	// Conversation options
	ConversationOptions *ConversationCreateOptions `json:"conversation_options,omitempty" url:"-"`
	// Request options
	RequestOptions *RequestCreateOptions `json:"request_options,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateConversationWithInitialRequestRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetIntentName sets the IntentName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequestRequest) SetIntentName(intentName string) {
	c.IntentName = intentName
	c.require(createConversationWithInitialRequestRequestFieldIntentName)
}

// SetMeshID sets the MeshID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequestRequest) SetMeshID(meshID string) {
	c.MeshID = meshID
	c.require(createConversationWithInitialRequestRequestFieldMeshID)
}

// SetVersionOrdinal sets the VersionOrdinal field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequestRequest) SetVersionOrdinal(versionOrdinal *int) {
	c.VersionOrdinal = versionOrdinal
	c.require(createConversationWithInitialRequestRequestFieldVersionOrdinal)
}

// SetVersionLabel sets the VersionLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequestRequest) SetVersionLabel(versionLabel *string) {
	c.VersionLabel = versionLabel
	c.require(createConversationWithInitialRequestRequestFieldVersionLabel)
}

func (c *CreateConversationWithInitialRequestRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConversationWithInitialRequestRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateConversationWithInitialRequestRequest(body)
	return nil
}

func (c *CreateConversationWithInitialRequestRequest) MarshalJSON() ([]byte, error) {
	type embed CreateConversationWithInitialRequestRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createEmptyConversationRequestFieldMeshID              = big.NewInt(1 << 0)
	createEmptyConversationRequestFieldVersionOrdinal      = big.NewInt(1 << 1)
	createEmptyConversationRequestFieldVersionLabel        = big.NewInt(1 << 2)
	createEmptyConversationRequestFieldTags                = big.NewInt(1 << 3)
	createEmptyConversationRequestFieldExternalID          = big.NewInt(1 << 4)
	createEmptyConversationRequestFieldNotificationChannel = big.NewInt(1 << 5)
	createEmptyConversationRequestFieldTracingEnabled      = big.NewInt(1 << 6)
)

type CreateEmptyConversationRequest struct {
	// The ID of the Mesh for which to create a conversation
	MeshID string `json:"-" url:"mesh_id"`
	// The version number of a mesh version for which to create the conversation
	VersionOrdinal *int `json:"-" url:"version_ordinal,omitempty"`
	// The version label of a mesh version for which to create the conversation
	VersionLabel *string `json:"-" url:"version_label,omitempty"`
	// Tags to apply to the conversation.
	Tags []*Tag `json:"tags,omitempty" url:"-"`
	// Optional client-chosen external ID for the new conversation: 1 to 255 bytes. Immutable and unique (per organization). Use `GET /conversations/lookup/external-id/{external_id}` to find the conversation again.
	ExternalID *string `json:"external_id,omitempty" url:"-"`
	// Optional notification channel for the new conversation: 1 to 255 bytes. Every request output the conversation writes is also appended to this channel, readable via `GET /notification-channels/{notification_channel_id}/events`. The channel is created implicitly on first use and cannot be changed later.
	NotificationChannel *string `json:"notification_channel,omitempty" url:"-"`
	// Optional. `false` turns off tracing for the new conversation, so the mesh records no trace for any of its requests. Defaults to `true`. Cannot be changed later.
	TracingEnabled *bool `json:"tracing_enabled,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateEmptyConversationRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetMeshID sets the MeshID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetMeshID(meshID string) {
	c.MeshID = meshID
	c.require(createEmptyConversationRequestFieldMeshID)
}

// SetVersionOrdinal sets the VersionOrdinal field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetVersionOrdinal(versionOrdinal *int) {
	c.VersionOrdinal = versionOrdinal
	c.require(createEmptyConversationRequestFieldVersionOrdinal)
}

// SetVersionLabel sets the VersionLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetVersionLabel(versionLabel *string) {
	c.VersionLabel = versionLabel
	c.require(createEmptyConversationRequestFieldVersionLabel)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetTags(tags []*Tag) {
	c.Tags = tags
	c.require(createEmptyConversationRequestFieldTags)
}

// SetExternalID sets the ExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetExternalID(externalID *string) {
	c.ExternalID = externalID
	c.require(createEmptyConversationRequestFieldExternalID)
}

// SetNotificationChannel sets the NotificationChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetNotificationChannel(notificationChannel *string) {
	c.NotificationChannel = notificationChannel
	c.require(createEmptyConversationRequestFieldNotificationChannel)
}

// SetTracingEnabled sets the TracingEnabled field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetTracingEnabled(tracingEnabled *bool) {
	c.TracingEnabled = tracingEnabled
	c.require(createEmptyConversationRequestFieldTracingEnabled)
}

func (c *CreateEmptyConversationRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateEmptyConversationRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateEmptyConversationRequest(body)
	return nil
}

func (c *CreateEmptyConversationRequest) MarshalJSON() ([]byte, error) {
	type embed CreateEmptyConversationRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	getConversationRequestFieldConversationID = big.NewInt(1 << 0)
)

type GetConversationRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetConversationRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequest) SetConversationID(conversationID string) {
	g.ConversationID = conversationID
	g.require(getConversationRequestFieldConversationID)
}

var (
	getConversationByExternalIDRequestFieldExternalID = big.NewInt(1 << 0)
)

type GetConversationByExternalIDRequest struct {
	// The external ID supplied when the conversation was created.
	ExternalID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetConversationByExternalIDRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetExternalID sets the ExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationByExternalIDRequest) SetExternalID(externalID string) {
	g.ExternalID = externalID
	g.require(getConversationByExternalIDRequestFieldExternalID)
}

var (
	getConversationRequestsRequestFieldConversationID = big.NewInt(1 << 0)
)

type GetConversationRequestsRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetConversationRequestsRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsRequest) SetConversationID(conversationID string) {
	g.ConversationID = conversationID
	g.require(getConversationRequestsRequestFieldConversationID)
}

var (
	getIntentsForConversationRequestFieldConversationID = big.NewInt(1 << 0)
)

type GetIntentsForConversationRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetIntentsForConversationRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIntentsForConversationRequest) SetConversationID(conversationID string) {
	g.ConversationID = conversationID
	g.require(getIntentsForConversationRequestFieldConversationID)
}

var (
	getIntentsForNewConversationRequestFieldMeshID         = big.NewInt(1 << 0)
	getIntentsForNewConversationRequestFieldVersionOrdinal = big.NewInt(1 << 1)
	getIntentsForNewConversationRequestFieldVersionLabel   = big.NewInt(1 << 2)
)

type GetIntentsForNewConversationRequest struct {
	// The ID of the Mesh for which to get intents
	MeshID string `json:"-" url:"mesh_id"`
	// The version number of a mesh version for which to get intents
	VersionOrdinal *int `json:"-" url:"version_ordinal,omitempty"`
	// The version label of a mesh version for which to get intents
	VersionLabel *string `json:"-" url:"version_label,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetIntentsForNewConversationRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetMeshID sets the MeshID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIntentsForNewConversationRequest) SetMeshID(meshID string) {
	g.MeshID = meshID
	g.require(getIntentsForNewConversationRequestFieldMeshID)
}

// SetVersionOrdinal sets the VersionOrdinal field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIntentsForNewConversationRequest) SetVersionOrdinal(versionOrdinal *int) {
	g.VersionOrdinal = versionOrdinal
	g.require(getIntentsForNewConversationRequestFieldVersionOrdinal)
}

// SetVersionLabel sets the VersionLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIntentsForNewConversationRequest) SetVersionLabel(versionLabel *string) {
	g.VersionLabel = versionLabel
	g.require(getIntentsForNewConversationRequestFieldVersionLabel)
}

var (
	getRequestEventsRequestFieldLastEventID    = big.NewInt(1 << 0)
	getRequestEventsRequestFieldConversationID = big.NewInt(1 << 1)
	getRequestEventsRequestFieldRequestID      = big.NewInt(1 << 2)
)

type GetRequestEventsRequest struct {
	// The ID of the last event received. If provided, the stream will start from the next event after this ID.
	LastEventID *string `json:"-" url:"-"`
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`
	// The ID of an existing request
	RequestID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetRequestEventsRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetLastEventID sets the LastEventID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestEventsRequest) SetLastEventID(lastEventID *string) {
	g.LastEventID = lastEventID
	g.require(getRequestEventsRequestFieldLastEventID)
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestEventsRequest) SetConversationID(conversationID string) {
	g.ConversationID = conversationID
	g.require(getRequestEventsRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestEventsRequest) SetRequestID(requestID string) {
	g.RequestID = requestID
	g.require(getRequestEventsRequestFieldRequestID)
}

var (
	getRequestResponseRequestFieldConversationID = big.NewInt(1 << 0)
	getRequestResponseRequestFieldRequestID      = big.NewInt(1 << 1)
	getRequestResponseRequestFieldWaitMs         = big.NewInt(1 << 2)
)

type GetRequestResponseRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`
	// The ID of an existing request
	RequestID string `json:"-" url:"-"`
	// How long to wait, in milliseconds, before answering 504. Defaults to 30000. A longer wait is shortened to 50000, so a proxy in front never sees a silent connection for a minute.
	WaitMs *int `json:"-" url:"wait_ms,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetRequestResponseRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestResponseRequest) SetConversationID(conversationID string) {
	g.ConversationID = conversationID
	g.require(getRequestResponseRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestResponseRequest) SetRequestID(requestID string) {
	g.RequestID = requestID
	g.require(getRequestResponseRequestFieldRequestID)
}

// SetWaitMs sets the WaitMs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestResponseRequest) SetWaitMs(waitMs *int) {
	g.WaitMs = waitMs
	g.require(getRequestResponseRequestFieldWaitMs)
}

var (
	getRequestTaskRequestFieldConversationID = big.NewInt(1 << 0)
	getRequestTaskRequestFieldRequestID      = big.NewInt(1 << 1)
	getRequestTaskRequestFieldWaitMs         = big.NewInt(1 << 2)
)

type GetRequestTaskRequest struct {
	// The ID of an existing conversation
	ConversationID string `json:"-" url:"-"`
	// The ID of an existing request
	RequestID string `json:"-" url:"-"`
	// How long to wait, in milliseconds, before answering 504. Defaults to 30000. A longer wait is shortened to 50000, so a proxy in front never sees a silent connection for a minute.
	WaitMs *int `json:"-" url:"wait_ms,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetRequestTaskRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestTaskRequest) SetConversationID(conversationID string) {
	g.ConversationID = conversationID
	g.require(getRequestTaskRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestTaskRequest) SetRequestID(requestID string) {
	g.RequestID = requestID
	g.require(getRequestTaskRequestFieldRequestID)
}

// SetWaitMs sets the WaitMs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestTaskRequest) SetWaitMs(waitMs *int) {
	g.WaitMs = waitMs
	g.require(getRequestTaskRequestFieldWaitMs)
}

var (
	attachmentInfoFieldReferenceURL = big.NewInt(1 << 0)
	attachmentInfoFieldFilename     = big.NewInt(1 << 1)
)

// Information about an uploaded attachment
type AttachmentInfo struct {
	// A URL that can be used to access the uploaded asset in references.
	ReferenceURL string `json:"reference_url" url:"reference_url"`
	// The original filename of the uploaded asset.
	Filename string `json:"filename" url:"filename"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AttachmentInfo) GetReferenceURL() string {
	if a == nil {
		return ""
	}
	return a.ReferenceURL
}

func (a *AttachmentInfo) GetFilename() string {
	if a == nil {
		return ""
	}
	return a.Filename
}

func (a *AttachmentInfo) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AttachmentInfo) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetReferenceURL sets the ReferenceURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AttachmentInfo) SetReferenceURL(referenceURL string) {
	a.ReferenceURL = referenceURL
	a.require(attachmentInfoFieldReferenceURL)
}

// SetFilename sets the Filename field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AttachmentInfo) SetFilename(filename string) {
	a.Filename = filename
	a.require(attachmentInfoFieldFilename)
}

func (a *AttachmentInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler AttachmentInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AttachmentInfo(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AttachmentInfo) MarshalJSON() ([]byte, error) {
	type embed AttachmentInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AttachmentInfo) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

var (
	conversationFieldConversationID = big.NewInt(1 << 0)
	conversationFieldExternalID     = big.NewInt(1 << 1)
	conversationFieldMeshID         = big.NewInt(1 << 2)
	conversationFieldMeshVersionID  = big.NewInt(1 << 3)
	conversationFieldTags           = big.NewInt(1 << 4)
	conversationFieldCreatedAt      = big.NewInt(1 << 5)
	conversationFieldUpdatedAt      = big.NewInt(1 << 6)
)

// A conversation and its client-facing attributes.
type Conversation struct {
	ConversationID string `json:"conversation_id" url:"conversation_id"`
	// The client-supplied external ID, if one was set at creation.
	ExternalID    *string `json:"external_id,omitempty" url:"external_id,omitempty"`
	MeshID        string  `json:"mesh_id" url:"mesh_id"`
	MeshVersionID string  `json:"mesh_version_id" url:"mesh_version_id"`
	// The tags supplied when the conversation was created. Empty when none were.
	Tags      []*Tag    `json:"tags" url:"tags"`
	CreatedAt time.Time `json:"created_at" url:"created_at"`
	UpdatedAt time.Time `json:"updated_at" url:"updated_at"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *Conversation) GetConversationID() string {
	if c == nil {
		return ""
	}
	return c.ConversationID
}

func (c *Conversation) GetExternalID() *string {
	if c == nil {
		return nil
	}
	return c.ExternalID
}

func (c *Conversation) GetMeshID() string {
	if c == nil {
		return ""
	}
	return c.MeshID
}

func (c *Conversation) GetMeshVersionID() string {
	if c == nil {
		return ""
	}
	return c.MeshVersionID
}

func (c *Conversation) GetTags() []*Tag {
	if c == nil {
		return nil
	}
	return c.Tags
}

func (c *Conversation) GetCreatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.CreatedAt
}

func (c *Conversation) GetUpdatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.UpdatedAt
}

func (c *Conversation) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *Conversation) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(conversationFieldConversationID)
}

// SetExternalID sets the ExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetExternalID(externalID *string) {
	c.ExternalID = externalID
	c.require(conversationFieldExternalID)
}

// SetMeshID sets the MeshID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetMeshID(meshID string) {
	c.MeshID = meshID
	c.require(conversationFieldMeshID)
}

// SetMeshVersionID sets the MeshVersionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetMeshVersionID(meshVersionID string) {
	c.MeshVersionID = meshVersionID
	c.require(conversationFieldMeshVersionID)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetTags(tags []*Tag) {
	c.Tags = tags
	c.require(conversationFieldTags)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetCreatedAt(createdAt time.Time) {
	c.CreatedAt = createdAt
	c.require(conversationFieldCreatedAt)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetUpdatedAt(updatedAt time.Time) {
	c.UpdatedAt = updatedAt
	c.require(conversationFieldUpdatedAt)
}

func (c *Conversation) UnmarshalJSON(data []byte) error {
	type embed Conversation
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = Conversation(unmarshaler.embed)
	c.CreatedAt = unmarshaler.CreatedAt.Time()
	c.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *Conversation) MarshalJSON() ([]byte, error) {
	type embed Conversation
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*c),
		CreatedAt: internal.NewDateTime(c.CreatedAt),
		UpdatedAt: internal.NewDateTime(c.UpdatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *Conversation) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

var (
	conversationCreateOptionsFieldTags                = big.NewInt(1 << 0)
	conversationCreateOptionsFieldExternalID          = big.NewInt(1 << 1)
	conversationCreateOptionsFieldNotificationChannel = big.NewInt(1 << 2)
	conversationCreateOptionsFieldTracingEnabled      = big.NewInt(1 << 3)
)

// Settings for conversation.
type ConversationCreateOptions struct {
	// Tags to apply to the conversation.
	Tags []*Tag `json:"tags,omitempty" url:"tags,omitempty"`
	// Optional client-chosen external ID for the new conversation: 1 to 255 bytes. Immutable and unique (per organization). Use `GET /conversations/lookup/external-id/{external_id}` to find the conversation again.
	ExternalID *string `json:"external_id,omitempty" url:"external_id,omitempty"`
	// Optional notification channel for the new conversation: 1 to 255 bytes. Every request output the conversation writes is also appended to this channel, readable via `GET /notification-channels/{notification_channel_id}/events`. The channel is created implicitly on first use and cannot be changed later.
	NotificationChannel *string `json:"notification_channel,omitempty" url:"notification_channel,omitempty"`
	// Optional. `false` turns off tracing for the new conversation, so the mesh records no trace for any of its requests. Defaults to `true`. Cannot be changed later.
	TracingEnabled *bool `json:"tracing_enabled,omitempty" url:"tracing_enabled,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ConversationCreateOptions) GetTags() []*Tag {
	if c == nil {
		return nil
	}
	return c.Tags
}

func (c *ConversationCreateOptions) GetExternalID() *string {
	if c == nil {
		return nil
	}
	return c.ExternalID
}

func (c *ConversationCreateOptions) GetNotificationChannel() *string {
	if c == nil {
		return nil
	}
	return c.NotificationChannel
}

func (c *ConversationCreateOptions) GetTracingEnabled() *bool {
	if c == nil {
		return nil
	}
	return c.TracingEnabled
}

func (c *ConversationCreateOptions) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ConversationCreateOptions) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationCreateOptions) SetTags(tags []*Tag) {
	c.Tags = tags
	c.require(conversationCreateOptionsFieldTags)
}

// SetExternalID sets the ExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationCreateOptions) SetExternalID(externalID *string) {
	c.ExternalID = externalID
	c.require(conversationCreateOptionsFieldExternalID)
}

// SetNotificationChannel sets the NotificationChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationCreateOptions) SetNotificationChannel(notificationChannel *string) {
	c.NotificationChannel = notificationChannel
	c.require(conversationCreateOptionsFieldNotificationChannel)
}

// SetTracingEnabled sets the TracingEnabled field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationCreateOptions) SetTracingEnabled(tracingEnabled *bool) {
	c.TracingEnabled = tracingEnabled
	c.require(conversationCreateOptionsFieldTracingEnabled)
}

func (c *ConversationCreateOptions) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationCreateOptions
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationCreateOptions(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConversationCreateOptions) MarshalJSON() ([]byte, error) {
	type embed ConversationCreateOptions
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ConversationCreateOptions) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

var (
	createConversationResponseFieldConversationID = big.NewInt(1 << 0)
	createConversationResponseFieldIntents        = big.NewInt(1 << 1)
)

// Response returned when a conversation is created without an initial request.
type CreateConversationResponse struct {
	ConversationID string      `json:"conversation_id" url:"conversation_id"`
	Intents        IntentsList `json:"intents" url:"intents"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateConversationResponse) GetConversationID() string {
	if c == nil {
		return ""
	}
	return c.ConversationID
}

func (c *CreateConversationResponse) GetIntents() IntentsList {
	if c == nil {
		return nil
	}
	return c.Intents
}

func (c *CreateConversationResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateConversationResponse) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationResponse) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(createConversationResponseFieldConversationID)
}

// SetIntents sets the Intents field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationResponse) SetIntents(intents IntentsList) {
	c.Intents = intents
	c.require(createConversationResponseFieldIntents)
}

func (c *CreateConversationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConversationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateConversationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateConversationResponse) MarshalJSON() ([]byte, error) {
	type embed CreateConversationResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateConversationResponse) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

var (
	createRequestResponseFieldConversationID = big.NewInt(1 << 0)
	createRequestResponseFieldRequestID      = big.NewInt(1 << 1)
)

type CreateRequestResponse struct {
	ConversationID string `json:"conversation_id" url:"conversation_id"`
	RequestID      string `json:"request_id" url:"request_id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateRequestResponse) GetConversationID() string {
	if c == nil {
		return ""
	}
	return c.ConversationID
}

func (c *CreateRequestResponse) GetRequestID() string {
	if c == nil {
		return ""
	}
	return c.RequestID
}

func (c *CreateRequestResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateRequestResponse) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestResponse) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(createRequestResponseFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestResponse) SetRequestID(requestID string) {
	c.RequestID = requestID
	c.require(createRequestResponseFieldRequestID)
}

func (c *CreateRequestResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateRequestResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateRequestResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateRequestResponse) MarshalJSON() ([]byte, error) {
	type embed CreateRequestResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateRequestResponse) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

var (
	getTaskResponseFieldTaskID = big.NewInt(1 << 0)
	getTaskResponseFieldData   = big.NewInt(1 << 1)
)

type GetTaskResponse struct {
	TaskID string    `json:"task_id" url:"task_id"`
	Data   *TaskData `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetTaskResponse) GetTaskID() string {
	if g == nil {
		return ""
	}
	return g.TaskID
}

func (g *GetTaskResponse) GetData() *TaskData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTaskResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetTaskResponse) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetTaskID sets the TaskID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTaskResponse) SetTaskID(taskID string) {
	g.TaskID = taskID
	g.require(getTaskResponseFieldTaskID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTaskResponse) SetData(data *TaskData) {
	g.Data = data
	g.require(getTaskResponseFieldData)
}

func (g *GetTaskResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetTaskResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetTaskResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetTaskResponse) MarshalJSON() ([]byte, error) {
	type embed GetTaskResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetTaskResponse) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// A list of intents available
type IntentsList = []*IntentsListItem

var (
	intentsListItemFieldName        = big.NewInt(1 << 0)
	intentsListItemFieldDescription = big.NewInt(1 << 1)
	intentsListItemFieldSchema      = big.NewInt(1 << 2)
	intentsListItemFieldResponses   = big.NewInt(1 << 3)
)

type IntentsListItem struct {
	Name        string                          `json:"name" url:"name"`
	Description *string                         `json:"description,omitempty" url:"description,omitempty"`
	Schema      map[string]any                  `json:"schema" url:"schema"`
	Responses   []*IntentsListItemResponsesItem `json:"responses,omitempty" url:"responses,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (i *IntentsListItem) GetName() string {
	if i == nil {
		return ""
	}
	return i.Name
}

func (i *IntentsListItem) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}

func (i *IntentsListItem) GetSchema() map[string]any {
	if i == nil {
		return nil
	}
	return i.Schema
}

func (i *IntentsListItem) GetResponses() []*IntentsListItemResponsesItem {
	if i == nil {
		return nil
	}
	return i.Responses
}

func (i *IntentsListItem) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i *IntentsListItem) require(field *big.Int) {
	next := new(big.Int)
	if i.explicitFields != nil {
		next.Set(i.explicitFields)
	}
	next.Or(next, field)
	i.explicitFields = next
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentsListItem) SetName(name string) {
	i.Name = name
	i.require(intentsListItemFieldName)
}

// SetDescription sets the Description field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentsListItem) SetDescription(description *string) {
	i.Description = description
	i.require(intentsListItemFieldDescription)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentsListItem) SetSchema(schema map[string]any) {
	i.Schema = schema
	i.require(intentsListItemFieldSchema)
}

// SetResponses sets the Responses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentsListItem) SetResponses(responses []*IntentsListItemResponsesItem) {
	i.Responses = responses
	i.require(intentsListItemFieldResponses)
}

func (i *IntentsListItem) UnmarshalJSON(data []byte) error {
	type unmarshaler IntentsListItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IntentsListItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	i.rawJSON = json.RawMessage(data)
	return nil
}

func (i *IntentsListItem) MarshalJSON() ([]byte, error) {
	type embed IntentsListItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, i.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (i *IntentsListItem) String() string {
	if i == nil {
		return "<nil>"
	}
	if len(i.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(i.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

var (
	intentsListItemResponsesItemFieldName        = big.NewInt(1 << 0)
	intentsListItemResponsesItemFieldDescription = big.NewInt(1 << 1)
	intentsListItemResponsesItemFieldSchema      = big.NewInt(1 << 2)
)

type IntentsListItemResponsesItem struct {
	Name        *string        `json:"name,omitempty" url:"name,omitempty"`
	Description *string        `json:"description,omitempty" url:"description,omitempty"`
	Schema      map[string]any `json:"schema,omitempty" url:"schema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (i *IntentsListItemResponsesItem) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *IntentsListItemResponsesItem) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}

func (i *IntentsListItemResponsesItem) GetSchema() map[string]any {
	if i == nil {
		return nil
	}
	return i.Schema
}

func (i *IntentsListItemResponsesItem) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i *IntentsListItemResponsesItem) require(field *big.Int) {
	next := new(big.Int)
	if i.explicitFields != nil {
		next.Set(i.explicitFields)
	}
	next.Or(next, field)
	i.explicitFields = next
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentsListItemResponsesItem) SetName(name *string) {
	i.Name = name
	i.require(intentsListItemResponsesItemFieldName)
}

// SetDescription sets the Description field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentsListItemResponsesItem) SetDescription(description *string) {
	i.Description = description
	i.require(intentsListItemResponsesItemFieldDescription)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentsListItemResponsesItem) SetSchema(schema map[string]any) {
	i.Schema = schema
	i.require(intentsListItemResponsesItemFieldSchema)
}

func (i *IntentsListItemResponsesItem) UnmarshalJSON(data []byte) error {
	type unmarshaler IntentsListItemResponsesItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IntentsListItemResponsesItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	i.rawJSON = json.RawMessage(data)
	return nil
}

func (i *IntentsListItemResponsesItem) MarshalJSON() ([]byte, error) {
	type embed IntentsListItemResponsesItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, i.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (i *IntentsListItemResponsesItem) String() string {
	if i == nil {
		return "<nil>"
	}
	if len(i.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(i.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

var (
	listIntentsResponseFieldIntents = big.NewInt(1 << 0)
)

// Response returned when listing intents for a mesh version.
type ListIntentsResponse struct {
	Intents IntentsList `json:"intents" url:"intents"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListIntentsResponse) GetIntents() IntentsList {
	if l == nil {
		return nil
	}
	return l.Intents
}

func (l *ListIntentsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListIntentsResponse) require(field *big.Int) {
	next := new(big.Int)
	if l.explicitFields != nil {
		next.Set(l.explicitFields)
	}
	next.Or(next, field)
	l.explicitFields = next
}

// SetIntents sets the Intents field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListIntentsResponse) SetIntents(intents IntentsList) {
	l.Intents = intents
	l.require(listIntentsResponseFieldIntents)
}

func (l *ListIntentsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListIntentsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListIntentsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListIntentsResponse) MarshalJSON() ([]byte, error) {
	type embed ListIntentsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListIntentsResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	requestCreateOptionsFieldTags = big.NewInt(1 << 0)
)

// Settings for request.
type RequestCreateOptions struct {
	// Tags to apply to the request.
	Tags []*Tag `json:"tags,omitempty" url:"tags,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RequestCreateOptions) GetTags() []*Tag {
	if r == nil {
		return nil
	}
	return r.Tags
}

func (r *RequestCreateOptions) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *RequestCreateOptions) require(field *big.Int) {
	next := new(big.Int)
	if r.explicitFields != nil {
		next.Set(r.explicitFields)
	}
	next.Or(next, field)
	r.explicitFields = next
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RequestCreateOptions) SetTags(tags []*Tag) {
	r.Tags = tags
	r.require(requestCreateOptionsFieldTags)
}

func (r *RequestCreateOptions) UnmarshalJSON(data []byte) error {
	type unmarshaler RequestCreateOptions
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RequestCreateOptions(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RequestCreateOptions) MarshalJSON() ([]byte, error) {
	type embed RequestCreateOptions
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *RequestCreateOptions) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// A chunk of data in the event stream
type ResponseEventData struct {
	Type         string
	Task         *TaskEventDetails
	TaskResponse *TaskResponseEventDetails
	Message      *MessageEventDetails
	Thinking     *ThinkingEventDetails
	Response     *ResponseEventDetails
	Error        *ErrorEventDetails
	Cancelled    *CancelEventDetails
	ServerError  *ServerErrorEventDetails

	rawJSON json.RawMessage
}

func (r *ResponseEventData) GetType() string {
	if r == nil {
		return ""
	}
	return r.Type
}

func (r *ResponseEventData) GetTask() *TaskEventDetails {
	if r == nil {
		return nil
	}
	return r.Task
}

func (r *ResponseEventData) GetTaskResponse() *TaskResponseEventDetails {
	if r == nil {
		return nil
	}
	return r.TaskResponse
}

func (r *ResponseEventData) GetMessage() *MessageEventDetails {
	if r == nil {
		return nil
	}
	return r.Message
}

func (r *ResponseEventData) GetThinking() *ThinkingEventDetails {
	if r == nil {
		return nil
	}
	return r.Thinking
}

func (r *ResponseEventData) GetResponse() *ResponseEventDetails {
	if r == nil {
		return nil
	}
	return r.Response
}

func (r *ResponseEventData) GetError() *ErrorEventDetails {
	if r == nil {
		return nil
	}
	return r.Error
}

func (r *ResponseEventData) GetCancelled() *CancelEventDetails {
	if r == nil {
		return nil
	}
	return r.Cancelled
}

func (r *ResponseEventData) GetServerError() *ServerErrorEventDetails {
	if r == nil {
		return nil
	}
	return r.ServerError
}

func (r *ResponseEventData) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	r.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", r)
	}
	switch unmarshaler.Type {
	case "task":
		value := new(TaskEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Task = value
	case "task_response":
		value := new(TaskResponseEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TaskResponse = value
	case "message":
		value := new(MessageEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Message = value
	case "thinking":
		value := new(ThinkingEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Thinking = value
	case "response":
		value := new(ResponseEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Response = value
	case "error":
		value := new(ErrorEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Error = value
	case "cancelled":
		value := new(CancelEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Cancelled = value
	case "server_error":
		value := new(ServerErrorEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.ServerError = value
	}
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r ResponseEventData) MarshalJSON() ([]byte, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	if r.Task != nil {
		return internal.MarshalJSONWithExtraProperty(r.Task, "type", "task")
	}
	if r.TaskResponse != nil {
		return internal.MarshalJSONWithExtraProperty(r.TaskResponse, "type", "task_response")
	}
	if r.Message != nil {
		return internal.MarshalJSONWithExtraProperty(r.Message, "type", "message")
	}
	if r.Thinking != nil {
		return internal.MarshalJSONWithExtraProperty(r.Thinking, "type", "thinking")
	}
	if r.Response != nil {
		return internal.MarshalJSONWithExtraProperty(r.Response, "type", "response")
	}
	if r.Error != nil {
		return internal.MarshalJSONWithExtraProperty(r.Error, "type", "error")
	}
	if r.Cancelled != nil {
		return internal.MarshalJSONWithExtraProperty(r.Cancelled, "type", "cancelled")
	}
	if r.ServerError != nil {
		return internal.MarshalJSONWithExtraProperty(r.ServerError, "type", "server_error")
	}
	if len(r.rawJSON) > 0 {
		return r.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", r)
}

type ResponseEventDataVisitor interface {
	VisitTask(*TaskEventDetails) error
	VisitTaskResponse(*TaskResponseEventDetails) error
	VisitMessage(*MessageEventDetails) error
	VisitThinking(*ThinkingEventDetails) error
	VisitResponse(*ResponseEventDetails) error
	VisitError(*ErrorEventDetails) error
	VisitCancelled(*CancelEventDetails) error
	VisitServerError(*ServerErrorEventDetails) error
}

func (r *ResponseEventData) Accept(visitor ResponseEventDataVisitor) error {
	if r.Task != nil {
		return visitor.VisitTask(r.Task)
	}
	if r.TaskResponse != nil {
		return visitor.VisitTaskResponse(r.TaskResponse)
	}
	if r.Message != nil {
		return visitor.VisitMessage(r.Message)
	}
	if r.Thinking != nil {
		return visitor.VisitThinking(r.Thinking)
	}
	if r.Response != nil {
		return visitor.VisitResponse(r.Response)
	}
	if r.Error != nil {
		return visitor.VisitError(r.Error)
	}
	if r.Cancelled != nil {
		return visitor.VisitCancelled(r.Cancelled)
	}
	if r.ServerError != nil {
		return visitor.VisitServerError(r.ServerError)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", r)
}

func (r *ResponseEventData) validate() error {
	if r == nil {
		return fmt.Errorf("type %T is nil", r)
	}
	var fields []string
	if r.Task != nil {
		fields = append(fields, "task")
	}
	if r.TaskResponse != nil {
		fields = append(fields, "task_response")
	}
	if r.Message != nil {
		fields = append(fields, "message")
	}
	if r.Thinking != nil {
		fields = append(fields, "thinking")
	}
	if r.Response != nil {
		fields = append(fields, "response")
	}
	if r.Error != nil {
		fields = append(fields, "error")
	}
	if r.Cancelled != nil {
		fields = append(fields, "cancelled")
	}
	if r.ServerError != nil {
		fields = append(fields, "server_error")
	}
	if len(fields) == 0 {
		if r.Type != "" {
			if len(r.rawJSON) > 0 {
				return nil
			}
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", r, r.Type)
		}
		return fmt.Errorf("type %T is empty", r)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", r, fields)
	}
	if r.Type != "" {
		field := fields[0]
		if r.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				r,
				r.Type,
				r,
			)
		}
	}
	return nil
}

var (
	tagFieldName  = big.NewInt(1 << 0)
	tagFieldValue = big.NewInt(1 << 1)
)

// A name/value tag.
type Tag struct {
	Name  string `json:"name" url:"name"`
	Value string `json:"value" url:"value"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *Tag) GetName() string {
	if t == nil {
		return ""
	}
	return t.Name
}

func (t *Tag) GetValue() string {
	if t == nil {
		return ""
	}
	return t.Value
}

func (t *Tag) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *Tag) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *Tag) SetName(name string) {
	t.Name = name
	t.require(tagFieldName)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *Tag) SetValue(value string) {
	t.Value = value
	t.require(tagFieldValue)
}

func (t *Tag) UnmarshalJSON(data []byte) error {
	type unmarshaler Tag
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = Tag(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *Tag) MarshalJSON() ([]byte, error) {
	type embed Tag
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *Tag) String() string {
	if t == nil {
		return "<nil>"
	}
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

var (
	attachAssetsToConversationResponseFieldAttachments = big.NewInt(1 << 0)
)

type AttachAssetsToConversationResponse struct {
	// A list of information about the uploaded assets.
	Attachments []*AttachmentInfo `json:"attachments" url:"attachments"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AttachAssetsToConversationResponse) GetAttachments() []*AttachmentInfo {
	if a == nil {
		return nil
	}
	return a.Attachments
}

func (a *AttachAssetsToConversationResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AttachAssetsToConversationResponse) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetAttachments sets the Attachments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AttachAssetsToConversationResponse) SetAttachments(attachments []*AttachmentInfo) {
	a.Attachments = attachments
	a.require(attachAssetsToConversationResponseFieldAttachments)
}

func (a *AttachAssetsToConversationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AttachAssetsToConversationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AttachAssetsToConversationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AttachAssetsToConversationResponse) MarshalJSON() ([]byte, error) {
	type embed AttachAssetsToConversationResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AttachAssetsToConversationResponse) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

var (
	getConversationRequestsResponseFieldRequests = big.NewInt(1 << 0)
)

type GetConversationRequestsResponse struct {
	// A list of requests within the conversation.
	Requests []*GetConversationRequestsResponseRequestsItem `json:"requests" url:"requests"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetConversationRequestsResponse) GetRequests() []*GetConversationRequestsResponseRequestsItem {
	if g == nil {
		return nil
	}
	return g.Requests
}

func (g *GetConversationRequestsResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetConversationRequestsResponse) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetRequests sets the Requests field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponse) SetRequests(requests []*GetConversationRequestsResponseRequestsItem) {
	g.Requests = requests
	g.require(getConversationRequestsResponseFieldRequests)
}

func (g *GetConversationRequestsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetConversationRequestsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetConversationRequestsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetConversationRequestsResponse) MarshalJSON() ([]byte, error) {
	type embed GetConversationRequestsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetConversationRequestsResponse) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

var (
	getConversationRequestsResponseRequestsItemFieldRequestID    = big.NewInt(1 << 0)
	getConversationRequestsResponseRequestsItemFieldInputMessage = big.NewInt(1 << 1)
	getConversationRequestsResponseRequestsItemFieldStatus       = big.NewInt(1 << 2)
	getConversationRequestsResponseRequestsItemFieldTags         = big.NewInt(1 << 3)
	getConversationRequestsResponseRequestsItemFieldEvents       = big.NewInt(1 << 4)
)

type GetConversationRequestsResponseRequestsItem struct {
	RequestID    string                                            `json:"request_id" url:"request_id"`
	InputMessage map[string]any                                    `json:"input_message" url:"input_message"`
	Status       GetConversationRequestsResponseRequestsItemStatus `json:"status" url:"status"`
	// The tags supplied when the request was created.
	Tags   []*Tag              `json:"tags" url:"tags"`
	Events []*MessageEventData `json:"events" url:"events"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetConversationRequestsResponseRequestsItem) GetRequestID() string {
	if g == nil {
		return ""
	}
	return g.RequestID
}

func (g *GetConversationRequestsResponseRequestsItem) GetInputMessage() map[string]any {
	if g == nil {
		return nil
	}
	return g.InputMessage
}

func (g *GetConversationRequestsResponseRequestsItem) GetStatus() GetConversationRequestsResponseRequestsItemStatus {
	if g == nil {
		return ""
	}
	return g.Status
}

func (g *GetConversationRequestsResponseRequestsItem) GetTags() []*Tag {
	if g == nil {
		return nil
	}
	return g.Tags
}

func (g *GetConversationRequestsResponseRequestsItem) GetEvents() []*MessageEventData {
	if g == nil {
		return nil
	}
	return g.Events
}

func (g *GetConversationRequestsResponseRequestsItem) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetConversationRequestsResponseRequestsItem) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetRequestID(requestID string) {
	g.RequestID = requestID
	g.require(getConversationRequestsResponseRequestsItemFieldRequestID)
}

// SetInputMessage sets the InputMessage field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetInputMessage(inputMessage map[string]any) {
	g.InputMessage = inputMessage
	g.require(getConversationRequestsResponseRequestsItemFieldInputMessage)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetStatus(status GetConversationRequestsResponseRequestsItemStatus) {
	g.Status = status
	g.require(getConversationRequestsResponseRequestsItemFieldStatus)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetTags(tags []*Tag) {
	g.Tags = tags
	g.require(getConversationRequestsResponseRequestsItemFieldTags)
}

// SetEvents sets the Events field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetEvents(events []*MessageEventData) {
	g.Events = events
	g.require(getConversationRequestsResponseRequestsItemFieldEvents)
}

func (g *GetConversationRequestsResponseRequestsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler GetConversationRequestsResponseRequestsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetConversationRequestsResponseRequestsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetConversationRequestsResponseRequestsItem) MarshalJSON() ([]byte, error) {
	type embed GetConversationRequestsResponseRequestsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetConversationRequestsResponseRequestsItem) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GetConversationRequestsResponseRequestsItemStatus string

const (
	GetConversationRequestsResponseRequestsItemStatusActive    GetConversationRequestsResponseRequestsItemStatus = "active"
	GetConversationRequestsResponseRequestsItemStatusCancelled GetConversationRequestsResponseRequestsItemStatus = "cancelled"
	GetConversationRequestsResponseRequestsItemStatusComplete  GetConversationRequestsResponseRequestsItemStatus = "complete"
)

func NewGetConversationRequestsResponseRequestsItemStatusFromString(s string) (GetConversationRequestsResponseRequestsItemStatus, error) {
	switch s {
	case "active":
		return GetConversationRequestsResponseRequestsItemStatusActive, nil
	case "cancelled":
		return GetConversationRequestsResponseRequestsItemStatusCancelled, nil
	case "complete":
		return GetConversationRequestsResponseRequestsItemStatusComplete, nil
	}
	var t GetConversationRequestsResponseRequestsItemStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetConversationRequestsResponseRequestsItemStatus) Ptr() *GetConversationRequestsResponseRequestsItemStatus {
	return &g
}
