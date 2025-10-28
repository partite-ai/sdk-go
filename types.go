// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/partite-ai/sdk-go/internal"
	big "math/big"
)

// Details about a bad request error
var (
	badRequestErrorResponseFieldError     = big.NewInt(1 << 0)
	badRequestErrorResponseFieldParameter = big.NewInt(1 << 1)
)

type BadRequestErrorResponse struct {
	// A description of why the request was invalid.
	Error string `json:"error" url:"error"`
	// The name of the parameter that caused the error, if applicable.
	Parameter *string `json:"parameter,omitempty" url:"parameter,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BadRequestErrorResponse) GetError() string {
	if b == nil {
		return ""
	}
	return b.Error
}

func (b *BadRequestErrorResponse) GetParameter() *string {
	if b == nil {
		return nil
	}
	return b.Parameter
}

func (b *BadRequestErrorResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BadRequestErrorResponse) require(field *big.Int) {
	if b.explicitFields == nil {
		b.explicitFields = big.NewInt(0)
	}
	b.explicitFields.Or(b.explicitFields, field)
}

// SetError sets the Error field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BadRequestErrorResponse) SetError(error_ string) {
	b.Error = error_
	b.require(badRequestErrorResponseFieldError)
}

// SetParameter sets the Parameter field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BadRequestErrorResponse) SetParameter(parameter *string) {
	b.Parameter = parameter
	b.require(badRequestErrorResponseFieldParameter)
}

func (b *BadRequestErrorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BadRequestErrorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BadRequestErrorResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BadRequestErrorResponse) MarshalJSON() ([]byte, error) {
	type embed BadRequestErrorResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, b.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (b *BadRequestErrorResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Create a request and attachments.
var (
	createRequestWithAttachmentsRequestFieldInputMessage = big.NewInt(1 << 0)
	createRequestWithAttachmentsRequestFieldAttachments  = big.NewInt(1 << 1)
)

type CreateRequestWithAttachmentsRequest struct {
	// The input message to provide to the mesh. The structure of this object must conform to the input message type schema defined on the Intent. Within the input message, any asset references may use the syntax `attachment://<filename>` format to refer to assets attached in the same multipart/form-data body.
	InputMessage map[string]interface{} `json:"input_message" url:"input_message"`
	// Multiple parts, each representing a file upload including a `Content-Type` header and a `filename` attribute in the `Content-Disposition` header for the part
	Attachments []string `json:"attachments,omitempty" url:"attachments,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateRequestWithAttachmentsRequest) GetInputMessage() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.InputMessage
}

func (c *CreateRequestWithAttachmentsRequest) GetAttachments() []string {
	if c == nil {
		return nil
	}
	return c.Attachments
}

func (c *CreateRequestWithAttachmentsRequest) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateRequestWithAttachmentsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetInputMessage sets the InputMessage field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestWithAttachmentsRequest) SetInputMessage(inputMessage map[string]interface{}) {
	c.InputMessage = inputMessage
	c.require(createRequestWithAttachmentsRequestFieldInputMessage)
}

// SetAttachments sets the Attachments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestWithAttachmentsRequest) SetAttachments(attachments []string) {
	c.Attachments = attachments
	c.require(createRequestWithAttachmentsRequestFieldAttachments)
}

func (c *CreateRequestWithAttachmentsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateRequestWithAttachmentsRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateRequestWithAttachmentsRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateRequestWithAttachmentsRequest) MarshalJSON() ([]byte, error) {
	type embed CreateRequestWithAttachmentsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateRequestWithAttachmentsRequest) String() string {
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

// A server-sent event containing an error
var (
	errorEventFieldId   = big.NewInt(1 << 0)
	errorEventFieldData = big.NewInt(1 << 1)
)

type ErrorEvent struct {
	Id   *string                  `json:"id,omitempty" url:"id,omitempty"`
	Data *ServerErrorEventDetails `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *ErrorEvent) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *ErrorEvent) GetData() *ServerErrorEventDetails {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *ErrorEvent) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *ErrorEvent) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ErrorEvent) SetId(id *string) {
	e.Id = id
	e.require(errorEventFieldId)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ErrorEvent) SetData(data *ServerErrorEventDetails) {
	e.Data = data
	e.require(errorEventFieldData)
}

