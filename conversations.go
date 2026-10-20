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
	addResourcesRequestFieldConversationID = big.NewInt(1 << 0)
)

type AddResourcesRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string      `json:"-" url:"-"`
	Files          []io.Reader `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (a *AddResourcesRequest) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AddResourcesRequest) SetConversationID(conversationID string) {
	a.ConversationID = conversationID
	a.require(addResourcesRequestFieldConversationID)
}

func (a *AddResourcesRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler AddResourcesRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*a = AddResourcesRequest(body)
	return nil
}

func (a *AddResourcesRequest) MarshalJSON() ([]byte, error) {
	type embed AddResourcesRequest
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
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`
	// The ID of a request of the conversation.
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
	cancelTaskRequestFieldConversationID = big.NewInt(1 << 0)
	cancelTaskRequestFieldRequestID      = big.NewInt(1 << 1)
	cancelTaskRequestFieldTaskID         = big.NewInt(1 << 2)
)

type CancelTaskRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`
	// The ID of a request of the conversation.
	RequestID string `json:"-" url:"-"`
	// The ID of an open task of the request.
	TaskID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CancelTaskRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelTaskRequest) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(cancelTaskRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelTaskRequest) SetRequestID(requestID string) {
	c.RequestID = requestID
	c.require(cancelTaskRequestFieldRequestID)
}

// SetTaskID sets the TaskID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelTaskRequest) SetTaskID(taskID string) {
	c.TaskID = taskID
	c.require(cancelTaskRequestFieldTaskID)
}

var (
	completeTaskRequestFieldConversationID = big.NewInt(1 << 0)
	completeTaskRequestFieldRequestID      = big.NewInt(1 << 1)
	completeTaskRequestFieldTaskID         = big.NewInt(1 << 2)
)

type CompleteTaskRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`
	// The ID of a request of the conversation.
	RequestID string `json:"-" url:"-"`
	// The ID of an open task of the request.
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
	createConversationRequestFieldWait = big.NewInt(1 << 0)
)

type CreateConversationRequest struct {
	// Seconds to hold the answer until the request leaves `running`. A wait longer than 50 seconds is shortened to 50.
	Wait        *int                      `json:"-" url:"wait,omitempty"`
	Attachments []io.Reader               `json:"-" url:"-"`
	Params      *CreateConversationParams `json:"params" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateConversationRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetWait sets the Wait field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationRequest) SetWait(wait *int) {
	c.Wait = wait
	c.require(createConversationRequestFieldWait)
}

func (c *CreateConversationRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConversationRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateConversationRequest(body)
	return nil
}

func (c *CreateConversationRequest) MarshalJSON() ([]byte, error) {
	type embed CreateConversationRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createRequestRequestFieldConversationID = big.NewInt(1 << 0)
	createRequestRequestFieldWait           = big.NewInt(1 << 1)
)

type CreateRequestRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`
	// Seconds to hold the answer until the request leaves `running`. A wait longer than 50 seconds is shortened to 50.
	Wait        *int                 `json:"-" url:"wait,omitempty"`
	Attachments []io.Reader          `json:"-" url:"-"`
	Params      *CreateRequestParams `json:"params" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateRequestRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestRequest) SetConversationID(conversationID string) {
	c.ConversationID = conversationID
	c.require(createRequestRequestFieldConversationID)
}

// SetWait sets the Wait field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestRequest) SetWait(wait *int) {
	c.Wait = wait
	c.require(createRequestRequestFieldWait)
}

func (c *CreateRequestRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateRequestRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateRequestRequest(body)
	return nil
}

func (c *CreateRequestRequest) MarshalJSON() ([]byte, error) {
	type embed CreateRequestRequest
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
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
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
	getRequestRequestFieldConversationID = big.NewInt(1 << 0)
	getRequestRequestFieldRequestID      = big.NewInt(1 << 1)
	getRequestRequestFieldWait           = big.NewInt(1 << 2)
)

type GetRequestRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`
	// The ID of a request of the conversation.
	RequestID string `json:"-" url:"-"`
	// Seconds to hold the answer until the request leaves `running`. A wait longer than 50 seconds is shortened to 50.
	Wait *int `json:"-" url:"wait,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetRequestRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestRequest) SetConversationID(conversationID string) {
	g.ConversationID = conversationID
	g.require(getRequestRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestRequest) SetRequestID(requestID string) {
	g.RequestID = requestID
	g.require(getRequestRequestFieldRequestID)
}

// SetWait sets the Wait field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestRequest) SetWait(wait *int) {
	g.Wait = wait
	g.require(getRequestRequestFieldWait)
}

var (
	listIntentsRequestFieldConversationID = big.NewInt(1 << 0)
)

type ListIntentsRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListIntentsRequest) require(field *big.Int) {
	next := new(big.Int)
	if l.explicitFields != nil {
		next.Set(l.explicitFields)
	}
	next.Or(next, field)
	l.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListIntentsRequest) SetConversationID(conversationID string) {
	l.ConversationID = conversationID
	l.require(listIntentsRequestFieldConversationID)
}

var (
	listRequestsRequestFieldConversationID = big.NewInt(1 << 0)
)

type ListRequestsRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListRequestsRequest) require(field *big.Int) {
	next := new(big.Int)
	if l.explicitFields != nil {
		next.Set(l.explicitFields)
	}
	next.Or(next, field)
	l.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListRequestsRequest) SetConversationID(conversationID string) {
	l.ConversationID = conversationID
	l.require(listRequestsRequestFieldConversationID)
}

var (
	listResourcesRequestFieldConversationID = big.NewInt(1 << 0)
)

type ListResourcesRequest struct {
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListResourcesRequest) require(field *big.Int) {
	next := new(big.Int)
	if l.explicitFields != nil {
		next.Set(l.explicitFields)
	}
	next.Or(next, field)
	l.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListResourcesRequest) SetConversationID(conversationID string) {
	l.ConversationID = conversationID
	l.require(listResourcesRequestFieldConversationID)
}

var (
	streamRequestEventsRequestFieldLastEventID    = big.NewInt(1 << 0)
	streamRequestEventsRequestFieldConversationID = big.NewInt(1 << 1)
	streamRequestEventsRequestFieldRequestID      = big.NewInt(1 << 2)
)

type StreamRequestEventsRequest struct {
	// The ID of the last event received. The stream starts after it.
	LastEventID *string `json:"-" url:"-"`
	// A conversation ID, or `extid=` followed by the percent-encoded external ID.
	ConversationID string `json:"-" url:"-"`
	// The ID of a request of the conversation.
	RequestID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (s *StreamRequestEventsRequest) require(field *big.Int) {
	next := new(big.Int)
	if s.explicitFields != nil {
		next.Set(s.explicitFields)
	}
	next.Or(next, field)
	s.explicitFields = next
}

// SetLastEventID sets the LastEventID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *StreamRequestEventsRequest) SetLastEventID(lastEventID *string) {
	s.LastEventID = lastEventID
	s.require(streamRequestEventsRequestFieldLastEventID)
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *StreamRequestEventsRequest) SetConversationID(conversationID string) {
	s.ConversationID = conversationID
	s.require(streamRequestEventsRequestFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *StreamRequestEventsRequest) SetRequestID(requestID string) {
	s.RequestID = requestID
	s.require(streamRequestEventsRequestFieldRequestID)
}

var (
	conversationFieldID         = big.NewInt(1 << 0)
	conversationFieldExternalID = big.NewInt(1 << 1)
	conversationFieldMesh       = big.NewInt(1 << 2)
	conversationFieldTags       = big.NewInt(1 << 3)
	conversationFieldCreatedAt  = big.NewInt(1 << 4)
	conversationFieldUpdatedAt  = big.NewInt(1 << 5)
)

type Conversation struct {
	ID         string  `json:"id" url:"id"`
	ExternalID *string `json:"external_id,omitempty" url:"external_id,omitempty"`
	// The mesh the conversation runs on, and the version it runs on.
	Mesh      *ConversationMeshVersion `json:"mesh" url:"mesh"`
	Tags      []*Tag                   `json:"tags" url:"tags"`
	CreatedAt time.Time                `json:"created_at" url:"created_at"`
	UpdatedAt time.Time                `json:"updated_at" url:"updated_at"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *Conversation) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *Conversation) GetExternalID() *string {
	if c == nil {
		return nil
	}
	return c.ExternalID
}

