// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/partite-ai/sdk-go/internal"
	big "math/big"
)

var (
	a2UIPresentationFieldContent = big.NewInt(1 << 0)
)

// An A2UI v0.9 surface for each state of the task.
type A2UIPresentation struct {
	Content *A2UISurfaces `json:"content" url:"content"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *A2UIPresentation) GetContent() *A2UISurfaces {
	if a == nil {
		return nil
	}
	return a.Content
}

func (a *A2UIPresentation) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *A2UIPresentation) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetContent sets the Content field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *A2UIPresentation) SetContent(content *A2UISurfaces) {
	a.Content = content
	a.require(a2UIPresentationFieldContent)
}

func (a *A2UIPresentation) UnmarshalJSON(data []byte) error {
	type unmarshaler A2UIPresentation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = A2UIPresentation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *A2UIPresentation) MarshalJSON() ([]byte, error) {
	type embed A2UIPresentation
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *A2UIPresentation) String() string {
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
	a2UISurfaceComponentsFieldComponents = big.NewInt(1 << 0)
)

type A2UISurfaceComponents struct {
	// The surface's A2UI v0.9 components. The entry component has id `root`.
	Components []map[string]any `json:"components" url:"components"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *A2UISurfaceComponents) GetComponents() []map[string]any {
	if a == nil {
		return nil
	}
	return a.Components
}

func (a *A2UISurfaceComponents) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *A2UISurfaceComponents) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetComponents sets the Components field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *A2UISurfaceComponents) SetComponents(components []map[string]any) {
	a.Components = components
	a.require(a2UISurfaceComponentsFieldComponents)
}

func (a *A2UISurfaceComponents) UnmarshalJSON(data []byte) error {
	type unmarshaler A2UISurfaceComponents
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = A2UISurfaceComponents(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *A2UISurfaceComponents) MarshalJSON() ([]byte, error) {
	type embed A2UISurfaceComponents
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *A2UISurfaceComponents) String() string {
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
	a2UISurfacesFieldCatalogID = big.NewInt(1 << 0)
	a2UISurfacesFieldOpen      = big.NewInt(1 << 1)
	a2UISurfacesFieldCompleted = big.NewInt(1 << 2)
)

// An A2UI v0.9 surface for each state of the task. While the task is open, a client renders `open`; once it is completed, `completed`, or no surface when there is none; once it is cancelled, no surface.
type A2UISurfaces struct {
	CatalogID string                 `json:"catalogId" url:"catalogId"`
	Open      *A2UISurfaceComponents `json:"open" url:"open"`
	Completed *A2UISurfaceComponents `json:"completed,omitempty" url:"completed,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *A2UISurfaces) GetCatalogID() string {
	if a == nil {
		return ""
	}
	return a.CatalogID
}

func (a *A2UISurfaces) GetOpen() *A2UISurfaceComponents {
	if a == nil {
		return nil
	}
	return a.Open
}

func (a *A2UISurfaces) GetCompleted() *A2UISurfaceComponents {
	if a == nil {
		return nil
	}
	return a.Completed
}

func (a *A2UISurfaces) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *A2UISurfaces) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetCatalogID sets the CatalogID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *A2UISurfaces) SetCatalogID(catalogID string) {
	a.CatalogID = catalogID
	a.require(a2UISurfacesFieldCatalogID)
}

// SetOpen sets the Open field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *A2UISurfaces) SetOpen(open *A2UISurfaceComponents) {
	a.Open = open
	a.require(a2UISurfacesFieldOpen)
}

// SetCompleted sets the Completed field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *A2UISurfaces) SetCompleted(completed *A2UISurfaceComponents) {
	a.Completed = completed
	a.require(a2UISurfacesFieldCompleted)
}

func (a *A2UISurfaces) UnmarshalJSON(data []byte) error {
	type unmarshaler A2UISurfaces
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = A2UISurfaces(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *A2UISurfaces) MarshalJSON() ([]byte, error) {
	type embed A2UISurfaces
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *A2UISurfaces) String() string {
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
	approvalToolFieldName   = big.NewInt(1 << 0)
	approvalToolFieldSource = big.NewInt(1 << 1)
	approvalToolFieldMemory = big.NewInt(1 << 2)
)

// The tool a call is for.
type ApprovalTool struct {
	// The tool's name, as the agent sees it.
	Name   string              `json:"name" url:"name"`
	Source *ApprovalToolSource `json:"source" url:"source"`
	// What only a memory tool has, set only when `source.kind` is `memory`.
	Memory *ApprovalToolMemory `json:"memory,omitempty" url:"memory,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalTool) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

func (a *ApprovalTool) GetSource() *ApprovalToolSource {
	if a == nil {
		return nil
	}
	return a.Source
}

func (a *ApprovalTool) GetMemory() *ApprovalToolMemory {
	if a == nil {
		return nil
	}
	return a.Memory
}

func (a *ApprovalTool) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *ApprovalTool) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ApprovalTool) SetName(name string) {
	a.Name = name
	a.require(approvalToolFieldName)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ApprovalTool) SetSource(source *ApprovalToolSource) {
	a.Source = source
	a.require(approvalToolFieldSource)
}

// SetMemory sets the Memory field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ApprovalTool) SetMemory(memory *ApprovalToolMemory) {
	a.Memory = memory
	a.require(approvalToolFieldMemory)
}

