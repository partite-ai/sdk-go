// Code generated from our API definition. DO NOT EDIT.

package partiteai

import (
	json "encoding/json"
	core "github.com/partite-ai/sdk-go/core"
)

// The request was invalid.
type BadRequestError struct {
	*core.APIError
	Body *Problem
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *Problem
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

func (b *BadRequestError) GetBody() *Problem {
	if b == nil {
		return nil
	}
	return b.Body
}

// The resource is not in a state that allows this, or one with the same unique identifier exists.
type ConflictError struct {
	*core.APIError
	Body *Problem
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	var body *Problem
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

func (c *ConflictError) GetBody() *Problem {
	if c == nil {
		return nil
	}
	return c.Body
}

// The body is too large: over 4 MiB for a JSON body, or over 32 MiB for a multipart body, its files included.
type ContentTooLargeError struct {
	*core.APIError
	Body *Problem
}

func (c *ContentTooLargeError) UnmarshalJSON(data []byte) error {
	var body *Problem
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 413
	c.Body = body
	return nil
}

func (c *ContentTooLargeError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

func (c *ContentTooLargeError) Unwrap() error {
	return c.APIError
}

func (c *ContentTooLargeError) GetBody() *Problem {
	if c == nil {
		return nil
	}
	return c.Body
}

// A boundary of the mesh, or of the access point the caller came through, refused the action. The detail is the refusing rule's message, or a generic one. An admin looks the decision up by its decision_id.
type ForbiddenError struct {
	*core.APIError
	Body *BoundaryProblem
}

func (f *ForbiddenError) UnmarshalJSON(data []byte) error {
	var body *BoundaryProblem
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	f.StatusCode = 403
	f.Body = body
	return nil
}

func (f *ForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Body)
}

func (f *ForbiddenError) Unwrap() error {
	return f.APIError
}

func (f *ForbiddenError) GetBody() *BoundaryProblem {
	if f == nil {
		return nil
	}
	return f.Body
}

// An unexpected error occurred on the server.
type InternalServerError struct {
	*core.APIError
	Body *Problem
}

func (i *InternalServerError) UnmarshalJSON(data []byte) error {
	var body *Problem
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

func (i *InternalServerError) GetBody() *Problem {
	if i == nil {
		return nil
	}
	return i.Body
}

// The Accept header allows none of the formats.
type NotAcceptableError struct {
	*core.APIError
	Body *Problem
}

func (n *NotAcceptableError) UnmarshalJSON(data []byte) error {
	var body *Problem
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

func (n *NotAcceptableError) GetBody() *Problem {
	if n == nil {
		return nil
	}
	return n.Body
}

// The resource was not found.
type NotFoundError struct {
	*core.APIError
	Body *Problem
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body *Problem
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

func (n *NotFoundError) GetBody() *Problem {
	if n == nil {
		return nil
	}
	return n.Body
}

// The body is neither JSON nor a multipart form.
type UnsupportedMediaTypeError struct {
	*core.APIError
	Body *Problem
}

func (u *UnsupportedMediaTypeError) UnmarshalJSON(data []byte) error {
	var body *Problem
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 415
	u.Body = body
	return nil
}

func (u *UnsupportedMediaTypeError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnsupportedMediaTypeError) Unwrap() error {
	return u.APIError
}

func (u *UnsupportedMediaTypeError) GetBody() *Problem {
	if u == nil {
		return nil
	}
	return u.Body
}