func (c *Conversation) GetMesh() *ConversationMeshVersion {
	if c == nil {
		return nil
	}
	return c.Mesh
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

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetID(id string) {
	c.ID = id
	c.require(conversationFieldID)
}

// SetExternalID sets the ExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetExternalID(externalID *string) {
	c.ExternalID = externalID
	c.require(conversationFieldExternalID)
}

// SetMesh sets the Mesh field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *Conversation) SetMesh(mesh *ConversationMeshVersion) {
	c.Mesh = mesh
	c.require(conversationFieldMesh)
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
	conversationMeshFieldID           = big.NewInt(1 << 0)
	conversationMeshFieldVersion      = big.NewInt(1 << 1)
	conversationMeshFieldVersionLabel = big.NewInt(1 << 2)
	conversationMeshFieldDraft        = big.NewInt(1 << 3)
)

// The mesh a conversation runs on, and which of its versions: exactly one of `version`, `version_label` and `draft: true`, so the draft is never used by accident.
type ConversationMesh struct {
	ID string `json:"id" url:"id"`
	// The mesh version's number.
	Version *int `json:"version,omitempty" url:"version,omitempty"`
	// The label of the mesh version.
	VersionLabel *string `json:"version_label,omitempty" url:"version_label,omitempty"`
	// Use the mesh's draft. Must be `true` when given.
	Draft *bool `json:"draft,omitempty" url:"draft,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ConversationMesh) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *ConversationMesh) GetVersion() *int {
	if c == nil {
		return nil
	}
	return c.Version
}

func (c *ConversationMesh) GetVersionLabel() *string {
	if c == nil {
		return nil
	}
	return c.VersionLabel
}

func (c *ConversationMesh) GetDraft() *bool {
	if c == nil {
		return nil
	}
	return c.Draft
}

func (c *ConversationMesh) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ConversationMesh) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationMesh) SetID(id string) {
	c.ID = id
	c.require(conversationMeshFieldID)
}

// SetVersion sets the Version field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationMesh) SetVersion(version *int) {
	c.Version = version
	c.require(conversationMeshFieldVersion)
}

// SetVersionLabel sets the VersionLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationMesh) SetVersionLabel(versionLabel *string) {
	c.VersionLabel = versionLabel
	c.require(conversationMeshFieldVersionLabel)
}

// SetDraft sets the Draft field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationMesh) SetDraft(draft *bool) {
	c.Draft = draft
	c.require(conversationMeshFieldDraft)
}

func (c *ConversationMesh) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationMesh
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationMesh(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConversationMesh) MarshalJSON() ([]byte, error) {
	type embed ConversationMesh
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ConversationMesh) String() string {
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
	conversationMeshVersionFieldID        = big.NewInt(1 << 0)
	conversationMeshVersionFieldVersionID = big.NewInt(1 << 1)
)

// The mesh the conversation runs on, and the version it runs on.
type ConversationMeshVersion struct {
	ID        string `json:"id" url:"id"`
	VersionID string `json:"version_id" url:"version_id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ConversationMeshVersion) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *ConversationMeshVersion) GetVersionID() string {
	if c == nil {
		return ""
	}
	return c.VersionID
}