func (a *ApprovalTool) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalTool
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApprovalTool(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalTool) MarshalJSON() ([]byte, error) {
	type embed ApprovalTool
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *ApprovalTool) String() string {
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
	approvalToolMemoryFieldOperation = big.NewInt(1 << 0)
)

// What only a memory tool has, set only when `source.kind` is `memory`.
type ApprovalToolMemory struct {
	// What the memory tool does.
	Operation ApprovalToolMemoryOperation `json:"operation" url:"operation"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalToolMemory) GetOperation() ApprovalToolMemoryOperation {
	if a == nil {
		return ""
	}
	return a.Operation
}

func (a *ApprovalToolMemory) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *ApprovalToolMemory) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetOperation sets the Operation field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ApprovalToolMemory) SetOperation(operation ApprovalToolMemoryOperation) {
	a.Operation = operation
	a.require(approvalToolMemoryFieldOperation)
}

func (a *ApprovalToolMemory) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalToolMemory
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApprovalToolMemory(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalToolMemory) MarshalJSON() ([]byte, error) {
	type embed ApprovalToolMemory
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *ApprovalToolMemory) String() string {
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

// What the memory tool does.
type ApprovalToolMemoryOperation string

const (
	ApprovalToolMemoryOperationQuery  ApprovalToolMemoryOperation = "query"
	ApprovalToolMemoryOperationWrite  ApprovalToolMemoryOperation = "write"
	ApprovalToolMemoryOperationExpand ApprovalToolMemoryOperation = "expand"
)

func NewApprovalToolMemoryOperationFromString(s string) (ApprovalToolMemoryOperation, error) {
	switch s {
	case "query":
		return ApprovalToolMemoryOperationQuery, nil
	case "write":
		return ApprovalToolMemoryOperationWrite, nil
	case "expand":
		return ApprovalToolMemoryOperationExpand, nil
	}
	var t ApprovalToolMemoryOperation
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalToolMemoryOperation) Ptr() *ApprovalToolMemoryOperation {
	return &a
}

var (
	approvalToolSourceFieldKind = big.NewInt(1 << 0)
	approvalToolSourceFieldID   = big.NewInt(1 << 1)
	approvalToolSourceFieldName = big.NewInt(1 << 2)
)

// Where a tool comes from. `id` and `name` are those of the MCP connection, installed particle or memory interaction profile the tool comes from, and are left out when there is none.
type ApprovalToolSource struct {
	Kind ApprovalToolSourceKind `json:"kind" url:"kind"`
	ID   *string                `json:"id,omitempty" url:"id,omitempty"`
	Name *string                `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalToolSource) GetKind() ApprovalToolSourceKind {
	if a == nil {
		return ""
	}
	return a.Kind
}

func (a *ApprovalToolSource) GetID() *string {
	if a == nil {
		return nil
	}
	return a.ID
}

func (a *ApprovalToolSource) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *ApprovalToolSource) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *ApprovalToolSource) require(field *big.Int) {
	next := new(big.Int)
	if a.explicitFields != nil {
		next.Set(a.explicitFields)
	}
	next.Or(next, field)
	a.explicitFields = next
}

