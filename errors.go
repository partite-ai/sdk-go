// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	core "github.com/partite-ai/sdk-go/core"
)

// The request was invalid or cannot be otherwise served.
type BadRequestError struct {
	*core.APIError
	Body *BadRequestErrorResponse
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *BadRequestErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

func (b *BadRequestError) GetBody() *BadRequestErrorResponse {
	if b == nil {
		return nil
	}
	return b.Body
}

// A resource with the same unique identifier already exists.
type ConflictError struct {
	*core.APIError
	Body *GenericErrorResponse
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	var body *GenericErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 409
	c.Body = body
	return nil
}

func (c *ConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

func (c *ConflictError) Unwrap() error {
	return c.APIError
}

func (c *ConflictError) GetBody() *GenericErrorResponse {
	if c == nil {
		return nil
	}
	return c.Body
}

// Nothing to return yet: the wait ended, or a proxy in front gave up on the connection. Poll again; the call is a GET and safe to repeat. Mesh sends no body and `Retry-After: 0`; a proxy's 504 may carry a body of its own, so ignore the body.
type GatewayTimeoutError struct {
	*core.APIError
	Body any
}

func (g *GatewayTimeoutError) UnmarshalJSON(data []byte) error {
	var body any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	g.StatusCode = 504
	g.Body = body
	return nil
}

func (g *GatewayTimeoutError) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.Body)
}

func (g *GatewayTimeoutError) Unwrap() error {
	return g.APIError
}

func (g *GatewayTimeoutError) GetBody() any {
	if g == nil {
		return nil
	}
	return g.Body
}

// An unexpected error occurred on the server.
type InternalServerError struct {
	*core.APIError
	Body *GenericErrorResponse
}

func (i *InternalServerError) UnmarshalJSON(data []byte) error {
	var body *GenericErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	i.StatusCode = 500
	i.Body = body
	return nil
}

func (i *InternalServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Body)
}

func (i *InternalServerError) Unwrap() error {
	return i.APIError
}

func (i *InternalServerError) GetBody() *GenericErrorResponse {
	if i == nil {
		return nil
	}
	return i.Body
}

// The Accept header allows none of the formats.
type NotAcceptableError struct {
	*core.APIError
	Body *GenericErrorResponse
}

func (n *NotAcceptableError) UnmarshalJSON(data []byte) error {
	var body *GenericErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 406
	n.Body = body
	return nil
}

func (n *NotAcceptableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotAcceptableError) Unwrap() error {
	return n.APIError
}

func (n *NotAcceptableError) GetBody() *GenericErrorResponse {
	if n == nil {
		return nil
	}
	return n.Body
}

// The specified resource was not found.
type NotFoundError struct {
	*core.APIError
	Body *GenericErrorResponse
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body *GenericErrorResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

func (n *NotFoundError) GetBody() *GenericErrorResponse {
	if n == nil {
		return nil
	}
	return n.Body
}
