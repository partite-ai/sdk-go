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
	getNotificationChannelEventsRequestFieldLimit                 = big.NewInt(1 << 2)
)

type GetNotificationChannelEventsRequest struct {
	// An event id (`rnco_…`) from a previous event on this channel. If provided, reading starts after that event. Event ids are ordered and unique per channel, not globally; an id is only meaningful for the channel it came from.
	LastEventID *string `json:"-" url:"-"`
	// The notification channel id supplied at conversation creation.
	NotificationChannelID string `json:"-" url:"-"`
	// Maximum number of events to return. `application/json` mode only; default 100, maximum 1000.
	Limit *int `json:"-" url:"limit,omitempty"`

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

// SetLimit sets the Limit field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetNotificationChannelEventsRequest) SetLimit(limit *int) {
	g.Limit = limit
	g.require(getNotificationChannelEventsRequestFieldLimit)
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
	notificationChannelEventFieldID                     = big.NewInt(1 << 0)
	notificationChannelEventFieldConversationID         = big.NewInt(1 << 1)
	notificationChannelEventFieldConversationExternalID = big.NewInt(1 << 2)
	notificationChannelEventFieldRequestID              = big.NewInt(1 << 3)
	notificationChannelEventFieldOutput                 = big.NewInt(1 << 4)
)

// One event delivered on a notification channel: the request output, exactly as the per-request events endpoint delivers it, wrapped in the event id and routing ids.
type NotificationChannelEvent struct {
	// The event's opaque id (`rnco_…`). Send it as `Last-Event-ID` to resume this channel after this event. Ids are ordered and unique per channel, not globally.
	ID string `json:"id" url:"id"`
	// The conversation that produced the event.
	ConversationID string `json:"conversation_id" url:"conversation_id"`
	// The conversation's client-chosen external ID, when one was set.
	ConversationExternalID *string `json:"conversation_external_id,omitempty" url:"conversation_external_id,omitempty"`
	// The request that produced the event.
	RequestID string            `json:"request_id" url:"request_id"`
	Output    *MessageEventData `json:"output" url:"output"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
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

func (n *NotificationChannelEvent) GetOutput() *MessageEventData {
	if n == nil {
		return nil
	}
	return n.Output
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

// SetOutput sets the Output field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEvent) SetOutput(output *MessageEventData) {
	n.Output = output
	n.require(notificationChannelEventFieldOutput)
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

var (
	notificationChannelEventListFieldEvents = big.NewInt(1 << 0)
	notificationChannelEventListFieldNext   = big.NewInt(1 << 1)
)

// One page of notification channel events.
type NotificationChannelEventList struct {
	Events []*NotificationChannelEvent `json:"events" url:"events"`
	// The event id to resume from: the last event's id, or, when `events` is empty, the position the request read from. Send it as `Last-Event-ID`.
	Next string `json:"next" url:"next"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NotificationChannelEventList) GetEvents() []*NotificationChannelEvent {
	if n == nil {
		return nil
	}
	return n.Events
}

func (n *NotificationChannelEventList) GetNext() string {
	if n == nil {
		return ""
	}
	return n.Next
}

func (n *NotificationChannelEventList) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n *NotificationChannelEventList) require(field *big.Int) {
	next := new(big.Int)
	if n.explicitFields != nil {
		next.Set(n.explicitFields)
	}
	next.Or(next, field)
	n.explicitFields = next
}

// SetEvents sets the Events field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEventList) SetEvents(events []*NotificationChannelEvent) {
	n.Events = events
	n.require(notificationChannelEventListFieldEvents)
}

// SetNext sets the Next field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelEventList) SetNext(next string) {
	n.Next = next
	n.require(notificationChannelEventListFieldNext)
}

func (n *NotificationChannelEventList) UnmarshalJSON(data []byte) error {
	type unmarshaler NotificationChannelEventList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NotificationChannelEventList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NotificationChannelEventList) MarshalJSON() ([]byte, error) {
	type embed NotificationChannelEventList
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, n.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (n *NotificationChannelEventList) String() string {
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