// SetKind sets the Kind field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ApprovalToolSource) SetKind(kind ApprovalToolSourceKind) {
	a.Kind = kind
	a.require(approvalToolSourceFieldKind)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ApprovalToolSource) SetID(id *string) {
	a.ID = id
	a.require(approvalToolSourceFieldID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ApprovalToolSource) SetName(name *string) {
	a.Name = name
	a.require(approvalToolSourceFieldName)
}

func (a *ApprovalToolSource) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalToolSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApprovalToolSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalToolSource) MarshalJSON() ([]byte, error) {
	type embed ApprovalToolSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *ApprovalToolSource) String() string {
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

type ApprovalToolSourceKind string

const (
	ApprovalToolSourceKindMcp            ApprovalToolSourceKind = "mcp"
	ApprovalToolSourceKindParticle       ApprovalToolSourceKind = "particle"
	ApprovalToolSourceKindClientTask     ApprovalToolSourceKind = "client_task"
	ApprovalToolSourceKindResourceReader ApprovalToolSourceKind = "resource_reader"
	ApprovalToolSourceKindSkill          ApprovalToolSourceKind = "skill"
	ApprovalToolSourceKindMemory         ApprovalToolSourceKind = "memory"
)

func NewApprovalToolSourceKindFromString(s string) (ApprovalToolSourceKind, error) {
	switch s {
	case "mcp":
		return ApprovalToolSourceKindMcp, nil
	case "particle":
		return ApprovalToolSourceKindParticle, nil
	case "client_task":
		return ApprovalToolSourceKindClientTask, nil
	case "resource_reader":
		return ApprovalToolSourceKindResourceReader, nil
	case "skill":
		return ApprovalToolSourceKindSkill, nil
	case "memory":
		return ApprovalToolSourceKindMemory, nil
	}
	var t ApprovalToolSourceKind
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalToolSourceKind) Ptr() *ApprovalToolSourceKind {
	return &a
}

var (
	boundaryProblemFieldType       = big.NewInt(1 << 0)
	boundaryProblemFieldTitle      = big.NewInt(1 << 1)
	boundaryProblemFieldStatus     = big.NewInt(1 << 2)
	boundaryProblemFieldDetail     = big.NewInt(1 << 3)
	boundaryProblemFieldParam      = big.NewInt(1 << 4)
	boundaryProblemFieldDecisionID = big.NewInt(1 << 5)
	boundaryProblemFieldBoundary   = big.NewInt(1 << 6)
	boundaryProblemFieldRule       = big.NewInt(1 << 7)
)

// A problem document for an action a boundary refused.
type BoundaryProblem struct {
	Type   string `json:"type" url:"type"`
	Title  string `json:"title" url:"title"`
	Status int    `json:"status" url:"status"`
	Detail string `json:"detail" url:"detail"`
	// The input the problem is with, when it is one.
	Param *string `json:"param,omitempty" url:"param,omitempty"`
	// The ID an admin looks the decision up by.
	DecisionID string `json:"decision_id" url:"decision_id"`
	// The boundary that refused. Left out for a caller through an access point.
	Boundary *string `json:"boundary,omitempty" url:"boundary,omitempty"`
	// The rule that refused. Left out for a caller through an access point.
	Rule *string `json:"rule,omitempty" url:"rule,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BoundaryProblem) GetType() string {
	if b == nil {
		return ""
	}
	return b.Type
}

func (b *BoundaryProblem) GetTitle() string {
	if b == nil {
		return ""
	}
	return b.Title
}

func (b *BoundaryProblem) GetStatus() int {
	if b == nil {
		return 0
	}
	return b.Status
}

func (b *BoundaryProblem) GetDetail() string {
	if b == nil {
		return ""
	}
	return b.Detail
}

func (b *BoundaryProblem) GetParam() *string {
	if b == nil {
		return nil
	}
	return b.Param
}

func (b *BoundaryProblem) GetDecisionID() string {
	if b == nil {
		return ""
	}
	return b.DecisionID
}

func (b *BoundaryProblem) GetBoundary() *string {
	if b == nil {
		return nil
	}
	return b.Boundary
}

func (b *BoundaryProblem) GetRule() *string {
	if b == nil {
		return nil
	}
	return b.Rule
}

func (b *BoundaryProblem) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b *BoundaryProblem) require(field *big.Int) {
	next := new(big.Int)
	if b.explicitFields != nil {
		next.Set(b.explicitFields)
	}
	next.Or(next, field)
	b.explicitFields = next
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetType(type_ string) {
	b.Type = type_
	b.require(boundaryProblemFieldType)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetTitle(title string) {
	b.Title = title
	b.require(boundaryProblemFieldTitle)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetStatus(status int) {
	b.Status = status
	b.require(boundaryProblemFieldStatus)
}

// SetDetail sets the Detail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetDetail(detail string) {
	b.Detail = detail
	b.require(boundaryProblemFieldDetail)
}

// SetParam sets the Param field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetParam(param *string) {
	b.Param = param
	b.require(boundaryProblemFieldParam)
}

// SetDecisionID sets the DecisionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetDecisionID(decisionID string) {
	b.DecisionID = decisionID
	b.require(boundaryProblemFieldDecisionID)
}

// SetBoundary sets the Boundary field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetBoundary(boundary *string) {
	b.Boundary = boundary
	b.require(boundaryProblemFieldBoundary)
}

// SetRule sets the Rule field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BoundaryProblem) SetRule(rule *string) {
	b.Rule = rule
	b.require(boundaryProblemFieldRule)
}

func (b *BoundaryProblem) UnmarshalJSON(data []byte) error {
	type unmarshaler BoundaryProblem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BoundaryProblem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BoundaryProblem) MarshalJSON() ([]byte, error) {
	type embed BoundaryProblem
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, b.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (b *BoundaryProblem) String() string {
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
	cancelledEventFieldID   = big.NewInt(1 << 0)
	cancelledEventFieldData = big.NewInt(1 << 1)
)

// The request was cancelled, which ends it.
type CancelledEvent struct {
	ID   string `json:"id" url:"id"`
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CancelledEvent) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *CancelledEvent) GetData() string {
	if c == nil {
		return ""
	}
	return c.Data
}

func (c *CancelledEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CancelledEvent) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelledEvent) SetID(id string) {
	c.ID = id
	c.require(cancelledEventFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CancelledEvent) SetData(data string) {
	c.Data = data
	c.require(cancelledEventFieldData)
}

func (c *CancelledEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler CancelledEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CancelledEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CancelledEvent) MarshalJSON() ([]byte, error) {
	type embed CancelledEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CancelledEvent) String() string {
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
	clientTaskFieldName          = big.NewInt(1 << 0)
	clientTaskFieldInput         = big.NewInt(1 << 1)
	clientTaskFieldPresentations = big.NewInt(1 << 2)
)

type ClientTask struct {
	// The client task definition the task came from, such as `ask_question`.
	Name string `json:"name" url:"name"`
	// What the user needs to do the task. It matches the client task definition's input schema.
	Input map[string]any `json:"input" url:"input"`
	// Ways a client can present the task, such as a UI form.
	Presentations []*Presentation `json:"presentations,omitempty" url:"presentations,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ClientTask) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *ClientTask) GetInput() map[string]any {
	if c == nil {
		return nil
	}
	return c.Input
}

func (c *ClientTask) GetPresentations() []*Presentation {
	if c == nil {
		return nil
	}
	return c.Presentations
}

func (c *ClientTask) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ClientTask) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ClientTask) SetName(name string) {
	c.Name = name
	c.require(clientTaskFieldName)
}

// SetInput sets the Input field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ClientTask) SetInput(input map[string]any) {
	c.Input = input
	c.require(clientTaskFieldInput)
}

// SetPresentations sets the Presentations field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ClientTask) SetPresentations(presentations []*Presentation) {
	c.Presentations = presentations
	c.require(clientTaskFieldPresentations)
}

