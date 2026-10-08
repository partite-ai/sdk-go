// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/partite-ai/sdk-go/internal"
	big "math/big"
)

var (
	badRequestErrorResponseFieldError     = big.NewInt(1 << 0)
	badRequestErrorResponseFieldParameter = big.NewInt(1 << 1)
)

// Details about a bad request error
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
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b *BadRequestErrorResponse) require(field *big.Int) {
	next := new(big.Int)
	if b.explicitFields != nil {
		next.Set(b.explicitFields)
	}
	next.Or(next, field)
	b.explicitFields = next
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
	if b == nil {
		return "<nil>"
	}
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

var (
	cancelEventDetailsFieldID   = big.NewInt(1 << 0)
	cancelEventDetailsFieldData = big.NewInt(1 << 1)
)

// Details about a cancellation event
type CancelEventDetails struct {
	// The ID of the cancellation event
	ID string `json:"id" url:"id"`
	// A description of the cancellation.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CancelEventDetails) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *CancelEventDetails) GetData() string {
	if c == nil {
		return ""
	}
	return c.Data
}

func (c *CancelEventDetails) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CancelEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelEventDetails) SetID(id string) {
	c.ID = id
	c.require(cancelEventDetailsFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelEventDetails) SetData(data string) {
	c.Data = data
	c.require(cancelEventDetailsFieldData)
}

func (c *CancelEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler CancelEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CancelEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CancelEventDetails) MarshalJSON() ([]byte, error) {
	type embed CancelEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CancelEventDetails) String() string {
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
	createConversationWithInitialRequestFieldInputMessage        = big.NewInt(1 << 0)
	createConversationWithInitialRequestFieldConversationOptions = big.NewInt(1 << 1)
	createConversationWithInitialRequestFieldRequestOptions      = big.NewInt(1 << 2)
	createConversationWithInitialRequestFieldAttachments         = big.NewInt(1 << 3)
)

// Create a conversation with initial request and attachments.
type CreateConversationWithInitialRequest struct {
	// The input message to provide to the mesh. The structure of this object must conform to the input message type schema defined on the Intent. Within the input message, any asset references may use the syntax `attachment://<filename>` format to refer to assets attached in the same multipart/form-data body.
	InputMessage map[string]any `json:"input_message" url:"input_message"`
	// Conversation options
	ConversationOptions *ConversationCreateOptions `json:"conversation_options,omitempty" url:"conversation_options,omitempty"`
	// Request options
	RequestOptions *RequestCreateOptions `json:"request_options,omitempty" url:"request_options,omitempty"`
	// Multiple parts, each representing a file upload including a `Content-Type` header and a `filename` attribute in the `Content-Disposition` header for the part
	Attachments []string `json:"attachments,omitempty" url:"attachments,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateConversationWithInitialRequest) GetInputMessage() map[string]any {
	if c == nil {
		return nil
	}
	return c.InputMessage
}

func (c *CreateConversationWithInitialRequest) GetConversationOptions() *ConversationCreateOptions {
	if c == nil {
		return nil
	}
	return c.ConversationOptions
}

func (c *CreateConversationWithInitialRequest) GetRequestOptions() *RequestCreateOptions {
	if c == nil {
		return nil
	}
	return c.RequestOptions
}

func (c *CreateConversationWithInitialRequest) GetAttachments() []string {
	if c == nil {
		return nil
	}
	return c.Attachments
}

func (c *CreateConversationWithInitialRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateConversationWithInitialRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetInputMessage sets the InputMessage field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequest) SetInputMessage(inputMessage map[string]any) {
	c.InputMessage = inputMessage
	c.require(createConversationWithInitialRequestFieldInputMessage)
}

// SetConversationOptions sets the ConversationOptions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequest) SetConversationOptions(conversationOptions *ConversationCreateOptions) {
	c.ConversationOptions = conversationOptions
	c.require(createConversationWithInitialRequestFieldConversationOptions)
}

// SetRequestOptions sets the RequestOptions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequest) SetRequestOptions(requestOptions *RequestCreateOptions) {
	c.RequestOptions = requestOptions
	c.require(createConversationWithInitialRequestFieldRequestOptions)
}

// SetAttachments sets the Attachments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequest) SetAttachments(attachments []string) {
	c.Attachments = attachments
	c.require(createConversationWithInitialRequestFieldAttachments)
}

func (c *CreateConversationWithInitialRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConversationWithInitialRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateConversationWithInitialRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateConversationWithInitialRequest) MarshalJSON() ([]byte, error) {
	type embed CreateConversationWithInitialRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateConversationWithInitialRequest) String() string {
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
	createRequestWithAttachmentsRequestFieldInputMessage   = big.NewInt(1 << 0)
	createRequestWithAttachmentsRequestFieldRequestOptions = big.NewInt(1 << 1)
	createRequestWithAttachmentsRequestFieldAttachments    = big.NewInt(1 << 2)
)

// Create a request and attachments.
type CreateRequestWithAttachmentsRequest struct {
	// The input message to provide to the mesh. The structure of this object must conform to the input message type schema defined on the Intent. Within the input message, any asset references may use the syntax `attachment://<filename>` format to refer to assets attached in the same multipart/form-data body.
	InputMessage map[string]any `json:"input_message" url:"input_message"`
	// Request options
	RequestOptions *RequestCreateOptions `json:"request_options,omitempty" url:"request_options,omitempty"`
	// Multiple parts, each representing a file upload including a `Content-Type` header and a `filename` attribute in the `Content-Disposition` header for the part
	Attachments []string `json:"attachments,omitempty" url:"attachments,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateRequestWithAttachmentsRequest) GetInputMessage() map[string]any {
	if c == nil {
		return nil
	}
	return c.InputMessage
}

func (c *CreateRequestWithAttachmentsRequest) GetRequestOptions() *RequestCreateOptions {
	if c == nil {
		return nil
	}
	return c.RequestOptions
}

func (c *CreateRequestWithAttachmentsRequest) GetAttachments() []string {
	if c == nil {
		return nil
	}
	return c.Attachments
}

func (c *CreateRequestWithAttachmentsRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateRequestWithAttachmentsRequest) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetInputMessage sets the InputMessage field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestWithAttachmentsRequest) SetInputMessage(inputMessage map[string]any) {
	c.InputMessage = inputMessage
	c.require(createRequestWithAttachmentsRequestFieldInputMessage)
}

// SetRequestOptions sets the RequestOptions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestWithAttachmentsRequest) SetRequestOptions(requestOptions *RequestCreateOptions) {
	c.RequestOptions = requestOptions
	c.require(createRequestWithAttachmentsRequestFieldRequestOptions)
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
	errorEventFieldID   = big.NewInt(1 << 0)
	errorEventFieldData = big.NewInt(1 << 1)
)

// A server-sent event containing an error
type ErrorEvent struct {
	ID   *string                  `json:"id,omitempty" url:"id,omitempty"`
	Data *ServerErrorEventDetails `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *ErrorEvent) GetID() *string {
	if e == nil {
		return nil
	}
	return e.ID
}

func (e *ErrorEvent) GetData() *ServerErrorEventDetails {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *ErrorEvent) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *ErrorEvent) require(field *big.Int) {
	next := new(big.Int)
	if e.explicitFields != nil {
		next.Set(e.explicitFields)
	}
	next.Or(next, field)
	e.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ErrorEvent) SetID(id *string) {
	e.ID = id
	e.require(errorEventFieldID)
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
	if e == nil {
		return "<nil>"
	}
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

var (
	errorEventDetailsFieldID   = big.NewInt(1 << 0)
	errorEventDetailsFieldData = big.NewInt(1 << 1)
)

// Details about an error that occurred during request processing
type ErrorEventDetails struct {
	// The ID of the error event
	ID string `json:"id" url:"id"`
	// A description of the error that occurred.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *ErrorEventDetails) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *ErrorEventDetails) GetData() string {
	if e == nil {
		return ""
	}
	return e.Data
}

func (e *ErrorEventDetails) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e *ErrorEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if e.explicitFields != nil {
		next.Set(e.explicitFields)
	}
	next.Or(next, field)
	e.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ErrorEventDetails) SetID(id string) {
	e.ID = id
	e.require(errorEventDetailsFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ErrorEventDetails) SetData(data string) {
	e.Data = data
	e.require(errorEventDetailsFieldData)
}

func (e *ErrorEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler ErrorEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ErrorEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *ErrorEventDetails) MarshalJSON() ([]byte, error) {
	type embed ErrorEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, e.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (e *ErrorEventDetails) String() string {
	if e == nil {
		return "<nil>"
	}
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

var (
	genericErrorResponseFieldError = big.NewInt(1 << 0)
)

// Details about a generic error
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
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GenericErrorResponse) require(field *big.Int) {
	next := new(big.Int)
	if g.explicitFields != nil {
		next.Set(g.explicitFields)
	}
	next.Or(next, field)
	g.explicitFields = next
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
	messageEventFieldID   = big.NewInt(1 << 0)
	messageEventFieldData = big.NewInt(1 << 1)
)

// A server-sent event containing a message
type MessageEvent struct {
	ID   string            `json:"id" url:"id"`
	Data *MessageEventData `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MessageEvent) GetID() string {
	if m == nil {
		return ""
	}
	return m.ID
}

func (m *MessageEvent) GetData() *MessageEventData {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MessageEvent) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m *MessageEvent) require(field *big.Int) {
	next := new(big.Int)
	if m.explicitFields != nil {
		next.Set(m.explicitFields)
	}
	next.Or(next, field)
	m.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (m *MessageEvent) SetID(id string) {
	m.ID = id
	m.require(messageEventFieldID)
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
	if m == nil {
		return "<nil>"
	}
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

// Details about a message from the model
type MessageEventData struct {
	Type         string
	Task         *TaskEventDetails
	TaskResponse *TaskResponseEventDetails
	Message      *MessageEventDetails
	Thinking     *ThinkingEventDetails
	Response     *ResponseEventDetails
	Error        *ErrorEventDetails
	Cancelled    *CancelEventDetails

	rawJSON json.RawMessage
}

func (m *MessageEventData) GetType() string {
	if m == nil {
		return ""
	}
	return m.Type
}

func (m *MessageEventData) GetTask() *TaskEventDetails {
	if m == nil {
		return nil
	}
	return m.Task
}

func (m *MessageEventData) GetTaskResponse() *TaskResponseEventDetails {
	if m == nil {
		return nil
	}
	return m.TaskResponse
}

func (m *MessageEventData) GetMessage() *MessageEventDetails {
	if m == nil {
		return nil
	}
	return m.Message
}

func (m *MessageEventData) GetThinking() *ThinkingEventDetails {
	if m == nil {
		return nil
	}
	return m.Thinking
}

func (m *MessageEventData) GetResponse() *ResponseEventDetails {
	if m == nil {
		return nil
	}
	return m.Response
}

func (m *MessageEventData) GetError() *ErrorEventDetails {
	if m == nil {
		return nil
	}
	return m.Error
}

func (m *MessageEventData) GetCancelled() *CancelEventDetails {
	if m == nil {
		return nil
	}
	return m.Cancelled
}

func (m *MessageEventData) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	m.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", m)
	}
	switch unmarshaler.Type {
	case "task":
		value := new(TaskEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Task = value
	case "task_response":
		value := new(TaskResponseEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.TaskResponse = value
	case "message":
		value := new(MessageEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Message = value
	case "thinking":
		value := new(ThinkingEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Thinking = value
	case "response":
		value := new(ResponseEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Response = value
	case "error":
		value := new(ErrorEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Error = value
	case "cancelled":
		value := new(CancelEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Cancelled = value
	}
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m MessageEventData) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Task != nil {
		return internal.MarshalJSONWithExtraProperty(m.Task, "type", "task")
	}
	if m.TaskResponse != nil {
		return internal.MarshalJSONWithExtraProperty(m.TaskResponse, "type", "task_response")
	}
	if m.Message != nil {
		return internal.MarshalJSONWithExtraProperty(m.Message, "type", "message")
	}
	if m.Thinking != nil {
		return internal.MarshalJSONWithExtraProperty(m.Thinking, "type", "thinking")
	}
	if m.Response != nil {
		return internal.MarshalJSONWithExtraProperty(m.Response, "type", "response")
	}
	if m.Error != nil {
		return internal.MarshalJSONWithExtraProperty(m.Error, "type", "error")
	}
	if m.Cancelled != nil {
		return internal.MarshalJSONWithExtraProperty(m.Cancelled, "type", "cancelled")
	}
	if len(m.rawJSON) > 0 {
		return m.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type MessageEventDataVisitor interface {
	VisitTask(*TaskEventDetails) error
	VisitTaskResponse(*TaskResponseEventDetails) error
	VisitMessage(*MessageEventDetails) error
	VisitThinking(*ThinkingEventDetails) error
	VisitResponse(*ResponseEventDetails) error
	VisitError(*ErrorEventDetails) error
	VisitCancelled(*CancelEventDetails) error
}

func (m *MessageEventData) Accept(visitor MessageEventDataVisitor) error {
	if m.Task != nil {
		return visitor.VisitTask(m.Task)
	}
	if m.TaskResponse != nil {
		return visitor.VisitTaskResponse(m.TaskResponse)
	}
	if m.Message != nil {
		return visitor.VisitMessage(m.Message)
	}
	if m.Thinking != nil {
		return visitor.VisitThinking(m.Thinking)
	}
	if m.Response != nil {
		return visitor.VisitResponse(m.Response)
	}
	if m.Error != nil {
		return visitor.VisitError(m.Error)
	}
	if m.Cancelled != nil {
		return visitor.VisitCancelled(m.Cancelled)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *MessageEventData) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Task != nil {
		fields = append(fields, "task")
	}
	if m.TaskResponse != nil {
		fields = append(fields, "task_response")
	}
	if m.Message != nil {
		fields = append(fields, "message")
	}
	if m.Thinking != nil {
		fields = append(fields, "thinking")
	}
	if m.Response != nil {
		fields = append(fields, "response")
	}
	if m.Error != nil {
		fields = append(fields, "error")
	}
	if m.Cancelled != nil {
		fields = append(fields, "cancelled")
	}
	if len(fields) == 0 {
		if m.Type != "" {
			if len(m.rawJSON) > 0 {
				return nil
			}
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", m, m.Type)
		}
		return fmt.Errorf("type %T is empty", m)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", m, fields)
	}
	if m.Type != "" {
		field := fields[0]
		if m.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				m,
				m.Type,
				m,
			)
		}
	}
	return nil
}

var (
	messageEventDetailsFieldID   = big.NewInt(1 << 0)
	messageEventDetailsFieldData = big.NewInt(1 << 1)
)

// Details about a message from the model
type MessageEventDetails struct {
	// The ID of the message event
	ID string `json:"id" url:"id"`
	// A progress message from the model to be shown to the user.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MessageEventDetails) GetID() string {
	if m == nil {
		return ""
	}
	return m.ID
}

func (m *MessageEventDetails) GetData() string {
	if m == nil {
		return ""
	}
	return m.Data
}

func (m *MessageEventDetails) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m *MessageEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if m.explicitFields != nil {
		next.Set(m.explicitFields)
	}
	next.Or(next, field)
	m.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (m *MessageEventDetails) SetID(id string) {
	m.ID = id
	m.require(messageEventDetailsFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (m *MessageEventDetails) SetData(data string) {
	m.Data = data
	m.require(messageEventDetailsFieldData)
}

func (m *MessageEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler MessageEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MessageEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	m.rawJSON = json.RawMessage(data)
	return nil
}

func (m *MessageEventDetails) MarshalJSON() ([]byte, error) {
	type embed MessageEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, m.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (m *MessageEventDetails) String() string {
	if m == nil {
		return "<nil>"
	}
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

var (
	notificationChannelMessageEventFieldID   = big.NewInt(1 << 0)
	notificationChannelMessageEventFieldData = big.NewInt(1 << 1)
)

// A server-sent event containing one notification channel event.
type NotificationChannelMessageEvent struct {
	// The event's id (`rnco_…`), also sent as the SSE `id:` line; use it as `Last-Event-ID` to resume.
	ID   string                    `json:"id" url:"id"`
	Data *NotificationChannelEvent `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NotificationChannelMessageEvent) GetID() string {
	if n == nil {
		return ""
	}
	return n.ID
}

func (n *NotificationChannelMessageEvent) GetData() *NotificationChannelEvent {
	if n == nil {
		return nil
	}
	return n.Data
}

func (n *NotificationChannelMessageEvent) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n *NotificationChannelMessageEvent) require(field *big.Int) {
	next := new(big.Int)
	if n.explicitFields != nil {
		next.Set(n.explicitFields)
	}
	next.Or(next, field)
	n.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelMessageEvent) SetID(id string) {
	n.ID = id
	n.require(notificationChannelMessageEventFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelMessageEvent) SetData(data *NotificationChannelEvent) {
	n.Data = data
	n.require(notificationChannelMessageEventFieldData)
}

func (n *NotificationChannelMessageEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler NotificationChannelMessageEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NotificationChannelMessageEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NotificationChannelMessageEvent) MarshalJSON() ([]byte, error) {
	type embed NotificationChannelMessageEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, n.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (n *NotificationChannelMessageEvent) String() string {
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

type NotificationChannelSseEvent struct {
	Event   string
	Error   *ErrorEvent
	Message *NotificationChannelMessageEvent

	rawJSON json.RawMessage
}

func (n *NotificationChannelSseEvent) GetEvent() string {
	if n == nil {
		return ""
	}
	return n.Event
}

func (n *NotificationChannelSseEvent) GetError() *ErrorEvent {
	if n == nil {
		return nil
	}
	return n.Error
}

func (n *NotificationChannelSseEvent) GetMessage() *NotificationChannelMessageEvent {
	if n == nil {
		return nil
	}
	return n.Message
}

func (n *NotificationChannelSseEvent) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Event string `json:"event"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	n.Event = unmarshaler.Event
	if unmarshaler.Event == "" {
		return fmt.Errorf("%T did not include discriminant event", n)
	}
	switch unmarshaler.Event {
	case "error":
		value := new(ErrorEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		n.Error = value
	case "message":
		value := new(NotificationChannelMessageEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		n.Message = value
	}
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n NotificationChannelSseEvent) MarshalJSON() ([]byte, error) {
	if err := n.validate(); err != nil {
		return nil, err
	}
	if n.Error != nil {
		return internal.MarshalJSONWithExtraProperty(n.Error, "event", "error")
	}
	if n.Message != nil {
		return internal.MarshalJSONWithExtraProperty(n.Message, "event", "message")
	}
	if len(n.rawJSON) > 0 {
		return n.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", n)
}

type NotificationChannelSseEventVisitor interface {
	VisitError(*ErrorEvent) error
	VisitMessage(*NotificationChannelMessageEvent) error
}

func (n *NotificationChannelSseEvent) Accept(visitor NotificationChannelSseEventVisitor) error {
	if n.Error != nil {
		return visitor.VisitError(n.Error)
	}
	if n.Message != nil {
		return visitor.VisitMessage(n.Message)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", n)
}

func (n *NotificationChannelSseEvent) validate() error {
	if n == nil {
		return fmt.Errorf("type %T is nil", n)
	}
	var fields []string
	if n.Error != nil {
		fields = append(fields, "error")
	}
	if n.Message != nil {
		fields = append(fields, "message")
	}
	if len(fields) == 0 {
		if n.Event != "" {
			if len(n.rawJSON) > 0 {
				return nil
			}
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", n, n.Event)
		}
		return fmt.Errorf("type %T is empty", n)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", n, fields)
	}
	if n.Event != "" {
		field := fields[0]
		if n.Event != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				n,
				n.Event,
				n,
			)
		}
	}
	return nil
}

var (
	responseEventDetailsFieldID   = big.NewInt(1 << 0)
	responseEventDetailsFieldData = big.NewInt(1 << 1)
)

// Details about the final response from the model
type ResponseEventDetails struct {
	// The ID of the response event
	ID string `json:"id" url:"id"`
	// The final response from the model. The structure of this object will conform to one of the output message type schemas defined on the Intent.
	Data map[string]any `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ResponseEventDetails) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *ResponseEventDetails) GetData() map[string]any {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *ResponseEventDetails) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ResponseEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if r.explicitFields != nil {
		next.Set(r.explicitFields)
	}
	next.Or(next, field)
	r.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResponseEventDetails) SetID(id string) {
	r.ID = id
	r.require(responseEventDetailsFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResponseEventDetails) SetData(data map[string]any) {
	r.Data = data
	r.require(responseEventDetailsFieldData)
}

func (r *ResponseEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler ResponseEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResponseEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResponseEventDetails) MarshalJSON() ([]byte, error) {
	type embed ResponseEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ResponseEventDetails) String() string {
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
	serverErrorEventDetailsFieldType  = big.NewInt(1 << 0)
	serverErrorEventDetailsFieldError = big.NewInt(1 << 1)
)

// Details about an error that occurred during request processing
type ServerErrorEventDetails struct {
	Type ServerErrorEventDetailsType `json:"type" url:"type"`
	// A description of the error that occurred.
	Error string `json:"error" url:"error"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *ServerErrorEventDetails) GetType() ServerErrorEventDetailsType {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *ServerErrorEventDetails) GetError() string {
	if s == nil {
		return ""
	}
	return s.Error
}

func (s *ServerErrorEventDetails) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *ServerErrorEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if s.explicitFields != nil {
		next.Set(s.explicitFields)
	}
	next.Or(next, field)
	s.explicitFields = next
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *ServerErrorEventDetails) SetType(type_ ServerErrorEventDetailsType) {
	s.Type = type_
	s.require(serverErrorEventDetailsFieldType)
}

// SetError sets the Error field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *ServerErrorEventDetails) SetError(error_ string) {
	s.Error = error_
	s.require(serverErrorEventDetailsFieldError)
}

func (s *ServerErrorEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler ServerErrorEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ServerErrorEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *ServerErrorEventDetails) MarshalJSON() ([]byte, error) {
	type embed ServerErrorEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *ServerErrorEventDetails) String() string {
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

type ServerErrorEventDetailsType string

const (
	ServerErrorEventDetailsTypeServerError ServerErrorEventDetailsType = "server_error"
)

func NewServerErrorEventDetailsTypeFromString(s string) (ServerErrorEventDetailsType, error) {
	switch s {
	case "server_error":
		return ServerErrorEventDetailsTypeServerError, nil
	}
	var t ServerErrorEventDetailsType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s ServerErrorEventDetailsType) Ptr() *ServerErrorEventDetailsType {
	return &s
}

type SseEvent struct {
	Event   string
	Error   *ErrorEvent
	Message *MessageEvent

	rawJSON json.RawMessage
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
	s.rawJSON = json.RawMessage(data)
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
	if len(s.rawJSON) > 0 {
		return s.rawJSON, nil
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
			if len(s.rawJSON) > 0 {
				return nil
			}
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

var (
	taskDataFieldInput        = big.NewInt(1 << 0)
	taskDataFieldTaskType     = big.NewInt(1 << 1)
	taskDataFieldOutputSchema = big.NewInt(1 << 2)
)

// Details about the task the user must perform to continue processing the request
type TaskData struct {
	// Any information necessary to perform the task. Conforms to the input schema of the task definition named by `task_type`. For example, for the built-in `ask_question` task this is an object with a single `question` string property.
	Input map[string]any `json:"input" url:"input"`
	// The name of the task definition that produced this task. Built-in definitions are `ask_question`, `pick_one`, `multi_select` and `confirmation`; authors may define additional task definitions with their own names.
	TaskType string `json:"task_type" url:"task_type"`
	// A JSON Schema describing the object that must be submitted to complete this task. The schema is captured on the task when it is created, so it remains authoritative for this task even if the task definition is later modified or deleted.
	OutputSchema map[string]any `json:"output_schema" url:"output_schema"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskData) GetInput() map[string]any {
	if t == nil {
		return nil
	}
	return t.Input
}

func (t *TaskData) GetTaskType() string {
	if t == nil {
		return ""
	}
	return t.TaskType
}

func (t *TaskData) GetOutputSchema() map[string]any {
	if t == nil {
		return nil
	}
	return t.OutputSchema
}

func (t *TaskData) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *TaskData) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetInput sets the Input field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskData) SetInput(input map[string]any) {
	t.Input = input
	t.require(taskDataFieldInput)
}

