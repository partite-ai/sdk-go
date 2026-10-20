// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/partite-ai/sdk-go/internal"
	big "math/big"
)

var (
	createNotificationChannelRequestFieldNotificationChannelID = big.NewInt(1 << 0)
)

type CreateNotificationChannelRequest struct {
	// The client-supplied id of the channel to create.
	NotificationChannelID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateNotificationChannelRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetNotificationChannelID sets the NotificationChannelID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateNotificationChannelRequest) SetNotificationChannelID(notificationChannelID string) {
	c.NotificationChannelID = notificationChannelID
	c.require(createNotificationChannelRequestFieldNotificationChannelID)
}

var (
	getNotificationChannelEventsRequestFieldLastEventID           = big.NewInt(1 << 0)
	getNotificationChannelEventsRequestFieldNotificationChannelID = big.NewInt(1 << 1)
)

type GetNotificationChannelEventsRequest struct {
	// An event id (`rnco_…`) from a previous event on this channel. If provided, reading starts after that event. Event ids are ordered and unique per channel, not globally; an id is only meaningful for the channel it came from.
	LastEventID *string `json:"-" url:"-"`
	// The notification channel id supplied at conversation creation.
	NotificationChannelID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetNotificationChannelEventsRequest) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
}

// SetLastEventID sets the LastEventID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetNotificationChannelEventsRequest) SetLastEventID(lastEventID *string) {
	g.LastEventID = lastEventID
	g.require(getNotificationChannelEventsRequestFieldLastEventID)
}

// SetNotificationChannelID sets the NotificationChannelID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetNotificationChannelEventsRequest) SetNotificationChannelID(notificationChannelID string) {
	g.NotificationChannelID = notificationChannelID
	g.require(getNotificationChannelEventsRequestFieldNotificationChannelID)
}

var (
	notificationChannelCreatedFieldID = big.NewInt(1 << 0)
)

// The created (or pre-existing) notification channel.
type NotificationChannelCreated struct {
	// The client-supplied channel id.
	ID string `json:"id" url:"id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NotificationChannelCreated) GetID() string {
	if n == nil {
		return ""
	}
	return n.ID
}

func (n *NotificationChannelCreated) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n *NotificationChannelCreated) require(field *big.Int) {
	next := new(big.Int)
	if n.explicitFields != nil {
		next.Set(n.explicitFields)
	}
	next.Or(next, field)
	n.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelCreated) SetID(id string) {
	n.ID = id
	n.require(notificationChannelCreatedFieldID)
}

func (n *NotificationChannelCreated) UnmarshalJSON(data []byte) error {
	type unmarshaler NotificationChannelCreated
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NotificationChannelCreated(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NotificationChannelCreated) MarshalJSON() ([]byte, error) {
	type embed NotificationChannelCreated
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, n.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (n *NotificationChannelCreated) String() string {
	if n == nil {
		return "<nil>"
	}
	if len(n.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(n.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

var (
	notificationChannelEventFieldType                   = big.NewInt(1 << 0)
	notificationChannelEventFieldID                     = big.NewInt(1 << 1)
	notificationChannelEventFieldConversationID         = big.NewInt(1 << 2)
	notificationChannelEventFieldConversationExternalID = big.NewInt(1 << 3)
	notificationChannelEventFieldRequestID              = big.NewInt(1 << 4)
	notificationChannelEventFieldEvent                  = big.NewInt(1 << 5)
)

// One event delivered on a notification channel: a request's event, exactly as the request's events endpoint delivers it, wrapped in the channel's event id and the conversation and request ids.
type NotificationChannelEvent struct {
	Type NotificationChannelEventType `json:"type" url:"type"`
	// The event's opaque id (`rnco_…`). Send it as `Last-Event-ID` to resume this channel after this event. Ids are ordered and unique per channel, not globally.
	ID string `json:"id" url:"id"`
	// The conversation that produced the event.
	ConversationID string `json:"conversation_id" url:"conversation_id"`
	// The conversation's client-chosen external ID, when one was set.
	ConversationExternalID *string `json:"conversation_external_id,omitempty" url:"conversation_external_id,omitempty"`
	// The request that produced the event.
	RequestID string `json:"request_id" url:"request_id"`
	// The request's event. Its `id` can be sent as `Last-Event-ID` to the request's events endpoint.
	Event *RequestEvent `json:"event" url:"event"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NotificationChannelEvent) GetType() NotificationChannelEventType {
	if n == nil {
		return ""
	}
	return n.Type
}

