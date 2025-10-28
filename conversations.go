// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/partite-ai/sdk-go/internal"
	io "io"
	big "math/big"
)

var (
	attachAssetsToConversationRequestFieldConversationId = big.NewInt(1 << 0)
)

type AttachAssetsToConversationRequest struct {
	// The ID of an existing conversation
	ConversationId string      `json:"-" url:"-"`
	Attachments    []io.Reader `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (a *AttachAssetsToConversationRequest) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AttachAssetsToConversationRequest) SetConversationId(conversationId string) {
	a.ConversationId = conversationId
	a.require(attachAssetsToConversationRequestFieldConversationId)
}

var (
	cancelRequestRequestFieldConversationId = big.NewInt(1 << 0)
	cancelRequestRequestFieldRequestId      = big.NewInt(1 << 1)
)

type CancelRequestRequest struct {
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`
	// The ID of an existing request
	RequestId string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CancelRequestRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelRequestRequest) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
	c.require(cancelRequestRequestFieldConversationId)
}

// SetRequestId sets the RequestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelRequestRequest) SetRequestId(requestId string) {
	c.RequestId = requestId
	c.require(cancelRequestRequestFieldRequestId)
}

var (
	completeTaskRequestFieldConversationId = big.NewInt(1 << 0)
	completeTaskRequestFieldRequestId      = big.NewInt(1 << 1)
	completeTaskRequestFieldTaskId         = big.NewInt(1 << 2)
)

type CompleteTaskRequest struct {
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`
	// The ID of an existing request
	RequestId string `json:"-" url:"-"`
	// The ID of the task to complete
	TaskId string      `json:"-" url:"-"`
	Body   interface{} `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CompleteTaskRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CompleteTaskRequest) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
	c.require(completeTaskRequestFieldConversationId)
}

// SetRequestId sets the RequestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CompleteTaskRequest) SetRequestId(requestId string) {
	c.RequestId = requestId
	c.require(completeTaskRequestFieldRequestId)
}

// SetTaskId sets the TaskId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CompleteTaskRequest) SetTaskId(taskId string) {
	c.TaskId = taskId
	c.require(completeTaskRequestFieldTaskId)
}

func (c *CompleteTaskRequest) UnmarshalJSON(data []byte) error {
	var body interface{}
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
	createConversationRequestRequestFieldConversationId = big.NewInt(1 << 0)
	createConversationRequestRequestFieldIntentName     = big.NewInt(1 << 1)
)