func (c *ConversationMeshVersion) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ConversationMeshVersion) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationMeshVersion) SetID(id string) {
	c.ID = id
	c.require(conversationMeshVersionFieldID)
}

// SetVersionID sets the VersionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConversationMeshVersion) SetVersionID(versionID string) {
	c.VersionID = versionID
	c.require(conversationMeshVersionFieldVersionID)
}

func (c *ConversationMeshVersion) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationMeshVersion
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationMeshVersion(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConversationMeshVersion) MarshalJSON() ([]byte, error) {
	type embed ConversationMeshVersion
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ConversationMeshVersion) String() string {
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
	createConversationParamsFieldMesh                = big.NewInt(1 << 0)
	createConversationParamsFieldRouteID             = big.NewInt(1 << 1)
	createConversationParamsFieldExternalID          = big.NewInt(1 << 2)
	createConversationParamsFieldTags                = big.NewInt(1 << 3)
	createConversationParamsFieldNotificationChannel = big.NewInt(1 << 4)
	createConversationParamsFieldTracing             = big.NewInt(1 << 5)
	createConversationParamsFieldIntent              = big.NewInt(1 << 6)
	createConversationParamsFieldInput               = big.NewInt(1 << 7)
)

// Name the target as `mesh` or as `route_id`, not both. A route picks its own version.
type CreateConversationParams struct {
	Mesh    *ConversationMesh `json:"mesh,omitempty" url:"mesh,omitempty"`
	RouteID *string           `json:"route_id,omitempty" url:"route_id,omitempty"`
	// A client-chosen ID, 1 to 255 bytes, unique and unchangeable. Address the conversation as `extid=<external ID>`.
	ExternalID *string `json:"external_id,omitempty" url:"external_id,omitempty"`
	// Tags of the conversation. Every request carries them alongside its own.
	Tags []*Tag `json:"tags,omitempty" url:"tags,omitempty"`
	// A notification channel that every request output of the conversation is also written to.
	NotificationChannel *string `json:"notification_channel,omitempty" url:"notification_channel,omitempty"`
	// `false` records no trace for any request of the conversation. Defaults to `true`.
	Tracing *bool `json:"tracing,omitempty" url:"tracing,omitempty"`
	// The intent of the first message. Required when the mesh has more than one.
	Intent *string `json:"intent,omitempty" url:"intent,omitempty"`
	// The first message, which starts a request. It matches the intent's input schema.
	Input map[string]any `json:"input,omitempty" url:"input,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateConversationParams) GetMesh() *ConversationMesh {
	if c == nil {
		return nil
	}
	return c.Mesh
}

func (c *CreateConversationParams) GetRouteID() *string {
	if c == nil {
		return nil
	}
	return c.RouteID
}

func (c *CreateConversationParams) GetExternalID() *string {
	if c == nil {
		return nil
	}
	return c.ExternalID
}

func (c *CreateConversationParams) GetTags() []*Tag {
	if c == nil {
		return nil
	}
	return c.Tags
}

func (c *CreateConversationParams) GetNotificationChannel() *string {
	if c == nil {
		return nil
	}
	return c.NotificationChannel
}

func (c *CreateConversationParams) GetTracing() *bool {
	if c == nil {
		return nil
	}
	return c.Tracing
}

func (c *CreateConversationParams) GetIntent() *string {
	if c == nil {
		return nil
	}
	return c.Intent
}

func (c *CreateConversationParams) GetInput() map[string]any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateConversationParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateConversationParams) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetMesh sets the Mesh field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetMesh(mesh *ConversationMesh) {
	c.Mesh = mesh
	c.require(createConversationParamsFieldMesh)
}

// SetRouteID sets the RouteID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetRouteID(routeID *string) {
	c.RouteID = routeID
	c.require(createConversationParamsFieldRouteID)
}

// SetExternalID sets the ExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetExternalID(externalID *string) {
	c.ExternalID = externalID
	c.require(createConversationParamsFieldExternalID)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetTags(tags []*Tag) {
	c.Tags = tags
	c.require(createConversationParamsFieldTags)
}

// SetNotificationChannel sets the NotificationChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetNotificationChannel(notificationChannel *string) {
	c.NotificationChannel = notificationChannel
	c.require(createConversationParamsFieldNotificationChannel)
}

// SetTracing sets the Tracing field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetTracing(tracing *bool) {
	c.Tracing = tracing
	c.require(createConversationParamsFieldTracing)
}

// SetIntent sets the Intent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetIntent(intent *string) {
	c.Intent = intent
	c.require(createConversationParamsFieldIntent)
}

// SetInput sets the Input field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationParams) SetInput(input map[string]any) {
	c.Input = input
	c.require(createConversationParamsFieldInput)
}

func (c *CreateConversationParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConversationParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateConversationParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateConversationParams) MarshalJSON() ([]byte, error) {
	type embed CreateConversationParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateConversationParams) String() string {
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
	createRequestParamsFieldIntent = big.NewInt(1 << 0)
	createRequestParamsFieldInput  = big.NewInt(1 << 1)
	createRequestParamsFieldTags   = big.NewInt(1 << 2)
)

type CreateRequestParams struct {
	// The intent the message is for. Required when the mesh has more than one.
	Intent *string `json:"intent,omitempty" url:"intent,omitempty"`
	// The message. It matches the intent's input schema.
	Input map[string]any `json:"input" url:"input"`
	// Tags of this request, alongside the conversation's.
	Tags []*Tag `json:"tags,omitempty" url:"tags,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateRequestParams) GetIntent() *string {
	if c == nil {
		return nil
	}
	return c.Intent
}