func (c *ClientTask) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientTask
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientTask(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientTask) MarshalJSON() ([]byte, error) {
	type embed ClientTask
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ClientTask) String() string {
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
	clientTaskDataFieldResponseSchema = big.NewInt(1 << 0)
	clientTaskDataFieldClientTask     = big.NewInt(1 << 1)
)

// A task an agent sent with one of its client task definitions, such as a question for the user.
type ClientTaskData struct {
	// The JSON Schema of the response that completes the task. It is the client task definition's output schema.
	ResponseSchema map[string]any `json:"response_schema" url:"response_schema"`
	ClientTask     *ClientTask    `json:"client_task" url:"client_task"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ClientTaskData) GetResponseSchema() map[string]any {
	if c == nil {
		return nil
	}
	return c.ResponseSchema
}

func (c *ClientTaskData) GetClientTask() *ClientTask {
	if c == nil {
		return nil
	}
	return c.ClientTask
}

func (c *ClientTaskData) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ClientTaskData) require(field *big.Int) {
	next := new(big.Int)
	if c.explicitFields != nil {
		next.Set(c.explicitFields)
	}
	next.Or(next, field)
	c.explicitFields = next
}

// SetResponseSchema sets the ResponseSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ClientTaskData) SetResponseSchema(responseSchema map[string]any) {
	c.ResponseSchema = responseSchema
	c.require(clientTaskDataFieldResponseSchema)
}

// SetClientTask sets the ClientTask field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ClientTaskData) SetClientTask(clientTask *ClientTask) {
	c.ClientTask = clientTask
	c.require(clientTaskDataFieldClientTask)
}

func (c *ClientTaskData) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientTaskData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientTaskData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientTaskData) MarshalJSON() ([]byte, error) {
	type embed ClientTaskData
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ClientTaskData) String() string {
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

// Why the request failed, which ends it.
type ErrorEvent struct {
	ID   string `json:"id" url:"id"`
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *ErrorEvent) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *ErrorEvent) GetData() string {
	if e == nil {
		return ""
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
func (e *ErrorEvent) SetID(id string) {
	e.ID = id
	e.require(errorEventFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ErrorEvent) SetData(data string) {
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
	messageEventFieldID   = big.NewInt(1 << 0)
	messageEventFieldData = big.NewInt(1 << 1)
)

// A progress message for the user.
type MessageEvent struct {
	ID   string `json:"id" url:"id"`
	Data string `json:"data" url:"data"`

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

func (m *MessageEvent) GetData() string {
	if m == nil {
		return ""
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
func (m *MessageEvent) SetData(data string) {
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
	ID   string `json:"id" url:"id"`
	Data string `json:"data" url:"data"`

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

func (n *NotificationChannelMessageEvent) GetData() string {
	if n == nil {
		return ""
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
func (n *NotificationChannelMessageEvent) SetData(data string) {
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

var (
	notificationChannelServerErrorEventFieldData = big.NewInt(1 << 0)
)

// A server-sent event saying the server could not read the stream further.
type NotificationChannelServerErrorEvent struct {
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NotificationChannelServerErrorEvent) GetData() string {
	if n == nil {
		return ""
	}
	return n.Data
}

func (n *NotificationChannelServerErrorEvent) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n *NotificationChannelServerErrorEvent) require(field *big.Int) {
	next := new(big.Int)
	if n.explicitFields != nil {
		next.Set(n.explicitFields)
	}
	next.Or(next, field)
	n.explicitFields = next
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (n *NotificationChannelServerErrorEvent) SetData(data string) {
	n.Data = data
	n.require(notificationChannelServerErrorEventFieldData)
}

func (n *NotificationChannelServerErrorEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler NotificationChannelServerErrorEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NotificationChannelServerErrorEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NotificationChannelServerErrorEvent) MarshalJSON() ([]byte, error) {
	type embed NotificationChannelServerErrorEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, n.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (n *NotificationChannelServerErrorEvent) String() string {
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

// One server-sent event of a notification channel stream. `data` is the event as JSON.
type NotificationChannelSseEvent struct {
	Event       string
	ServerError *NotificationChannelServerErrorEvent
	Message     *NotificationChannelMessageEvent

	rawJSON json.RawMessage
}

func (n *NotificationChannelSseEvent) GetEvent() string {
	if n == nil {
		return ""
	}
	return n.Event
}

func (n *NotificationChannelSseEvent) GetServerError() *NotificationChannelServerErrorEvent {
	if n == nil {
		return nil
	}
	return n.ServerError
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
	case "server_error":
		value := new(NotificationChannelServerErrorEvent)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		n.ServerError = value
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
	if n.ServerError != nil {
		return internal.MarshalJSONWithExtraProperty(n.ServerError, "event", "server_error")
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
	VisitServerError(*NotificationChannelServerErrorEvent) error
	VisitMessage(*NotificationChannelMessageEvent) error
}

func (n *NotificationChannelSseEvent) Accept(visitor NotificationChannelSseEventVisitor) error {
	if n.ServerError != nil {
		return visitor.VisitServerError(n.ServerError)
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
	if n.ServerError != nil {
		fields = append(fields, "server_error")
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

// One way a client can present a task to the user. `type` says what `content` holds.
type Presentation struct {
	Type string
	A2UI *A2UIPresentation

	rawJSON json.RawMessage
}

func (p *Presentation) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *Presentation) GetA2UI() *A2UIPresentation {
	if p == nil {
		return nil
	}
	return p.A2UI
}

func (p *Presentation) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", p)
	}
	switch unmarshaler.Type {
	case "a2ui":
		value := new(A2UIPresentation)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.A2UI = value
	}
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p Presentation) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.A2UI != nil {
		return internal.MarshalJSONWithExtraProperty(p.A2UI, "type", "a2ui")
	}
	if len(p.rawJSON) > 0 {
		return p.rawJSON, nil
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type PresentationVisitor interface {
	VisitA2UI(*A2UIPresentation) error
}

func (p *Presentation) Accept(visitor PresentationVisitor) error {
	if p.A2UI != nil {
		return visitor.VisitA2UI(p.A2UI)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *Presentation) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.A2UI != nil {
		fields = append(fields, "a2ui")
	}
	if len(fields) == 0 {
		if p.Type != "" {
			if len(p.rawJSON) > 0 {
				return nil
			}
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.Type)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.Type != "" {
		field := fields[0]
		if p.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.Type,
				p,
			)
		}
	}
	return nil
}

var (
	problemFieldType   = big.NewInt(1 << 0)
	problemFieldTitle  = big.NewInt(1 << 1)
	problemFieldStatus = big.NewInt(1 << 2)
	problemFieldDetail = big.NewInt(1 << 3)
	problemFieldParam  = big.NewInt(1 << 4)
)

// An RFC 9457 problem document.
type Problem struct {
	Type   string `json:"type" url:"type"`
	Title  string `json:"title" url:"title"`
	Status int    `json:"status" url:"status"`
	Detail string `json:"detail" url:"detail"`
	// The input the problem is with, when it is one.
	Param *string `json:"param,omitempty" url:"param,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *Problem) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *Problem) GetTitle() string {
	if p == nil {
		return ""
	}
	return p.Title
}