type CreateConversationRequestRequest struct {
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`
	// The name of the Intent defined on the target mesh which will be used for this request
	IntentName  string      `json:"-" url:"-"`
	Attachments []io.Reader `json:"-" url:"-"`
	// The input message to provide to the mesh. The structure of this object must conform to the input message type schema defined on the Intent. Within the input message, any asset references may use the syntax `attachment://<filename>` format to refer to assets attached in the same multipart/form-data body.
	InputMessage map[string]interface{} `json:"input_message,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateConversationRequestRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationRequestRequest) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
	c.require(createConversationRequestRequestFieldConversationId)
}

// SetIntentName sets the IntentName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationRequestRequest) SetIntentName(intentName string) {
	c.IntentName = intentName
	c.require(createConversationRequestRequestFieldIntentName)
}

var (
	createConversationWithInitialRequestRequestFieldIntentName     = big.NewInt(1 << 0)
	createConversationWithInitialRequestRequestFieldMeshId         = big.NewInt(1 << 1)
	createConversationWithInitialRequestRequestFieldVersionOrdinal = big.NewInt(1 << 2)
	createConversationWithInitialRequestRequestFieldVersionLabel   = big.NewInt(1 << 3)
)

type CreateConversationWithInitialRequestRequest struct {
	// The name of the Intent defined on the target mesh which will be used to start the conversation
	IntentName string `json:"-" url:"-"`
	// The ID of the Mesh for which to create a conversation
	MeshId string `json:"-" url:"mesh_id"`
	// The version number of a mesh version for which to create the conversation
	VersionOrdinal *int `json:"-" url:"version_ordinal,omitempty"`
	// The version label of a mesh version for which to create the conversation
	VersionLabel *string     `json:"-" url:"version_label,omitempty"`
	Attachments  []io.Reader `json:"-" url:"-"`
	// The input message to provide to the mesh. The structure of this object must conform to the input message type schema defined on the Intent. Within the input message, any asset references may use the syntax `attachment://<filename>` format to refer to assets attached in the same multipart/form-data body.
	InputMessage map[string]interface{} `json:"input_message,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateConversationWithInitialRequestRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetIntentName sets the IntentName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequestRequest) SetIntentName(intentName string) {
	c.IntentName = intentName
	c.require(createConversationWithInitialRequestRequestFieldIntentName)
}

// SetMeshId sets the MeshId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationWithInitialRequestRequest) SetMeshId(meshId string) {
	c.MeshId = meshId
	c.require(createConversationWithInitialRequestRequestFieldMeshId)
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

var (
	createEmptyConversationRequestFieldMeshId         = big.NewInt(1 << 0)
	createEmptyConversationRequestFieldVersionOrdinal = big.NewInt(1 << 1)
	createEmptyConversationRequestFieldVersionLabel   = big.NewInt(1 << 2)
)

type CreateEmptyConversationRequest struct {
	// The ID of the Mesh for which to create a conversation
	MeshId string `json:"-" url:"mesh_id"`
	// The version number of a mesh version for which to create the conversation
	VersionOrdinal *int `json:"-" url:"version_ordinal,omitempty"`
	// The version label of a mesh version for which to create the conversation
	VersionLabel *string `json:"-" url:"version_label,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateEmptyConversationRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetMeshId sets the MeshId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateEmptyConversationRequest) SetMeshId(meshId string) {
	c.MeshId = meshId
	c.require(createEmptyConversationRequestFieldMeshId)
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

var (
	getConversationRequestsRequestFieldConversationId = big.NewInt(1 << 0)
)

type GetConversationRequestsRequest struct {
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetConversationRequestsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsRequest) SetConversationId(conversationId string) {
	g.ConversationId = conversationId
	g.require(getConversationRequestsRequestFieldConversationId)
}

var (
	getIntentsForConversationRequestFieldConversationId = big.NewInt(1 << 0)
)

type GetIntentsForConversationRequest struct {
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetIntentsForConversationRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIntentsForConversationRequest) SetConversationId(conversationId string) {
	g.ConversationId = conversationId
	g.require(getIntentsForConversationRequestFieldConversationId)
}

var (
	getIntentsForNewConversationRequestFieldMeshId         = big.NewInt(1 << 0)
	getIntentsForNewConversationRequestFieldVersionOrdinal = big.NewInt(1 << 1)
	getIntentsForNewConversationRequestFieldVersionLabel   = big.NewInt(1 << 2)
)

type GetIntentsForNewConversationRequest struct {
	// The ID of the Mesh for which to get intents
	MeshId string `json:"-" url:"mesh_id"`
	// The version number of a mesh version for which to get intents
	VersionOrdinal *int `json:"-" url:"version_ordinal,omitempty"`
	// The version label of a mesh version for which to get intents
	VersionLabel *string `json:"-" url:"version_label,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetIntentsForNewConversationRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetMeshId sets the MeshId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetIntentsForNewConversationRequest) SetMeshId(meshId string) {
	g.MeshId = meshId
	g.require(getIntentsForNewConversationRequestFieldMeshId)
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
	getRequestEventsRequestFieldLastEventId    = big.NewInt(1 << 0)
	getRequestEventsRequestFieldConversationId = big.NewInt(1 << 1)
	getRequestEventsRequestFieldRequestId      = big.NewInt(1 << 2)
)

type GetRequestEventsRequest struct {
	// The ID of the last event received. If provided, the stream will start from the next event after this ID.
	LastEventId *string `json:"-" url:"-"`
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`
	// The ID of an existing request
	RequestId string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetRequestEventsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetLastEventId sets the LastEventId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestEventsRequest) SetLastEventId(lastEventId *string) {
	g.LastEventId = lastEventId
	g.require(getRequestEventsRequestFieldLastEventId)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestEventsRequest) SetConversationId(conversationId string) {
	g.ConversationId = conversationId
	g.require(getRequestEventsRequestFieldConversationId)
}

// SetRequestId sets the RequestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestEventsRequest) SetRequestId(requestId string) {
	g.RequestId = requestId
	g.require(getRequestEventsRequestFieldRequestId)
}