func (n *NotificationChannelEvent) GetID() string {
	if n == nil {
		return ""
	}
	return n.ID
}

func (n *NotificationChannelEvent) GetConversationID() string {
	if n == nil {
		return ""
	}
	return n.ConversationID
}

func (n *NotificationChannelEvent) GetConversationExternalID() *string {
	if n == nil {
		return nil
	}
	return n.ConversationExternalID
}

func (n *NotificationChannelEvent) GetRequestID() string {
	if n == nil {
		return ""
	}
	return n.RequestID
}

func (n *NotificationChannelEvent) GetEvent() *RequestEvent {
	if n == nil {
		return nil
	}
	return n.Event
}

func (n *NotificationChannelEvent) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n *NotificationChannelEvent) require(field *big.Int) {
	next := new(big.Int)
	if n.explicitFields != nil {
		next.Set(n.explicitFields)
	}
	next.Or(next, field)
	n.explicitFields = next
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEvent) SetType(type_ NotificationChannelEventType) {
	n.Type = type_
	n.require(notificationChannelEventFieldType)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEvent) SetID(id string) {
	n.ID = id
	n.require(notificationChannelEventFieldID)
}

// SetConversationID sets the ConversationID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEvent) SetConversationID(conversationID string) {
	n.ConversationID = conversationID
	n.require(notificationChannelEventFieldConversationID)
}

// SetConversationExternalID sets the ConversationExternalID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEvent) SetConversationExternalID(conversationExternalID *string) {
	n.ConversationExternalID = conversationExternalID
	n.require(notificationChannelEventFieldConversationExternalID)
}

// SetRequestID sets the RequestID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEvent) SetRequestID(requestID string) {
	n.RequestID = requestID
	n.require(notificationChannelEventFieldRequestID)
}

// SetEvent sets the Event field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEvent) SetEvent(event *RequestEvent) {
	n.Event = event
	n.require(notificationChannelEventFieldEvent)
}