func (p *Problem) GetStatus() int {
	if p == nil {
		return 0
	}
	return p.Status
}

func (p *Problem) GetDetail() string {
	if p == nil {
		return ""
	}
	return p.Detail
}

func (p *Problem) GetParam() *string {
	if p == nil {
		return nil
	}
	return p.Param
}

func (p *Problem) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *Problem) require(field *big.Int) {
	next := new(big.Int)
	if p.explicitFields != nil {
		next.Set(p.explicitFields)
	}
	next.Or(next, field)
	p.explicitFields = next
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Problem) SetType(type_ string) {
	p.Type = type_
	p.require(problemFieldType)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Problem) SetTitle(title string) {
	p.Title = title
	p.require(problemFieldTitle)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Problem) SetStatus(status int) {
	p.Status = status
	p.require(problemFieldStatus)
}

// SetDetail sets the Detail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Problem) SetDetail(detail string) {
	p.Detail = detail
	p.require(problemFieldDetail)
}

// SetParam sets the Param field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *Problem) SetParam(param *string) {
	p.Param = param
	p.require(problemFieldParam)
}

func (p *Problem) UnmarshalJSON(data []byte) error {
	type unmarshaler Problem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = Problem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *Problem) MarshalJSON() ([]byte, error) {
	type embed Problem
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *Problem) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

var (
	requestServerSentEventFieldID    = big.NewInt(1 << 0)
	requestServerSentEventFieldEvent = big.NewInt(1 << 1)
	requestServerSentEventFieldData  = big.NewInt(1 << 2)
)