var (
	getRequestResponseRequestFieldConversationId = big.NewInt(1 << 0)
	getRequestResponseRequestFieldRequestId      = big.NewInt(1 << 1)
)

type GetRequestResponseRequest struct {
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`
	// The ID of an existing request
	RequestId string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetRequestResponseRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestResponseRequest) SetConversationId(conversationId string) {
	g.ConversationId = conversationId
	g.require(getRequestResponseRequestFieldConversationId)
}

// SetRequestId sets the RequestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestResponseRequest) SetRequestId(requestId string) {
	g.RequestId = requestId
	g.require(getRequestResponseRequestFieldRequestId)
}

var (
	getRequestTaskRequestFieldConversationId = big.NewInt(1 << 0)
	getRequestTaskRequestFieldRequestId      = big.NewInt(1 << 1)
)

type GetRequestTaskRequest struct {
	// The ID of an existing conversation
	ConversationId string `json:"-" url:"-"`
	// The ID of an existing request
	RequestId string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetRequestTaskRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestTaskRequest) SetConversationId(conversationId string) {
	g.ConversationId = conversationId
	g.require(getRequestTaskRequestFieldConversationId)
}

// SetRequestId sets the RequestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetRequestTaskRequest) SetRequestId(requestId string) {
	g.RequestId = requestId
	g.require(getRequestTaskRequestFieldRequestId)
}

// Information about an uploaded attachment
var (
	attachmentInfoFieldReferenceUrl = big.NewInt(1 << 0)
	attachmentInfoFieldFilename     = big.NewInt(1 << 1)
)

type AttachmentInfo struct {
	// A URL that can be used to access the uploaded asset in references.
	ReferenceUrl string `json:"reference_url" url:"reference_url"`
	// The original filename of the uploaded asset.
	Filename string `json:"filename" url:"filename"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AttachmentInfo) GetReferenceUrl() string {
	if a == nil {
		return ""
	}
	return a.ReferenceUrl
}

func (a *AttachmentInfo) GetFilename() string {
	if a == nil {
		return ""
	}
	return a.Filename
}

func (a *AttachmentInfo) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AttachmentInfo) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetReferenceUrl sets the ReferenceUrl field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AttachmentInfo) SetReferenceUrl(referenceUrl string) {
	a.ReferenceUrl = referenceUrl
	a.require(attachmentInfoFieldReferenceUrl)
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

// Details about a cancellation event
var (
	cancelEventDetailsFieldId   = big.NewInt(1 << 0)
	cancelEventDetailsFieldData = big.NewInt(1 << 1)
)

type CancelEventDetails struct {
	// The ID of the cancellation event
	Id string `json:"id" url:"id"`
	// A description of the cancellation.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CancelEventDetails) GetId() string {
	if c == nil {
		return ""
	}
	return c.Id
}

func (c *CancelEventDetails) GetData() string {
	if c == nil {
		return ""
	}
	return c.Data
}

func (c *CancelEventDetails) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CancelEventDetails) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelEventDetails) SetId(id string) {
	c.Id = id
	c.require(cancelEventDetailsFieldId)
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

// Response returned when a conversation is created without an initial request.
var (
	createConversationResponseFieldConversationId = big.NewInt(1 << 0)
	createConversationResponseFieldIntents        = big.NewInt(1 << 1)
)

type CreateConversationResponse struct {
	ConversationId string      `json:"conversation_id" url:"conversation_id"`
	Intents        IntentsList `json:"intents" url:"intents"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateConversationResponse) GetConversationId() string {
	if c == nil {
		return ""
	}
	return c.ConversationId
}

func (c *CreateConversationResponse) GetIntents() IntentsList {
	if c == nil {
		return nil
	}
	return c.Intents
}

func (c *CreateConversationResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateConversationResponse) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConversationResponse) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
	c.require(createConversationResponseFieldConversationId)
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
	createRequestResponseFieldConversationId = big.NewInt(1 << 0)
	createRequestResponseFieldRequestId      = big.NewInt(1 << 1)
)

type CreateRequestResponse struct {
	ConversationId string `json:"conversation_id" url:"conversation_id"`
	RequestId      string `json:"request_id" url:"request_id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateRequestResponse) GetConversationId() string {
	if c == nil {
		return ""
	}
	return c.ConversationId
}

func (c *CreateRequestResponse) GetRequestId() string {
	if c == nil {
		return ""
	}
	return c.RequestId
}

func (c *CreateRequestResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateRequestResponse) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestResponse) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
	c.require(createRequestResponseFieldConversationId)
}