func (e *ErrorEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler ErrorEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ErrorEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *ErrorEvent) MarshalJSON() ([]byte, error) {
	type embed ErrorEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *ErrorEvent) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// Details about a generic error
var (
	genericErrorResponseFieldError = big.NewInt(1 << 0)
)

type GenericErrorResponse struct {
	// A description of the error.
	Error string `json:"error" url:"error"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GenericErrorResponse) GetError() string {
	if g == nil {
		return ""
	}
	return g.Error
}

func (g *GenericErrorResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GenericErrorResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetError sets the Error field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GenericErrorResponse) SetError(error_ string) {
	g.Error = error_
	g.require(genericErrorResponseFieldError)
}

func (g *GenericErrorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GenericErrorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GenericErrorResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GenericErrorResponse) MarshalJSON() ([]byte, error) {
	type embed GenericErrorResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GenericErrorResponse) String() string {
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

// A server-sent event containing a message
var (
	messageEventFieldId   = big.NewInt(1 << 0)
	messageEventFieldData = big.NewInt(1 << 1)
)

type MessageEvent struct {
	Id   string            `json:"id" url:"id"`
	Data *MessageEventData `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MessageEvent) GetId() string {
	if m == nil {
		return ""
	}
	return m.Id
}

func (m *MessageEvent) GetData() *MessageEventData {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MessageEvent) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MessageEvent) require(field *big.Int) {
	if m.explicitFields == nil {
		m.explicitFields = big.NewInt(0)
	}
	m.explicitFields.Or(m.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (m *MessageEvent) SetId(id string) {
	m.Id = id
	m.require(messageEventFieldId)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (m *MessageEvent) SetData(data *MessageEventData) {
	m.Data = data
	m.require(messageEventFieldData)
}

func (m *MessageEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler MessageEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MessageEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MessageEvent) MarshalJSON() ([]byte, error) {
	type embed MessageEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, m.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (m *MessageEvent) String() string {
	if len(m.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(m.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type SseEvent struct {
	Event   string
	Error   *ErrorEvent
	Message *MessageEvent
}

func (s *SseEvent) GetEvent() string {
	if s == nil {
		return ""
	}
	return s.Event
}

func (s *SseEvent) GetError() *ErrorEvent {
	if s == nil {
		return nil
	}
	return s.Error
}

func (s *SseEvent) GetMessage() *MessageEvent {
	if s == nil {
		return nil
	}
	return s.Message
}

func (s *SseEvent) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Event string `json:"event"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Event = unmarshaler.Event
	if unmarshaler.Event == "" {
		return fmt.Errorf("%T did not include discriminant event", s)
	}
	switch unmarshaler.Event {
	case "error":
		value := new(ErrorEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Error = value
	case "message":
		value := new(MessageEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Message = value
	}
	return nil
}

func (s SseEvent) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Error != nil {
		return internal.MarshalJSONWithExtraProperty(s.Error, "event", "error")
	}
	if s.Message != nil {
		return internal.MarshalJSONWithExtraProperty(s.Message, "event", "message")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SseEventVisitor interface {
	VisitError(*ErrorEvent) error
	VisitMessage(*MessageEvent) error
}

func (s *SseEvent) Accept(visitor SseEventVisitor) error {
	if s.Error != nil {
		return visitor.VisitError(s.Error)
	}
	if s.Message != nil {
		return visitor.VisitMessage(s.Message)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SseEvent) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Error != nil {
		fields = append(fields, "error")
	}
	if s.Message != nil {
		fields = append(fields, "message")
	}
	if len(fields) == 0 {
		if s.Event != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Event)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Event != "" {
		field := fields[0]
		if s.Event != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Event,
				s,
			)
		}
	}
	return nil
}
