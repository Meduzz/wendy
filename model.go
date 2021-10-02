package wendy

import (
	"encoding/json"
)

type (
	Request struct {
		Module  string            `json:"module"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers,omitempty"`
		Body    json.RawMessage   `json:"body,omitempty"`
	}

	Response struct {
		Code    int               `json:"status"`
		Headers map[string]string `json:"headers,omitempty"`
		Body    json.RawMessage   `json:"body,omitempty"`
	}

	Handler = func(*Request) *Response
)

func (r *Request) Bind(into interface{}) error {
	return json.Unmarshal(r.Body, into)
}

func (r *Request) SetBody(any interface{}) error {
	bs, err := json.Marshal(any)

	if err != nil {
		return err
	}

	r.Body = json.RawMessage(bs)

	return nil
}

func (r *Response) SetBody(any interface{}) error {
	bs, err := json.Marshal(any)

	if err != nil {
		return err
	}

	r.Body = json.RawMessage(bs)

	return nil
}

func (r *Response) SetHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}

	r.Headers[key] = value
}

func (r *Response) Bind(into interface{}) error {
	return json.Unmarshal(r.Body, into)
}

func Ok(body interface{}) *Response {
	res := &Response{200, nil, nil}

	if body != nil {
		res.SetBody(body)
	}

	return res
}

func Error(body interface{}) *Response {
	res := &Response{500, nil, nil}

	if body != nil {
		res.SetBody(body)
	}

	return res
}

func BadRequest(body interface{}) *Response {
	res := &Response{400, nil, nil}

	if body != nil {
		res.SetBody(body)
	}

	return res
}

func Forbidden(body interface{}) *Response {
	res := &Response{403, nil, nil}

	if body != nil {
		res.SetBody(body)
	}

	return res
}

func Authorize() *Response {
	return &Response{401, nil, nil}
}

func NotFound() *Response {
	return &Response{404, nil, nil}
}