func (n *NotificationChannelEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler NotificationChannelEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NotificationChannelEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NotificationChannelEvent) MarshalJSON() ([]byte, error) {
	type embed NotificationChannelEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, n.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (n *NotificationChannelEvent) String() string {
	if n == nil {
		return "<nil>"
	}
	if len(n.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(n.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

type NotificationChannelEventType string

const (
	NotificationChannelEventTypeNotification NotificationChannelEventType = "notification"
)

func NewNotificationChannelEventTypeFromString(s string) (NotificationChannelEventType, error) {
	switch s {
	case "notification":
		return NotificationChannelEventTypeNotification, nil
	}
	var t NotificationChannelEventType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (n NotificationChannelEventType) Ptr() *NotificationChannelEventType {
	return &n
}

// One line of a notification channel stream, either an event or a server error.
type NotificationChannelStreamData struct {
	Type         string
	Notification *NotificationChannelEvent
	ServerError  *ServerErrorEvent

	rawJSON json.RawMessage
}

func (n *NotificationChannelStreamData) GetType() string {
	if n == nil {
		return ""
	}
	return n.Type
}

func (n *NotificationChannelStreamData) GetNotification() *NotificationChannelEvent {
	if n == nil {
		return nil
	}
	return n.Notification
}

func (n *NotificationChannelStreamData) GetServerError() *ServerErrorEvent {
	if n == nil {
		return nil
	}
	return n.ServerError
}

func (n *NotificationChannelStreamData) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	n.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", n)
	}
	switch unmarshaler.Type {
	case "notification":
		value := new(NotificationChannelEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		n.Notification = value
	case "server_error":
		value := new(ServerErrorEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		n.ServerError = value
	}
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n NotificationChannelStreamData) MarshalJSON() ([]byte, error) {
	if err := n.validate(); err != nil {
		return nil, err
	}
	if n.Notification != nil {
		return internal.MarshalJSONWithExtraProperty(n.Notification, "type", "notification")
	}
	if n.ServerError != nil {
		return internal.MarshalJSONWithExtraProperty(n.ServerError, "type", "server_error")
	}
	if len(n.rawJSON) > 0 {
		return n.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", n)
}

type NotificationChannelStreamDataVisitor interface {
	VisitNotification(*NotificationChannelEvent) error
	VisitServerError(*ServerErrorEvent) error
}

func (n *NotificationChannelStreamData) Accept(visitor NotificationChannelStreamDataVisitor) error {
	if n.Notification != nil {
		return visitor.VisitNotification(n.Notification)
	}
	if n.ServerError != nil {
		return visitor.VisitServerError(n.ServerError)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", n)
}

func (n *NotificationChannelStreamData) validate() error {
	if n == nil {
		return fmt.Errorf("type %T is nil", n)
	}
	var fields []string
	if n.Notification != nil {
		fields = append(fields, "notification")
	}
	if n.ServerError != nil {
		fields = append(fields, "server_error")
	}
	if len(fields) == 0 {
		if n.Type != "" {
			if len(n.rawJSON) > 0 {
				return nil
			}
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", n, n.Type)
		}
		return fmt.Errorf("type %T is empty", n)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", n, fields)
	}
	if n.Type != "" {
		field := fields[0]
		if n.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				n,
				n.Type,
				n,
			)
		}
	}
	return nil
}

// One stored event of a request.
type RequestEvent struct {
	Type         string
	Task         *TaskEvent
	TaskResponse *TaskResponseEvent
	Message      *MessageEvent
	Thinking     *ThinkingEvent
	Response     *ResponseEvent
	Error        *ErrorEvent
	Cancelled    *CancelledEvent

	rawJSON json.RawMessage
}

func (r *RequestEvent) GetType() string {
	if r == nil {
		return ""
	}
	return r.Type
}

func (r *RequestEvent) GetTask() *TaskEvent {
	if r == nil {
		return nil
	}
	return r.Task
}

func (r *RequestEvent) GetTaskResponse() *TaskResponseEvent {
	if r == nil {
		return nil
	}
	return r.TaskResponse
}

func (r *RequestEvent) GetMessage() *MessageEvent {
	if r == nil {
		return nil
	}
	return r.Message
}

func (r *RequestEvent) GetThinking() *ThinkingEvent {
	if r == nil {
		return nil
	}
	return r.Thinking
}

func (r *RequestEvent) GetResponse() *ResponseEvent {
	if r == nil {
		return nil
	}
	return r.Response
}

func (r *RequestEvent) GetError() *ErrorEvent {
	if r == nil {
		return nil
	}
	return r.Error
}

func (r *RequestEvent) GetCancelled() *CancelledEvent {
	if r == nil {
		return nil
	}
	return r.Cancelled
}

func (r *RequestEvent) UnmarshalJSON(data []byte) error {
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
	}
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r RequestEvent) MarshalJSON() ([]byte, error) {
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
	if len(r.rawJSON) > 0 {
		return r.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", r)
}

type RequestEventVisitor interface {
	VisitTask(*TaskEvent) error
	VisitTaskResponse(*TaskResponseEvent) error
	VisitMessage(*MessageEvent) error
	VisitThinking(*ThinkingEvent) error
	VisitResponse(*ResponseEvent) error
	VisitError(*ErrorEvent) error
	VisitCancelled(*CancelledEvent) error
}

func (r *RequestEvent) Accept(visitor RequestEventVisitor) error {
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
	return fmt.Errorf("type %T does not define a non-empty union type", r)
}

func (r *RequestEvent) validate() error {
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