// SetTaskType sets the TaskType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskData) SetTaskType(taskType string) {
	t.TaskType = taskType
	t.require(taskDataFieldTaskType)
}

// SetOutputSchema sets the OutputSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskData) SetOutputSchema(outputSchema map[string]any) {
	t.OutputSchema = outputSchema
	t.require(taskDataFieldOutputSchema)
}

func (t *TaskData) UnmarshalJSON(data []byte) error {
	type unmarshaler TaskData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TaskData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TaskData) MarshalJSON() ([]byte, error) {
	type embed TaskData
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *TaskData) String() string {
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
	taskEventDetailsFieldID   = big.NewInt(1 << 0)
	taskEventDetailsFieldData = big.NewInt(1 << 1)
)

// Details about a task the user must perform to continue processing the request
type TaskEventDetails struct {
	// The ID of the task event
	ID   string    `json:"id" url:"id"`
	Data *TaskData `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskEventDetails) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

func (t *TaskEventDetails) GetData() *TaskData {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TaskEventDetails) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *TaskEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEventDetails) SetID(id string) {
	t.ID = id
	t.require(taskEventDetailsFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEventDetails) SetData(data *TaskData) {
	t.Data = data
	t.require(taskEventDetailsFieldData)
}

func (t *TaskEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler TaskEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TaskEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TaskEventDetails) MarshalJSON() ([]byte, error) {
	type embed TaskEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *TaskEventDetails) String() string {
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
	taskResponseEventDetailsFieldID        = big.NewInt(1 << 0)
	taskResponseEventDetailsFieldData      = big.NewInt(1 << 1)
	taskResponseEventDetailsFieldInReplyTo = big.NewInt(1 << 2)
)

// Details about a task the user completed
type TaskResponseEventDetails struct {
	// The ID of the task response event
	ID string `json:"id" url:"id"`
	// The response submitted to complete the task. An object conforming to the `output_schema` of the task event referenced by `in_reply_to` (e.g. `{"answer": "..."}` for the built-in `ask_question` task).
	Data map[string]any `json:"data" url:"data"`
	// The ID of the task event this is a response to.
	InReplyTo string `json:"in_reply_to" url:"in_reply_to"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskResponseEventDetails) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