// SetRequestId sets the RequestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateRequestResponse) SetRequestId(requestId string) {
	c.RequestId = requestId
	c.require(createRequestResponseFieldRequestId)
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

// Details about an error that occurred during request processing
var (
	errorEventDetailsFieldId   = big.NewInt(1 << 0)
	errorEventDetailsFieldData = big.NewInt(1 << 1)
)

type ErrorEventDetails struct {
	// The ID of the error event
	Id string `json:"id" url:"id"`
	// A description of the error that occurred.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *ErrorEventDetails) GetId() string {
	if e == nil {
		return ""
	}
	return e.Id
}

func (e *ErrorEventDetails) GetData() string {
	if e == nil {
		return ""
	}
	return e.Data
}

func (e *ErrorEventDetails) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *ErrorEventDetails) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ErrorEventDetails) SetId(id string) {
	e.Id = id
	e.require(errorEventDetailsFieldId)
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
	getTaskResponseFieldTaskId = big.NewInt(1 << 0)
	getTaskResponseFieldData   = big.NewInt(1 << 1)
)

type GetTaskResponse struct {
	TaskId string `json:"task_id" url:"task_id"`
	// Details about the task the user must perform to continue processing the request
	Data *GetTaskResponseData `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetTaskResponse) GetTaskId() string {
	if g == nil {
		return ""
	}
	return g.TaskId
}

func (g *GetTaskResponse) GetData() *GetTaskResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTaskResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetTaskResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetTaskId sets the TaskId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTaskResponse) SetTaskId(taskId string) {
	g.TaskId = taskId
	g.require(getTaskResponseFieldTaskId)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTaskResponse) SetData(data *GetTaskResponseData) {
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

// Details about the task the user must perform to continue processing the request
var (
	getTaskResponseDataFieldInput    = big.NewInt(1 << 0)
	getTaskResponseDataFieldTaskType = big.NewInt(1 << 1)
)

type GetTaskResponseData struct {
	// Any information necessary to perform the task. Varies according to `task_type`. For the `ask_question` task type (the only currently supported type), will consist of an object with a single `question` property containing the question the model desires the answer for. In the future, custom task definitions will be allowed, and the object structure will conform to the schema defined in the task type.
	Input map[string]interface{} `json:"input" url:"input"`
	// Currently will be "ask_question", additional task types will be defined in the future.
	TaskType string `json:"task_type" url:"task_type"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetTaskResponseData) GetInput() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.Input
}

func (g *GetTaskResponseData) GetTaskType() string {
	if g == nil {
		return ""
	}
	return g.TaskType
}

func (g *GetTaskResponseData) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetTaskResponseData) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetInput sets the Input field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTaskResponseData) SetInput(input map[string]interface{}) {
	g.Input = input
	g.require(getTaskResponseDataFieldInput)
}

// SetTaskType sets the TaskType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTaskResponseData) SetTaskType(taskType string) {
	g.TaskType = taskType
	g.require(getTaskResponseDataFieldTaskType)
}