// One server-sent event of a request's stream. `event` is the event's `type`, except that an `error` is named `request_error`, because a browser's `EventSource` also fires `error` when the connection fails. `data` is the event as JSON. The `started` and `server_error` events have no `id`.
type RequestServerSentEvent struct {
	// The event's ID. Send it as `Last-Event-ID` to resume after this event.
	ID    *string                     `json:"id,omitempty" url:"id,omitempty"`
	Event RequestServerSentEventEvent `json:"event" url:"event"`
	Data  string                      `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RequestServerSentEvent) GetID() *string {
	if r == nil {
		return nil
	}
	return r.ID
}

func (r *RequestServerSentEvent) GetEvent() RequestServerSentEventEvent {
	if r == nil {
		return ""
	}
	return r.Event
}

func (r *RequestServerSentEvent) GetData() string {
	if r == nil {
		return ""
	}
	return r.Data
}

func (r *RequestServerSentEvent) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *RequestServerSentEvent) require(field *big.Int) {
	next := new(big.Int)
	if r.explicitFields != nil {
		next.Set(r.explicitFields)
	}
	next.Or(next, field)
	r.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RequestServerSentEvent) SetID(id *string) {
	r.ID = id
	r.require(requestServerSentEventFieldID)
}

// SetEvent sets the Event field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RequestServerSentEvent) SetEvent(event RequestServerSentEventEvent) {
	r.Event = event
	r.require(requestServerSentEventFieldEvent)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RequestServerSentEvent) SetData(data string) {
	r.Data = data
	r.require(requestServerSentEventFieldData)
}

func (r *RequestServerSentEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler RequestServerSentEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RequestServerSentEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RequestServerSentEvent) MarshalJSON() ([]byte, error) {
	type embed RequestServerSentEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *RequestServerSentEvent) String() string {
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

type RequestServerSentEventEvent string

const (
	RequestServerSentEventEventStarted      RequestServerSentEventEvent = "started"
	RequestServerSentEventEventTask         RequestServerSentEventEvent = "task"
	RequestServerSentEventEventTaskResponse RequestServerSentEventEvent = "task_response"
	RequestServerSentEventEventMessage      RequestServerSentEventEvent = "message"
	RequestServerSentEventEventThinking     RequestServerSentEventEvent = "thinking"
	RequestServerSentEventEventResponse     RequestServerSentEventEvent = "response"
	RequestServerSentEventEventRequestError RequestServerSentEventEvent = "request_error"
	RequestServerSentEventEventCancelled    RequestServerSentEventEvent = "cancelled"
	RequestServerSentEventEventServerError  RequestServerSentEventEvent = "server_error"
)

func NewRequestServerSentEventEventFromString(s string) (RequestServerSentEventEvent, error) {
	switch s {
	case "started":
		return RequestServerSentEventEventStarted, nil
	case "task":
		return RequestServerSentEventEventTask, nil
	case "task_response":
		return RequestServerSentEventEventTaskResponse, nil
	case "message":
		return RequestServerSentEventEventMessage, nil
	case "thinking":
		return RequestServerSentEventEventThinking, nil
	case "response":
		return RequestServerSentEventEventResponse, nil
	case "request_error":
		return RequestServerSentEventEventRequestError, nil
	case "cancelled":
		return RequestServerSentEventEventCancelled, nil
	case "server_error":
		return RequestServerSentEventEventServerError, nil
	}
	var t RequestServerSentEventEvent
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestServerSentEventEvent) Ptr() *RequestServerSentEventEvent {
	return &r
}

var (
	responseEventFieldID   = big.NewInt(1 << 0)
	responseEventFieldData = big.NewInt(1 << 1)
)

// The request's response, which ends it.
type ResponseEvent struct {
	ID   string         `json:"id" url:"id"`
	Data map[string]any `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ResponseEvent) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *ResponseEvent) GetData() map[string]any {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *ResponseEvent) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ResponseEvent) require(field *big.Int) {
	next := new(big.Int)
	if r.explicitFields != nil {
		next.Set(r.explicitFields)
	}
	next.Or(next, field)
	r.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResponseEvent) SetID(id string) {
	r.ID = id
	r.require(responseEventFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResponseEvent) SetData(data map[string]any) {
	r.Data = data
	r.require(responseEventFieldData)
}

func (r *ResponseEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler ResponseEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResponseEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResponseEvent) MarshalJSON() ([]byte, error) {
	type embed ResponseEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ResponseEvent) String() string {
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
	serverErrorEventFieldError = big.NewInt(1 << 0)
)

// The server could not read the stream further. Resume it.
type ServerErrorEvent struct {
	Error string `json:"error" url:"error"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *ServerErrorEvent) GetError() string {
	if s == nil {
		return ""
	}
	return s.Error
}

func (s *ServerErrorEvent) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *ServerErrorEvent) require(field *big.Int) {
	next := new(big.Int)
	if s.explicitFields != nil {
		next.Set(s.explicitFields)
	}
	next.Or(next, field)
	s.explicitFields = next
}

// SetError sets the Error field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *ServerErrorEvent) SetError(error_ string) {
	s.Error = error_
	s.require(serverErrorEventFieldError)
}

func (s *ServerErrorEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler ServerErrorEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ServerErrorEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *ServerErrorEvent) MarshalJSON() ([]byte, error) {
	type embed ServerErrorEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *ServerErrorEvent) String() string {
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

// What a task holds. `kind` says what the task is, and the task's details are under the property named after it. Every kind has a `response_schema`, the JSON Schema of the response that completes the task.
type TaskData struct {
	Kind             string
	ClientTask       *ClientTaskData
	ToolCallApproval *ToolCallApprovalData

	rawJSON json.RawMessage
}

func (t *TaskData) GetKind() string {
	if t == nil {
		return ""
	}
	return t.Kind
}

func (t *TaskData) GetClientTask() *ClientTaskData {
	if t == nil {
		return nil
	}
	return t.ClientTask
}

func (t *TaskData) GetToolCallApproval() *ToolCallApprovalData {
	if t == nil {
		return nil
	}
	return t.ToolCallApproval
}

func (t *TaskData) UnmarshalJSON(data []byte) error {
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
		value := new(ClientTaskData)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ClientTask = value
	case "tool_call_approval":
		value := new(ToolCallApprovalData)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ToolCallApproval = value
	}
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t TaskData) MarshalJSON() ([]byte, error) {
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

type TaskDataVisitor interface {
	VisitClientTask(*ClientTaskData) error
	VisitToolCallApproval(*ToolCallApprovalData) error
}

func (t *TaskData) Accept(visitor TaskDataVisitor) error {
	if t.ClientTask != nil {
		return visitor.VisitClientTask(t.ClientTask)
	}
	if t.ToolCallApproval != nil {
		return visitor.VisitToolCallApproval(t.ToolCallApproval)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TaskData) validate() error {
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
	taskEventFieldID   = big.NewInt(1 << 0)
	taskEventFieldData = big.NewInt(1 << 1)
)

// A task opened. Its data is the task without its ID; the event's ID is the task's.
type TaskEvent struct {
	ID   string    `json:"id" url:"id"`
	Data *TaskData `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskEvent) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

func (t *TaskEvent) GetData() *TaskData {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TaskEvent) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *TaskEvent) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEvent) SetID(id string) {
	t.ID = id
	t.require(taskEventFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskEvent) SetData(data *TaskData) {
	t.Data = data
	t.require(taskEventFieldData)
}

func (t *TaskEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler TaskEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TaskEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TaskEvent) MarshalJSON() ([]byte, error) {
	type embed TaskEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *TaskEvent) String() string {
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
	taskResponseEventFieldID        = big.NewInt(1 << 0)
	taskResponseEventFieldData      = big.NewInt(1 << 1)
	taskResponseEventFieldInReplyTo = big.NewInt(1 << 2)
)

// A task was answered.
type TaskResponseEvent struct {
	ID string `json:"id" url:"id"`
	// The task's response, or the string `cancelled` when the task was cancelled.
	Data *TaskResponseEventData `json:"data" url:"data"`
	// The ID of the task answered.
	InReplyTo string `json:"in_reply_to" url:"in_reply_to"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TaskResponseEvent) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

func (t *TaskResponseEvent) GetData() *TaskResponseEventData {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TaskResponseEvent) GetInReplyTo() string {
	if t == nil {
		return ""
	}
	return t.InReplyTo
}

func (t *TaskResponseEvent) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *TaskResponseEvent) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEvent) SetID(id string) {
	t.ID = id
	t.require(taskResponseEventFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEvent) SetData(data *TaskResponseEventData) {
	t.Data = data
	t.require(taskResponseEventFieldData)
}

// SetInReplyTo sets the InReplyTo field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TaskResponseEvent) SetInReplyTo(inReplyTo string) {
	t.InReplyTo = inReplyTo
	t.require(taskResponseEventFieldInReplyTo)
}

func (t *TaskResponseEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler TaskResponseEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TaskResponseEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TaskResponseEvent) MarshalJSON() ([]byte, error) {
	type embed TaskResponseEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *TaskResponseEvent) String() string {
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

// The task's response, or the string `cancelled` when the task was cancelled.
type TaskResponseEventData struct {
	StringUnknownMap         map[string]any
	TaskResponseEventDataOne TaskResponseEventDataOne

	typ string
}

func (t *TaskResponseEventData) GetStringUnknownMap() map[string]any {
	if t == nil {
		return nil
	}
	return t.StringUnknownMap
}

func (t *TaskResponseEventData) GetTaskResponseEventDataOne() TaskResponseEventDataOne {
	if t == nil {
		return ""
	}
	return t.TaskResponseEventDataOne
}

func (t *TaskResponseEventData) UnmarshalJSON(data []byte) error {
	var valueStringUnknownMap map[string]any
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		t.typ = "StringUnknownMap"
		t.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	var valueTaskResponseEventDataOne TaskResponseEventDataOne
	if err := json.Unmarshal(data, &valueTaskResponseEventDataOne); err == nil {
		t.typ = "TaskResponseEventDataOne"
		t.TaskResponseEventDataOne = valueTaskResponseEventDataOne
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TaskResponseEventData) MarshalJSON() ([]byte, error) {
	if t.typ == "StringUnknownMap" || t.StringUnknownMap != nil {
		return json.Marshal(t.StringUnknownMap)
	}
	if t.typ == "TaskResponseEventDataOne" || t.TaskResponseEventDataOne != "" {
		return json.Marshal(t.TaskResponseEventDataOne)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TaskResponseEventDataVisitor interface {
	VisitStringUnknownMap(map[string]any) error
	VisitTaskResponseEventDataOne(TaskResponseEventDataOne) error
}

func (t *TaskResponseEventData) Accept(visitor TaskResponseEventDataVisitor) error {
	if t.typ == "StringUnknownMap" || t.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(t.StringUnknownMap)
	}
	if t.typ == "TaskResponseEventDataOne" || t.TaskResponseEventDataOne != "" {
		return visitor.VisitTaskResponseEventDataOne(t.TaskResponseEventDataOne)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TaskResponseEventDataOne string

const (
	TaskResponseEventDataOneCancelled TaskResponseEventDataOne = "cancelled"
)

func NewTaskResponseEventDataOneFromString(s string) (TaskResponseEventDataOne, error) {
	switch s {
	case "cancelled":
		return TaskResponseEventDataOneCancelled, nil
	}
	var t TaskResponseEventDataOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TaskResponseEventDataOne) Ptr() *TaskResponseEventDataOne {
	return &t
}

var (
	thinkingEventFieldID   = big.NewInt(1 << 0)
	thinkingEventFieldData = big.NewInt(1 << 1)
)

// What the model is thinking.
type ThinkingEvent struct {
	ID   string `json:"id" url:"id"`
	Data string `json:"data" url:"data"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *ThinkingEvent) GetID() string {
	if t == nil {
		return ""
	}
	return t.ID
}

func (t *ThinkingEvent) GetData() string {
	if t == nil {
		return ""
	}
	return t.Data
}

func (t *ThinkingEvent) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *ThinkingEvent) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ThinkingEvent) SetID(id string) {
	t.ID = id
	t.require(thinkingEventFieldID)
}

