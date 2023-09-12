package wendy

import (
	"encoding/json"
	"log"
)

type (
	Request struct {
		App     string            `json:"app"`
		Module  string            `json:"module"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers,omitempty"`
		Body    *Body             `json:"body,omitempty"`
	}

	Response struct {
		Code    int               `json:"status"`
		Headers map[string]string `json:"headers,omitempty"`
		Body    *Body             `json:"body,omitempty"`
	}

	Body struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	Handler = func(*Request) *Response
)

const (
	JSON = "application/json"
	HTML = "text/html"
	TEXT = "text/plain"
	JS   = "text/javascript"
	CSS  = "text/css"
	FORM = "application/x-www-form-urlencoded"
)

func (r *Response) SetHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}

	r.Headers[key] = value
}

func Ok(body *Body) *Response {
	res := &Response{200, nil, nil}

	if body != nil {
		res.Body = body
	}

	return res
}

func Error(body *Body) *Response {
	res := &Response{500, nil, nil}

	if body != nil {
		res.Body = body
	}

	return res
}

func BadRequest(body *Body) *Response {
	res := &Response{400, nil, nil}

	if body != nil {
		res.Body = body
	}

	return res
}

func Forbidden(body *Body) *Response {
	res := &Response{401, nil, nil}

	if body != nil {
		res.Body = body
	}

	return res
}

func NotAllowed(body *Body) *Response {
	res := &Response{403, nil, nil}

	if body != nil {
		res.Body = body
	}

	return res
}

func NotFound() *Response {
	return &Response{404, nil, nil}
}

func Redirect(url string) *Response {
	headers := make(map[string]string)
	headers["Location"] = url

	return &Response{303, headers, nil}
}

func Invalid(body *Body) *Response {
	res := &Response{409, nil, nil}

	if body != nil {
		res.Body = body
	}

	return res
}

func Json(data interface{}) *Body {
	bs, err := json.Marshal(data)

	if err != nil {
		log.Printf("Turning data into json failed: %v\n", err)
	}

	return &Body{JSON, bs}
}

func Static(encoding string, data []byte) *Body {
	return &Body{encoding, data}
}

func Text(text string) *Body {
	return &Body{TEXT, []byte(text)}
}

func Form(bytes []byte) *Body {
	return &Body{FORM, bytes}
}

func (b *Body) Bind(into interface{}) error {
	return json.Unmarshal(b.Data, into)
}

func (b *Body) Text() string {
	return string(b.Data)
}