func (c *CreateRequestParams) GetInput() map[string]any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *CreateRequestParams) GetTags() []*Tag {
	if c == nil {
		return nil
	}
	return c.Tags
}

func (c *CreateRequestParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateRequestParams) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetIntent sets the Intent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestParams) SetIntent(intent *string) {
	c.Intent = intent
	c.require(createRequestParamsFieldIntent)
}

// SetInput sets the Input field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestParams) SetInput(input map[string]any) {
	c.Input = input
	c.require(createRequestParamsFieldInput)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestParams) SetTags(tags []*Tag) {
	c.Tags = tags
	c.require(createRequestParamsFieldTags)
}

func (c *CreateRequestParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateRequestParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateRequestParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateRequestParams) MarshalJSON() ([]byte, error) {
	type embed CreateRequestParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateRequestParams) String() string {
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
	createdConversationFieldID         = big.NewInt(1 << 0)
	createdConversationFieldExternalID = big.NewInt(1 << 1)
	createdConversationFieldMesh       = big.NewInt(1 << 2)
	createdConversationFieldTags       = big.NewInt(1 << 3)
	createdConversationFieldCreatedAt  = big.NewInt(1 << 4)
	createdConversationFieldUpdatedAt  = big.NewInt(1 << 5)
	createdConversationFieldRequest    = big.NewInt(1 << 6)
)

// The new conversation, with the request its first message started when it had one.
type CreatedConversation struct {
	ID         string  `json:"id" url:"id"`
	ExternalID *string `json:"external_id,omitempty" url:"external_id,omitempty"`
	// The mesh the conversation runs on, and the version it runs on.
	Mesh      *ConversationMeshVersion `json:"mesh" url:"mesh"`
	Tags      []*Tag                   `json:"tags" url:"tags"`
	CreatedAt time.Time                `json:"created_at" url:"created_at"`
	UpdatedAt time.Time                `json:"updated_at" url:"updated_at"`
	Request   *Request                 `json:"request,omitempty" url:"request,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreatedConversation) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *CreatedConversation) GetExternalID() *string {
	if c == nil {
		return nil
	}
	return c.ExternalID
}

func (c *CreatedConversation) GetMesh() *ConversationMeshVersion {
	if c == nil {
		return nil
	}
	return c.Mesh
}

func (c *CreatedConversation) GetTags() []*Tag {
	if c == nil {
		return nil
	}
	return c.Tags
}

func (c *CreatedConversation) GetCreatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.CreatedAt
}

func (c *CreatedConversation) GetUpdatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.UpdatedAt
}

func (c *CreatedConversation) GetRequest() *Request {
	if c == nil {
		return nil
	}
	return c.Request
}

func (c *CreatedConversation) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreatedConversation) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreatedConversation) SetID(id string) {
	c.ID = id
	c.require(createdConversationFieldID)
}

// SetExternalID sets the ExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreatedConversation) SetExternalID(externalID *string) {
	c.ExternalID = externalID
	c.require(createdConversationFieldExternalID)
}

// SetMesh sets the Mesh field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreatedConversation) SetMesh(mesh *ConversationMeshVersion) {
	c.Mesh = mesh
	c.require(createdConversationFieldMesh)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreatedConversation) SetTags(tags []*Tag) {
	c.Tags = tags
	c.require(createdConversationFieldTags)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreatedConversation) SetCreatedAt(createdAt time.Time) {
	c.CreatedAt = createdAt
	c.require(createdConversationFieldCreatedAt)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreatedConversation) SetUpdatedAt(updatedAt time.Time) {
	c.UpdatedAt = updatedAt
	c.require(createdConversationFieldUpdatedAt)
}

// SetRequest sets the Request field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreatedConversation) SetRequest(request *Request) {
	c.Request = request
	c.require(createdConversationFieldRequest)
}

func (c *CreatedConversation) UnmarshalJSON(data []byte) error {
	type embed CreatedConversation
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
	*c = CreatedConversation(unmarshaler.embed)
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

func (c *CreatedConversation) MarshalJSON() ([]byte, error) {
	type embed CreatedConversation
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

func (c *CreatedConversation) String() string {
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
	intentFieldName        = big.NewInt(1 << 0)
	intentFieldDescription = big.NewInt(1 << 1)
	intentFieldSchema      = big.NewInt(1 << 2)
	intentFieldResponses   = big.NewInt(1 << 3)
)

type Intent struct {
	Name        string  `json:"name" url:"name"`
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// The JSON Schema of the intent's input.
	Schema    map[string]any         `json:"schema" url:"schema"`
	Responses []*IntentResponsesItem `json:"responses" url:"responses"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (i *Intent) GetName() string {
	if i == nil {
		return ""
	}
	return i.Name
}

func (i *Intent) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}

func (i *Intent) GetSchema() map[string]any {
	if i == nil {
		return nil
	}
	return i.Schema
}

func (i *Intent) GetResponses() []*IntentResponsesItem {
	if i == nil {
		return nil
	}
	return i.Responses
}

func (i *Intent) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i *Intent) require(field *big.Int) {
	next := new(big.Int)
	if i.explicitFields != nil {
		next.Set(i.explicitFields)
	}
	next.Or(next, field)
	i.explicitFields = next
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *Intent) SetName(name string) {
	i.Name = name
	i.require(intentFieldName)
}