func (g *GetTaskResponseData) UnmarshalJSON(data []byte) error {
	type unmarshaler GetTaskResponseData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetTaskResponseData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetTaskResponseData) MarshalJSON() ([]byte, error) {
	type embed GetTaskResponseData
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetTaskResponseData) String() string {
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
	Schema      map[string]interface{}          `json:"schema" url:"schema"`
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

func (i *IntentsListItem) GetSchema() map[string]interface{} {
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
	return i.extraProperties
}

func (i *IntentsListItem) require(field *big.Int) {
	if i.explicitFields == nil {
		i.explicitFields = big.NewInt(0)
	}
	i.explicitFields.Or(i.explicitFields, field)
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
func (i *IntentsListItem) SetSchema(schema map[string]interface{}) {
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
	Name        *string                `json:"name,omitempty" url:"name,omitempty"`
	Description *string                `json:"description,omitempty" url:"description,omitempty"`
	Schema      map[string]interface{} `json:"schema,omitempty" url:"schema,omitempty"`

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

func (i *IntentsListItemResponsesItem) GetSchema() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.Schema
}

func (i *IntentsListItemResponsesItem) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *IntentsListItemResponsesItem) require(field *big.Int) {
	if i.explicitFields == nil {
		i.explicitFields = big.NewInt(0)
	}
	i.explicitFields.Or(i.explicitFields, field)
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
func (i *IntentsListItemResponsesItem) SetSchema(schema map[string]interface{}) {
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

// Response returned when listing intents for a mesh version.
var (
	listIntentsResponseFieldIntents = big.NewInt(1 << 0)
)

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
	return l.extraProperties
}

func (l *ListIntentsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
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

// Details about a message from the model
type MessageEventData struct {
	Type         string
	Task         *TaskEventDetails
	TaskResponse *TaskResponseEventDetails
	Message      *MessageEventDetails
	Thinking     *ThinkingEventDetails
	Response     *ResponseEventDetails
	Error        *ErrorEventDetails
	Cancel       *CancelEventDetails
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

func (m *MessageEventData) GetCancel() *CancelEventDetails {
	if m == nil {
		return nil
	}
	return m.Cancel
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
	case "cancel":
		value := new(CancelEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Cancel = value
	}
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
	if m.Cancel != nil {
		return internal.MarshalJSONWithExtraProperty(m.Cancel, "type", "cancel")
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
	VisitCancel(*CancelEventDetails) error
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
	if m.Cancel != nil {
		return visitor.VisitCancel(m.Cancel)
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
	if m.Cancel != nil {
		fields = append(fields, "cancel")
	}
	if len(fields) == 0 {
		if m.Type != "" {
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

// Details about a message from the model
var (
	messageEventDetailsFieldId   = big.NewInt(1 << 0)
	messageEventDetailsFieldData = big.NewInt(1 << 1)
)

type MessageEventDetails struct {
	// The ID of the message event
	Id string `json:"id" url:"id"`
	// A progress message from the model to be shown to the user.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (m *MessageEventDetails) GetId() string {
	if m == nil {
		return ""
	}
	return m.Id
}

func (m *MessageEventDetails) GetData() string {
	if m == nil {
		return ""
	}
	return m.Data
}

func (m *MessageEventDetails) GetExtraProperties() map[string]interface{} {
	return m.extraProperties
}

func (m *MessageEventDetails) require(field *big.Int) {
	if m.explicitFields == nil {
		m.explicitFields = big.NewInt(0)
	}
	m.explicitFields.Or(m.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (m *MessageEventDetails) SetId(id string) {
	m.Id = id
	m.require(messageEventDetailsFieldId)
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

// A chunk of data in the event stream
type ResponseEventData struct {
	Type         string
	Task         *TaskEventDetails
	TaskResponse *TaskResponseEventDetails
	Message      *MessageEventDetails
	Thinking     *ThinkingEventDetails
	Response     *ResponseEventDetails
	Error        *ErrorEventDetails
	Cancel       *CancelEventDetails
	ServerError  *ServerErrorEventDetails
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

func (r *ResponseEventData) GetCancel() *CancelEventDetails {
	if r == nil {
		return nil
	}
	return r.Cancel
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
	case "cancel":
		value := new(CancelEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Cancel = value
	case "server_error":
		value := new(ServerErrorEventDetails)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.ServerError = value
	}
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
	if r.Cancel != nil {
		return internal.MarshalJSONWithExtraProperty(r.Cancel, "type", "cancel")
	}
	if r.ServerError != nil {
		return internal.MarshalJSONWithExtraProperty(r.ServerError, "type", "server_error")
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
	VisitCancel(*CancelEventDetails) error
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
	if r.Cancel != nil {
		return visitor.VisitCancel(r.Cancel)
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
	if r.Cancel != nil {
		fields = append(fields, "cancel")
	}
	if r.ServerError != nil {
		fields = append(fields, "server_error")
	}
	if len(fields) == 0 {
		if r.Type != "" {
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

// Details about the final response from the model
var (
	responseEventDetailsFieldId   = big.NewInt(1 << 0)
	responseEventDetailsFieldData = big.NewInt(1 << 1)
)

type ResponseEventDetails struct {
	// The ID of the response event
	Id string `json:"id" url:"id"`
	// The final response from the model. The structure of this object will conform to one of the output message type schemas defined on the Intent.
	Data map[string]interface{} `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ResponseEventDetails) GetId() string {
	if r == nil {
		return ""
	}
	return r.Id
}

func (r *ResponseEventDetails) GetData() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *ResponseEventDetails) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResponseEventDetails) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResponseEventDetails) SetId(id string) {
	r.Id = id
	r.require(responseEventDetailsFieldId)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResponseEventDetails) SetData(data map[string]interface{}) {
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

// Details about an error that occurred during request processing
var (
	serverErrorEventDetailsFieldError = big.NewInt(1 << 0)
)

type ServerErrorEventDetails struct {
	// A description of the error that occurred.
	Error string `json:"error" url:"error"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *ServerErrorEventDetails) GetError() string {
	if s == nil {
		return ""
	}
	return s.Error
}

func (s *ServerErrorEventDetails) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *ServerErrorEventDetails) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
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

// Details about a task the user must perform to continue processing the request
var (
	taskEventDetailsFieldId   = big.NewInt(1 << 0)
	taskEventDetailsFieldData = big.NewInt(1 << 1)
)

type TaskEventDetails struct {
	// The ID of the task event
	Id   string                `json:"id" url:"id"`
	Data *TaskEventDetailsData `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskEventDetails) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *TaskEventDetails) GetData() *TaskEventDetailsData {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TaskEventDetails) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TaskEventDetails) require(field *big.Int) {
	if t.explicitFields == nil {
		t.explicitFields = big.NewInt(0)
	}
	t.explicitFields.Or(t.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEventDetails) SetId(id string) {
	t.Id = id
	t.require(taskEventDetailsFieldId)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEventDetails) SetData(data *TaskEventDetailsData) {
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
	taskEventDetailsDataFieldInput    = big.NewInt(1 << 0)
	taskEventDetailsDataFieldTaskType = big.NewInt(1 << 1)
)

type TaskEventDetailsData struct {
	// Any information necessary to perform the task. Varies according to `task_type`. For the `ask_question` task type (the only currently supported type), will consist of an object with a single `question` property containing the question the model desires the answer for. In the future, custom task definitions will be allowed, and the object structure will conform to the schema defined in the task type.
	Input map[string]interface{} `json:"input" url:"input"`
	// Currently will be "ask_question", additional task types will be defined in the future.
	TaskType string `json:"task_type" url:"task_type"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskEventDetailsData) GetInput() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.Input
}

func (t *TaskEventDetailsData) GetTaskType() string {
	if t == nil {
		return ""
	}
	return t.TaskType
}

func (t *TaskEventDetailsData) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TaskEventDetailsData) require(field *big.Int) {
	if t.explicitFields == nil {
		t.explicitFields = big.NewInt(0)
	}
	t.explicitFields.Or(t.explicitFields, field)
}

// SetInput sets the Input field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEventDetailsData) SetInput(input map[string]interface{}) {
	t.Input = input
	t.require(taskEventDetailsDataFieldInput)
}

// SetTaskType sets the TaskType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEventDetailsData) SetTaskType(taskType string) {
	t.TaskType = taskType
	t.require(taskEventDetailsDataFieldTaskType)
}

func (t *TaskEventDetailsData) UnmarshalJSON(data []byte) error {
	type unmarshaler TaskEventDetailsData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TaskEventDetailsData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TaskEventDetailsData) MarshalJSON() ([]byte, error) {
	type embed TaskEventDetailsData
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *TaskEventDetailsData) String() string {
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

// Details about a task the user completed
var (
	taskResponseEventDetailsFieldId        = big.NewInt(1 << 0)
	taskResponseEventDetailsFieldData      = big.NewInt(1 << 1)
	taskResponseEventDetailsFieldInReplyTo = big.NewInt(1 << 2)
)

type TaskResponseEventDetails struct {
	// The ID of the task response event
	Id   string      `json:"id" url:"id"`
	Data interface{} `json:"data" url:"data"`
	// The ID of the task event this is a response to.
	InReplyTo string `json:"in_reply_to" url:"in_reply_to"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskResponseEventDetails) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *TaskResponseEventDetails) GetData() interface{} {
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
	return t.extraProperties
}

func (t *TaskResponseEventDetails) require(field *big.Int) {
	if t.explicitFields == nil {
		t.explicitFields = big.NewInt(0)
	}
	t.explicitFields.Or(t.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEventDetails) SetId(id string) {
	t.Id = id
	t.require(taskResponseEventDetailsFieldId)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEventDetails) SetData(data interface{}) {
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

// Details about a thinking event from the model
var (
	thinkingEventDetailsFieldId   = big.NewInt(1 << 0)
	thinkingEventDetailsFieldData = big.NewInt(1 << 1)
)

type ThinkingEventDetails struct {
	// The ID of the thinking event
	Id string `json:"id" url:"id"`
	// A description of the model's current thinking process.
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *ThinkingEventDetails) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *ThinkingEventDetails) GetData() string {
	if t == nil {
		return ""
	}
	return t.Data
}

func (t *ThinkingEventDetails) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *ThinkingEventDetails) require(field *big.Int) {
	if t.explicitFields == nil {
		t.explicitFields = big.NewInt(0)
	}
	t.explicitFields.Or(t.explicitFields, field)
}

// SetId sets the Id field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ThinkingEventDetails) SetId(id string) {
	t.Id = id
	t.require(thinkingEventDetailsFieldId)
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
	return a.extraProperties
}

func (a *AttachAssetsToConversationResponse) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
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
	return g.extraProperties
}

func (g *GetConversationRequestsResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
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
	getConversationRequestsResponseRequestsItemFieldRequestId    = big.NewInt(1 << 0)
	getConversationRequestsResponseRequestsItemFieldInputMessage = big.NewInt(1 << 1)
	getConversationRequestsResponseRequestsItemFieldStatus       = big.NewInt(1 << 2)
	getConversationRequestsResponseRequestsItemFieldEvents       = big.NewInt(1 << 3)
)

type GetConversationRequestsResponseRequestsItem struct {
	RequestId    string                                            `json:"request_id" url:"request_id"`
	InputMessage map[string]interface{}                            `json:"input_message" url:"input_message"`
	Status       GetConversationRequestsResponseRequestsItemStatus `json:"status" url:"status"`
	Events       []*MessageEventData                               `json:"events" url:"events"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetConversationRequestsResponseRequestsItem) GetRequestId() string {
	if g == nil {
		return ""
	}
	return g.RequestId
}

func (g *GetConversationRequestsResponseRequestsItem) GetInputMessage() map[string]interface{} {
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

func (g *GetConversationRequestsResponseRequestsItem) GetEvents() []*MessageEventData {
	if g == nil {
		return nil
	}
	return g.Events
}

func (g *GetConversationRequestsResponseRequestsItem) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetConversationRequestsResponseRequestsItem) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetRequestId sets the RequestId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetRequestId(requestId string) {
	g.RequestId = requestId
	g.require(getConversationRequestsResponseRequestsItemFieldRequestId)
}

// SetInputMessage sets the InputMessage field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetInputMessage(inputMessage map[string]interface{}) {
	g.InputMessage = inputMessage
	g.require(getConversationRequestsResponseRequestsItemFieldInputMessage)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationRequestsResponseRequestsItem) SetStatus(status GetConversationRequestsResponseRequestsItemStatus) {
	g.Status = status
	g.require(getConversationRequestsResponseRequestsItemFieldStatus)
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
