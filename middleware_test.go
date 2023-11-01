package wendy

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

type (
	myMiddleware struct {
		text string
	}
)

func TestMiddleware(t *testing.T) {
	req := &Request{}
	req.App = ""
	req.Module = "test"
	req.Method = "test"
	req.Body = Text("world")

	t.Run("with no middlewares", func(t *testing.T) {
		emptySubject := Chain(nil, handlerFunc)

		res := emptySubject.Handle(context.Background(), req)

		if res == nil {
			t.Error("res was nil")
		}

		if res.Body == nil {
			t.Error("res.body was nil")
		}

		if res.Code != 200 {
			t.Errorf("code was not 200 but %d", res.Code)
		}

		if res.Body.Type != TEXT {
			t.Errorf("body type was not text but %s", res.Body.Type)
		}

	})

	t.Run("with one middleware", func(t *testing.T) {
		m := &myMiddleware{"m1"}
		ms := make([]Middleware, 0)
		ms = append(ms, m)

		subject := Chain(ms, handlerFunc)
		res := subject.Handle(context.Background(), req)

		if res == nil {
			t.Error("res was nil")
		}

		if res.Code != 200 {
			t.Errorf("code was not 200 but %d", res.Code)
		}

		if res.Body == nil && res.Body.Type != TEXT {
			t.Error("Body did not match expectations")
		}

		if string(res.Body.Data) != "HELLO WORLD!" {
			t.Errorf("response was not correct, was %s", string(res.Body.Data))
		}
	})

	t.Run("it's all gone wrong", func(t *testing.T) {
		req := &Request{}
		req.Body = Json("blaha")

		m := &myMiddleware{"m1"}
		ms := make([]Middleware, 0)
		ms = append(ms, m)

		subject := Chain(ms, handlerFunc)
		res := subject.Handle(context.Background(), req)

		if res == nil {
			t.Error("res was empty")
		}

		if res.Code != 404 {
			t.Errorf("response was not 404 byt %d", res.Code)
		}
	})
}

func handlerFunc(ctx context.Context, req *Request) *Response {
	println("hello from handlerFunc")
	if req.Body != nil && req.Body.Type == TEXT {
		greeting := string(req.Body.Data)
		return Ok(Text(fmt.Sprintf("Hello %s!", greeting)))
	} else {
		return NotFound()
	}
}

func (m *myMiddleware) Handle(ctx context.Context, req *Request, next HandlerFunc) *Response {
	println(fmt.Sprintf("hello from middle ware %s - before", m.text))
	if req.Body != nil && req.Body.Type == TEXT {
		req.Body.Data = []byte(strings.ToUpper(string(req.Body.Data)))
	}

	res := next(ctx, req)

	if res.Body != nil && res.Body.Type == TEXT {
		res.Body.Data = []byte(strings.ToUpper(string(res.Body.Data)))
	}

	println(fmt.Sprintf("hello from middle ware %s - after", m.text))

	return res
}