// SetDescription sets the Description field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *Intent) SetDescription(description *string) {
	i.Description = description
	i.require(intentFieldDescription)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *Intent) SetSchema(schema map[string]any) {
	i.Schema = schema
	i.require(intentFieldSchema)
}

// SetResponses sets the Responses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *Intent) SetResponses(responses []*IntentResponsesItem) {
	i.Responses = responses
	i.require(intentFieldResponses)
}

func (i *Intent) UnmarshalJSON(data []byte) error {
	type unmarshaler Intent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = Intent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	i.rawJSON = json.RawMessage(data)
	return nil
}

func (i *Intent) MarshalJSON() ([]byte, error) {
	type embed Intent
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, i.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (i *Intent) String() string {
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
	intentResponsesItemFieldName        = big.NewInt(1 << 0)
	intentResponsesItemFieldDescription = big.NewInt(1 << 1)
	intentResponsesItemFieldSchema      = big.NewInt(1 << 2)
)

type IntentResponsesItem struct {
	Name        string         `json:"name" url:"name"`
	Description *string        `json:"description,omitempty" url:"description,omitempty"`
	Schema      map[string]any `json:"schema" url:"schema"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (i *IntentResponsesItem) GetName() string {
	if i == nil {
		return ""
	}
	return i.Name
}

func (i *IntentResponsesItem) GetDescription() *string {
	if i == nil {
		return nil
	}
	return i.Description
}

func (i *IntentResponsesItem) GetSchema() map[string]any {
	if i == nil {
		return nil
	}
	return i.Schema
}

func (i *IntentResponsesItem) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i *IntentResponsesItem) require(field *big.Int) {
	next := new(big.Int)
	if i.explicitFields != nil {
		next.Set(i.explicitFields)
	}
	next.Or(next, field)
	i.explicitFields = next
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentResponsesItem) SetName(name string) {
	i.Name = name
	i.require(intentResponsesItemFieldName)
}

// SetDescription sets the Description field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentResponsesItem) SetDescription(description *string) {
	i.Description = description
	i.require(intentResponsesItemFieldDescription)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (i *IntentResponsesItem) SetSchema(schema map[string]any) {
	i.Schema = schema
	i.require(intentResponsesItemFieldSchema)
}

func (i *IntentResponsesItem) UnmarshalJSON(data []byte) error {
	type unmarshaler IntentResponsesItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IntentResponsesItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	i.rawJSON = json.RawMessage(data)
	return nil
}

func (i *IntentResponsesItem) MarshalJSON() ([]byte, error) {
	type embed IntentResponsesItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, i.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (i *IntentResponsesItem) String() string {
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
	openClientTaskFieldResponseSchema = big.NewInt(1 << 0)
	openClientTaskFieldClientTask     = big.NewInt(1 << 1)
	openClientTaskFieldID             = big.NewInt(1 << 2)
)

// A `client_task` task of a request.
type OpenClientTask struct {
	// The JSON Schema of the response that completes the task. It is the client task definition's output schema.
	ResponseSchema map[string]any `json:"response_schema" url:"response_schema"`
	ClientTask     *ClientTask    `json:"client_task" url:"client_task"`
	// The task's ID, which is also the ID of the `task` event that opened it.
	ID string `json:"id" url:"id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (o *OpenClientTask) GetResponseSchema() map[string]any {
	if o == nil {
		return nil
	}
	return o.ResponseSchema
}

func (o *OpenClientTask) GetClientTask() *ClientTask {
	if o == nil {
		return nil
	}
	return o.ClientTask
}

func (o *OpenClientTask) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OpenClientTask) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o *OpenClientTask) require(field *big.Int) {
	next := new(big.Int)
	if o.explicitFields != nil {
		next.Set(o.explicitFields)
	}
	next.Or(next, field)
	o.explicitFields = next
}

// SetResponseSchema sets the ResponseSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OpenClientTask) SetResponseSchema(responseSchema map[string]any) {
	o.ResponseSchema = responseSchema
	o.require(openClientTaskFieldResponseSchema)
}

// SetClientTask sets the ClientTask field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OpenClientTask) SetClientTask(clientTask *ClientTask) {
	o.ClientTask = clientTask
	o.require(openClientTaskFieldClientTask)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OpenClientTask) SetID(id string) {
	o.ID = id
	o.require(openClientTaskFieldID)
}

func (o *OpenClientTask) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenClientTask
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenClientTask(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	o.rawJSON = json.RawMessage(data)
	return nil
}