// SetData sets the Data field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ThinkingEvent) SetData(data string) {
	t.Data = data
	t.require(thinkingEventFieldData)
}

func (t *ThinkingEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler ThinkingEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ThinkingEvent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *ThinkingEvent) MarshalJSON() ([]byte, error) {
	type embed ThinkingEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *ThinkingEvent) String() string {
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
	toolCallApprovalFieldTool       = big.NewInt(1 << 0)
	toolCallApprovalFieldArguments  = big.NewInt(1 << 1)
	toolCallApprovalFieldMessage    = big.NewInt(1 << 2)
	toolCallApprovalFieldDecisionID = big.NewInt(1 << 3)
)

type ToolCallApproval struct {
	Tool *ApprovalTool `json:"tool" url:"tool"`
	// The arguments the call will be made with.
	Arguments map[string]any `json:"arguments" url:"arguments"`
	// Why the call needs approval, for the user to read, when there is a reason.
	Message *string `json:"message,omitempty" url:"message,omitempty"`
	// An ID an admin can look up to see why approval was asked for.
	DecisionID string `json:"decision_id" url:"decision_id"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *ToolCallApproval) GetTool() *ApprovalTool {
	if t == nil {
		return nil
	}
	return t.Tool
}

func (t *ToolCallApproval) GetArguments() map[string]any {
	if t == nil {
		return nil
	}
	return t.Arguments
}

func (t *ToolCallApproval) GetMessage() *string {
	if t == nil {
		return nil
	}
	return t.Message
}

func (t *ToolCallApproval) GetDecisionID() string {
	if t == nil {
		return ""
	}
	return t.DecisionID
}

func (t *ToolCallApproval) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *ToolCallApproval) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetTool sets the Tool field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApproval) SetTool(tool *ApprovalTool) {
	t.Tool = tool
	t.require(toolCallApprovalFieldTool)
}

// SetArguments sets the Arguments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApproval) SetArguments(arguments map[string]any) {
	t.Arguments = arguments
	t.require(toolCallApprovalFieldArguments)
}

// SetMessage sets the Message field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApproval) SetMessage(message *string) {
	t.Message = message
	t.require(toolCallApprovalFieldMessage)
}

// SetDecisionID sets the DecisionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApproval) SetDecisionID(decisionID string) {
	t.DecisionID = decisionID
	t.require(toolCallApprovalFieldDecisionID)
}

func (t *ToolCallApproval) UnmarshalJSON(data []byte) error {
	type unmarshaler ToolCallApproval
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ToolCallApproval(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *ToolCallApproval) MarshalJSON() ([]byte, error) {
	type embed ToolCallApproval
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *ToolCallApproval) String() string {
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
	toolCallApprovalDataFieldResponseSchema   = big.NewInt(1 << 0)
	toolCallApprovalDataFieldToolCallApproval = big.NewInt(1 << 1)
)

// A tool call the agent wants to make, held until the user approves it. The `response_schema` is always that of `ToolCallApprovalResponse`.
type ToolCallApprovalData struct {
	// The JSON Schema of `ToolCallApprovalResponse`.
	ResponseSchema   map[string]any    `json:"response_schema" url:"response_schema"`
	ToolCallApproval *ToolCallApproval `json:"tool_call_approval" url:"tool_call_approval"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *ToolCallApprovalData) GetResponseSchema() map[string]any {
	if t == nil {
		return nil
	}
	return t.ResponseSchema
}