func (t *TaskResponseEventDetails) GetData() map[string]any {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TaskResponseEventDetails) GetInReplyTo() string {
	if t == nil {
		return ""
	}
	return t.InReplyTo
}

func (t *TaskResponseEventDetails) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *TaskResponseEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEventDetails) SetID(id string) {
	t.ID = id
	t.require(taskResponseEventDetailsFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEventDetails) SetData(data map[string]any) {
	t.Data = data
	t.require(taskResponseEventDetailsFieldData)
}

// SetInReplyTo sets the InReplyTo field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEventDetails) SetInReplyTo(inReplyTo string) {
	t.InReplyTo = inReplyTo
	t.require(taskResponseEventDetailsFieldInReplyTo)
}

func (t *TaskResponseEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler TaskResponseEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TaskResponseEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TaskResponseEventDetails) MarshalJSON() ([]byte, error) {
	type embed TaskResponseEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *TaskResponseEventDetails) String() string {
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
	thinkingEventDetailsFieldID   = big.NewInt(1 << 0)
	thinkingEventDetailsFieldData = big.NewInt(1 << 1)
)

// Details about a thinking event from the model
type ThinkingEventDetails struct {
	// The ID of the thinking event
	ID string `json:"id" url:"id"`
	// A description of the model's current thinking process.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *ThinkingEventDetails) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

func (t *ThinkingEventDetails) GetData() string {
	if t == nil {
		return ""
	}
	return t.Data
}

func (t *ThinkingEventDetails) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *ThinkingEventDetails) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ThinkingEventDetails) SetID(id string) {
	t.ID = id
	t.require(thinkingEventDetailsFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ThinkingEventDetails) SetData(data string) {
	t.Data = data
	t.require(thinkingEventDetailsFieldData)
}

func (t *ThinkingEventDetails) UnmarshalJSON(data []byte) error {
	type unmarshaler ThinkingEventDetails
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThinkingEventDetails(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThinkingEventDetails) MarshalJSON() ([]byte, error) {
	type embed ThinkingEventDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *ThinkingEventDetails) String() string {
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