func (o *OpenClientTask) MarshalJSON() ([]byte, error) {
	type embed OpenClientTask
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, o.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (o *OpenClientTask) String() string {
	if o == nil {
		return "<nil>"
	}
	if len(o.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(o.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}

var (
	openToolCallApprovalFieldResponseSchema   = big.NewInt(1 << 0)
	openToolCallApprovalFieldToolCallApproval = big.NewInt(1 << 1)
	openToolCallApprovalFieldID               = big.NewInt(1 << 2)
)

// A `tool_call_approval` task of a request.
type OpenToolCallApproval struct {
	// The JSON Schema of `ToolCallApprovalResponse`.
	ResponseSchema   map[string]any    `json:"response_schema" url:"response_schema"`
	ToolCallApproval *ToolCallApproval `json:"tool_call_approval" url:"tool_call_approval"`
	// The task's ID, which is also the ID of the `task` event that opened it.
	ID string `json:"id" url:"id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (o *OpenToolCallApproval) GetResponseSchema() map[string]any {
	if o == nil {
		return nil
	}
	return o.ResponseSchema
}

func (o *OpenToolCallApproval) GetToolCallApproval() *ToolCallApproval {
	if o == nil {
		return nil
	}
	return o.ToolCallApproval
}

func (o *OpenToolCallApproval) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OpenToolCallApproval) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o *OpenToolCallApproval) require(field *big.Int) {
	next := new(big.Int)
	if o.explicitFields != nil {
		next.Set(o.explicitFields)
	}
	next.Or(next, field)
	o.explicitFields = next
}

// SetResponseSchema sets the ResponseSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OpenToolCallApproval) SetResponseSchema(responseSchema map[string]any) {
	o.ResponseSchema = responseSchema
	o.require(openToolCallApprovalFieldResponseSchema)
}

// SetToolCallApproval sets the ToolCallApproval field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OpenToolCallApproval) SetToolCallApproval(toolCallApproval *ToolCallApproval) {
	o.ToolCallApproval = toolCallApproval
	o.require(openToolCallApprovalFieldToolCallApproval)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (o *OpenToolCallApproval) SetID(id string) {
	o.ID = id
	o.require(openToolCallApprovalFieldID)
}

func (o *OpenToolCallApproval) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenToolCallApproval
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenToolCallApproval(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	o.rawJSON = json.RawMessage(data)
	return nil
}

func (o *OpenToolCallApproval) MarshalJSON() ([]byte, error) {
	type embed OpenToolCallApproval
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, o.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (o *OpenToolCallApproval) String() string {
	if o == nil {
		return "<nil>"
	}
	if len(o.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(o.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}

// An object with one key, `$ref`, whose value names the resource, such as `resource://r1`.
type Reference = map[string]string

var (
	requestFieldID        = big.NewInt(1 << 0)
	requestFieldStatus    = big.NewInt(1 << 1)
	requestFieldInput     = big.NewInt(1 << 2)
	requestFieldTags      = big.NewInt(1 << 3)
	requestFieldCreatedAt = big.NewInt(1 << 4)
	requestFieldResponse  = big.NewInt(1 << 5)
	requestFieldTasks     = big.NewInt(1 << 6)
	requestFieldError     = big.NewInt(1 << 7)
)

// A request and where it stands. `running` and `waiting` requests are not done: a `waiting` one waits on its open `tasks`. `completed` ones have a `response`, `failed` ones an `error`.
type Request struct {
	ID        string         `json:"id" url:"id"`
	Status    RequestStatus  `json:"status" url:"status"`
	Input     map[string]any `json:"input" url:"input"`
	Tags      []*Tag         `json:"tags" url:"tags"`
	CreatedAt time.Time      `json:"created_at" url:"created_at"`
	// The response, when `completed`. It matches one of the intent's response schemas.
	Response map[string]any `json:"response,omitempty" url:"response,omitempty"`
	// The open tasks, when `waiting`.
	Tasks []*Task `json:"tasks,omitempty" url:"tasks,omitempty"`
	// Why the request failed, when `failed`.
	Error *string `json:"error,omitempty" url:"error,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *Request) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *Request) GetStatus() RequestStatus {
	if r == nil {
		return ""
	}
	return r.Status
}

func (r *Request) GetInput() map[string]any {
	if r == nil {
		return nil
	}
	return r.Input
}

func (r *Request) GetTags() []*Tag {
	if r == nil {
		return nil
	}
	return r.Tags
}

func (r *Request) GetCreatedAt() time.Time {
	if r == nil {
		return time.Time{}
	}
	return r.CreatedAt
}

func (r *Request) GetResponse() map[string]any {
	if r == nil {
		return nil
	}
	return r.Response
}

func (r *Request) GetTasks() []*Task {
	if r == nil {
		return nil
	}
	return r.Tasks
}

func (r *Request) GetError() *string {
	if r == nil {
		return nil
	}
	return r.Error
}

func (r *Request) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *Request) require(field *big.Int) {
	next := new(big.Int)
	if r.explicitFields != nil {
		next.Set(r.explicitFields)
	}
	next.Or(next, field)
	r.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetID(id string) {
	r.ID = id
	r.require(requestFieldID)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetStatus(status RequestStatus) {
	r.Status = status
	r.require(requestFieldStatus)
}

// SetInput sets the Input field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetInput(input map[string]any) {
	r.Input = input
	r.require(requestFieldInput)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetTags(tags []*Tag) {
	r.Tags = tags
	r.require(requestFieldTags)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
	r.require(requestFieldCreatedAt)
}

// SetResponse sets the Response field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetResponse(response map[string]any) {
	r.Response = response
	r.require(requestFieldResponse)
}

// SetTasks sets the Tasks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetTasks(tasks []*Task) {
	r.Tasks = tasks
	r.require(requestFieldTasks)
}

// SetError sets the Error field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Request) SetError(error_ *string) {
	r.Error = error_
	r.require(requestFieldError)
}

func (r *Request) UnmarshalJSON(data []byte) error {
	type embed Request
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
	}{
		embed: embed(*r),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*r = Request(unmarshaler.embed)
	r.CreatedAt = unmarshaler.CreatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *Request) MarshalJSON() ([]byte, error) {
	type embed Request
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
	}{
		embed:     embed(*r),
		CreatedAt: internal.NewDateTime(r.CreatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *Request) String() string {
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

type RequestStatus string

const (
	RequestStatusRunning   RequestStatus = "running"
	RequestStatusWaiting   RequestStatus = "waiting"
	RequestStatusCompleted RequestStatus = "completed"
	RequestStatusFailed    RequestStatus = "failed"
	RequestStatusCancelled RequestStatus = "cancelled"
)

func NewRequestStatusFromString(s string) (RequestStatus, error) {
	switch s {
	case "running":
		return RequestStatusRunning, nil
	case "waiting":
		return RequestStatusWaiting, nil
	case "completed":
		return RequestStatusCompleted, nil
	case "failed":
		return RequestStatusFailed, nil
	case "cancelled":
		return RequestStatusCancelled, nil
	}
	var t RequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestStatus) Ptr() *RequestStatus {
	return &r
}

// One event of a request's stream.
type RequestStreamEvent struct {
	Type         string
	Started      *StartedEvent
	Task         *TaskEvent
	TaskResponse *TaskResponseEvent
	Message      *MessageEvent
	Thinking     *ThinkingEvent
	Response     *ResponseEvent
	Error        *ErrorEvent
	Cancelled    *CancelledEvent
	ServerError  *ServerErrorEvent

	rawJSON json.RawMessage
}

func (r *RequestStreamEvent) GetType() string {
	if r == nil {
		return ""
	}
	return r.Type
}

func (r *RequestStreamEvent) GetStarted() *StartedEvent {
	if r == nil {
		return nil
	}
	return r.Started
}

func (r *RequestStreamEvent) GetTask() *TaskEvent {
	if r == nil {
		return nil
	}
	return r.Task
}

func (r *RequestStreamEvent) GetTaskResponse() *TaskResponseEvent {
	if r == nil {
		return nil
	}
	return r.TaskResponse
}

func (r *RequestStreamEvent) GetMessage() *MessageEvent {
	if r == nil {
		return nil
	}
	return r.Message
}

func (r *RequestStreamEvent) GetThinking() *ThinkingEvent {
	if r == nil {
		return nil
	}
	return r.Thinking
}

func (r *RequestStreamEvent) GetResponse() *ResponseEvent {
	if r == nil {
		return nil
	}
	return r.Response
}

func (r *RequestStreamEvent) GetError() *ErrorEvent {
	if r == nil {
		return nil
	}
	return r.Error
}

func (r *RequestStreamEvent) GetCancelled() *CancelledEvent {
	if r == nil {
		return nil
	}
	return r.Cancelled
}

func (r *RequestStreamEvent) GetServerError() *ServerErrorEvent {
	if r == nil {
		return nil
	}
	return r.ServerError
}

func (r *RequestStreamEvent) UnmarshalJSON(data []byte) error {
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
	case "started":
		value := new(StartedEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Started = value
	case "task":
		value := new(TaskEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Task = value
	case "task_response":
		value := new(TaskResponseEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.TaskResponse = value
	case "message":
		value := new(MessageEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Message = value
	case "thinking":
		value := new(ThinkingEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Thinking = value
	case "response":
		value := new(ResponseEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Response = value
	case "error":
		value := new(ErrorEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Error = value
	case "cancelled":
		value := new(CancelledEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Cancelled = value
	case "server_error":
		value := new(ServerErrorEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.ServerError = value
	}
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r RequestStreamEvent) MarshalJSON() ([]byte, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	if r.Started != nil {
		return internal.MarshalJSONWithExtraProperty(r.Started, "type", "started")
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

type RequestStreamEventVisitor interface {
	VisitStarted(*StartedEvent) error
	VisitTask(*TaskEvent) error
	VisitTaskResponse(*TaskResponseEvent) error
	VisitMessage(*MessageEvent) error
	VisitThinking(*ThinkingEvent) error
	VisitResponse(*ResponseEvent) error
	VisitError(*ErrorEvent) error
	VisitCancelled(*CancelledEvent) error
	VisitServerError(*ServerErrorEvent) error
}

func (r *RequestStreamEvent) Accept(visitor RequestStreamEventVisitor) error {
	if r.Started != nil {
		return visitor.VisitStarted(r.Started)
	}
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

func (r *RequestStreamEvent) validate() error {
	if r == nil {
		return fmt.Errorf("type %T is nil", r)
	}
	var fields []string
	if r.Started != nil {
		fields = append(fields, "started")
	}
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
	resourceFieldReference   = big.NewInt(1 << 0)
	resourceFieldName        = big.NewInt(1 << 1)
	resourceFieldDescription = big.NewInt(1 << 2)
	resourceFieldMimeType    = big.NewInt(1 << 3)
)

// Something the conversation holds, such as a file.
type Resource struct {
	// The value to put in an input wherever it refers to this resource.
	Reference   Reference `json:"reference" url:"reference"`
	Name        string    `json:"name" url:"name"`
	Description string    `json:"description" url:"description"`
	MimeType    string    `json:"mime_type" url:"mime_type"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *Resource) GetReference() Reference {
	if r == nil {
		return nil
	}
	return r.Reference
}

func (r *Resource) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

func (r *Resource) GetDescription() string {
	if r == nil {
		return ""
	}
	return r.Description
}

func (r *Resource) GetMimeType() string {
	if r == nil {
		return ""
	}
	return r.MimeType
}

func (r *Resource) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *Resource) require(field *big.Int) {
	next := new(big.Int)
	if r.explicitFields != nil {
		next.Set(r.explicitFields)
	}
	next.Or(next, field)
	r.explicitFields = next
}

// SetReference sets the Reference field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Resource) SetReference(reference Reference) {
	r.Reference = reference
	r.require(resourceFieldReference)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Resource) SetName(name string) {
	r.Name = name
	r.require(resourceFieldName)
}

// SetDescription sets the Description field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Resource) SetDescription(description string) {
	r.Description = description
	r.require(resourceFieldDescription)
}

// SetMimeType sets the MimeType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *Resource) SetMimeType(mimeType string) {
	r.MimeType = mimeType
	r.require(resourceFieldMimeType)
}