func (t *ToolCallApprovalData) GetToolCallApproval() *ToolCallApproval {
	if t == nil {
		return nil
	}
	return t.ToolCallApproval
}

func (t *ToolCallApprovalData) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *ToolCallApprovalData) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetResponseSchema sets the ResponseSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApprovalData) SetResponseSchema(responseSchema map[string]any) {
	t.ResponseSchema = responseSchema
	t.require(toolCallApprovalDataFieldResponseSchema)
}

// SetToolCallApproval sets the ToolCallApproval field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApprovalData) SetToolCallApproval(toolCallApproval *ToolCallApproval) {
	t.ToolCallApproval = toolCallApproval
	t.require(toolCallApprovalDataFieldToolCallApproval)
}

func (t *ToolCallApprovalData) UnmarshalJSON(data []byte) error {
	type unmarshaler ToolCallApprovalData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ToolCallApprovalData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *ToolCallApprovalData) MarshalJSON() ([]byte, error) {
	type embed ToolCallApprovalData
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *ToolCallApprovalData) String() string {
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
	toolCallApprovalResponseFieldApproved = big.NewInt(1 << 0)
	toolCallApprovalResponseFieldComment  = big.NewInt(1 << 1)
)

// The response that completes a `tool_call_approval`.
type ToolCallApprovalResponse struct {
	// Whether the tool call may be made.
	Approved bool `json:"approved" url:"approved"`
	// Why, for the agent to read when the call is declined.
	Comment *string `json:"comment,omitempty" url:"comment,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *ToolCallApprovalResponse) GetApproved() bool {
	if t == nil {
		return false
	}
	return t.Approved
}

func (t *ToolCallApprovalResponse) GetComment() *string {
	if t == nil {
		return nil
	}
	return t.Comment
}

func (t *ToolCallApprovalResponse) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t *ToolCallApprovalResponse) require(field *big.Int) {
	next := new(big.Int)
	if t.explicitFields != nil {
		next.Set(t.explicitFields)
	}
	next.Or(next, field)
	t.explicitFields = next
}

// SetApproved sets the Approved field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApprovalResponse) SetApproved(approved bool) {
	t.Approved = approved
	t.require(toolCallApprovalResponseFieldApproved)
}

// SetComment sets the Comment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *ToolCallApprovalResponse) SetComment(comment *string) {
	t.Comment = comment
	t.require(toolCallApprovalResponseFieldComment)
}

func (t *ToolCallApprovalResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ToolCallApprovalResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ToolCallApprovalResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *ToolCallApprovalResponse) MarshalJSON() ([]byte, error) {
	type embed ToolCallApprovalResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, t.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (t *ToolCallApprovalResponse) String() string {
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