func (r *Resource) UnmarshalJSON(data []byte) error {
	type unmarshaler Resource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = Resource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *Resource) MarshalJSON() ([]byte, error) {
	type embed Resource
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *Resource) String() string {
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

var (
	resourceListFieldResources = big.NewInt(1 << 0)
)

type ResourceList struct {
	Resources []*Resource `json:"resources" url:"resources"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ResourceList) GetResources() []*Resource {
	if r == nil {
		return nil
	}
	return r.Resources
}

func (r *ResourceList) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ResourceList) require(field *big.Int) {
	next := new(big.Int)
	if r.explicitFields != nil {
		next.Set(r.explicitFields)
	}
	next.Or(next, field)
	r.explicitFields = next
}

// SetResources sets the Resources field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResourceList) SetResources(resources []*Resource) {
	r.Resources = resources
	r.require(resourceListFieldResources)
}

func (r *ResourceList) UnmarshalJSON(data []byte) error {
	type unmarshaler ResourceList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResourceList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResourceList) MarshalJSON() ([]byte, error) {
	type embed ResourceList
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ResourceList) String() string {
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

var (
	startedEventFieldConversationID = big.NewInt(1 << 0)
	startedEventFieldRequestID      = big.NewInt(1 << 1)
)

// The first event of every stream. It names the conversation and the request, and has no event ID.
type StartedEvent struct {
	ConversationID string `json:"conversation_id" url:"conversation_id"`
	RequestID      string `json:"request_id" url:"request_id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *StartedEvent) GetConversationID() string {
	if s == nil {
		return ""
	}
	return s.ConversationID
}

func (s *StartedEvent) GetRequestID() string {
	if s == nil {
		return ""
	}
	return s.RequestID
}

func (s *StartedEvent) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *StartedEvent) require(field *big.Int) {
	next := new(big.Int)
	if s.explicitFields != nil {
		next.Set(s.explicitFields)
	}
	next.Or(next, field)
	s.explicitFields = next
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *StartedEvent) SetConversationID(conversationID string) {
	s.ConversationID = conversationID
	s.require(startedEventFieldConversationID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *StartedEvent) SetRequestID(requestID string) {
	s.RequestID = requestID
	s.require(startedEventFieldRequestID)
}

func (s *StartedEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler StartedEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StartedEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *StartedEvent) MarshalJSON() ([]byte, error) {
	type embed StartedEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *StartedEvent) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
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

// Something the user must do for the request to go on.
type Task struct {
	Kind             string
	ClientTask       *OpenClientTask
	ToolCallApproval *OpenToolCallApproval

	rawJSON json.RawMessage
}

func (t *Task) GetKind() string {
	if t == nil {
		return ""
	}
	return t.Kind
}

func (t *Task) GetClientTask() *OpenClientTask {
	if t == nil {
		return nil
	}
	return t.ClientTask
}

func (t *Task) GetToolCallApproval() *OpenToolCallApproval {
	if t == nil {
		return nil
	}
	return t.ToolCallApproval
}

func (t *Task) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", t)
	}
	switch unmarshaler.Kind {
	case "client_task":
		value := new(OpenClientTask)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ClientTask = value
	case "tool_call_approval":
		value := new(OpenToolCallApproval)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ToolCallApproval = value
	}
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t Task) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.ClientTask != nil {
		return internal.MarshalJSONWithExtraProperty(t.ClientTask, "kind", "client_task")
	}
	if t.ToolCallApproval != nil {
		return internal.MarshalJSONWithExtraProperty(t.ToolCallApproval, "kind", "tool_call_approval")
	}
	if len(t.rawJSON) > 0 {
		return t.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TaskVisitor interface {
	VisitClientTask(*OpenClientTask) error
	VisitToolCallApproval(*OpenToolCallApproval) error
}

func (t *Task) Accept(visitor TaskVisitor) error {
	if t.ClientTask != nil {
		return visitor.VisitClientTask(t.ClientTask)
	}
	if t.ToolCallApproval != nil {
		return visitor.VisitToolCallApproval(t.ToolCallApproval)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *Task) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.ClientTask != nil {
		fields = append(fields, "client_task")
	}
	if t.ToolCallApproval != nil {
		fields = append(fields, "tool_call_approval")
	}
	if len(fields) == 0 {
		if t.Kind != "" {
			if len(t.rawJSON) > 0 {
				return nil
			}
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.Kind)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.Kind != "" {
		field := fields[0]
		if t.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.Kind,
				t,
			)
		}
	}
	return nil
}

var (
	listIntentsResponseFieldIntents = big.NewInt(1 << 0)
)

type ListIntentsResponse struct {
	Intents []*Intent `json:"intents" url:"intents"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListIntentsResponse) GetIntents() []*Intent {
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
func (l *ListIntentsResponse) SetIntents(intents []*Intent) {
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
	listRequestsResponseFieldRequests = big.NewInt(1 << 0)
)

type ListRequestsResponse struct {
	Requests []*Request `json:"requests" url:"requests"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListRequestsResponse) GetRequests() []*Request {
	if l == nil {
		return nil
	}
	return l.Requests
}

func (l *ListRequestsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListRequestsResponse) require(field *big.Int) {
	next := new(big.Int)
	if l.explicitFields != nil {
		next.Set(l.explicitFields)
	}
	next.Or(next, field)
	l.explicitFields = next
}

// SetRequests sets the Requests field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListRequestsResponse) SetRequests(requests []*Request) {
	l.Requests = requests
	l.require(listRequestsResponseFieldRequests)
}

func (l *ListRequestsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListRequestsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListRequestsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListRequestsResponse) MarshalJSON() ([]byte, error) {
	type embed ListRequestsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListRequestsResponse) String() string {
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
