package wendy

import (
	"context"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	Middleware interface {
		Handle(ctx context.Context, req *Request, next HandlerFunc) *Response
	}

	middleware struct {
		middlewares []Middleware
		handler     HandlerFunc
	}

	MiddlewareFunc func(context.Context, *Request, func(context.Context, *Request) *Response) *Response
)

func Chain(middlewares []Middleware, handler HandlerFunc) Wendy {
	return &middleware{middlewares, handler}
}

func (m *middleware) Handle(ctx context.Context, req *Request) *Response {
	chain := m.chain()

	res := chain(ctx, req)

	if res == nil {
		println("res was still nil")
	}

	return res
}

func (m *middleware) chain() HandlerFunc {
	if len(m.middlewares) == 0 {
		return m.nextHandler(m.handler)
	}

	head := slice.Head(m.middlewares)
	tail := slice.Tail(m.middlewares)

	return m.nextMiddleware(head, tail, m.handler)
}

func (m *middleware) nextMiddleware(head Middleware, tail []Middleware, handler HandlerFunc) HandlerFunc {
	return func(ctx context.Context, req *Request) *Response {
		var next HandlerFunc

		if len(tail) == 0 {
			next = m.nextHandler(handler)
		} else {
			h := slice.Head(tail)
			t := slice.Tail(tail)

			next = m.nextMiddleware(h, t, handler)
		}

		return head.Handle(ctx, req, next)
	}
}

func (m *middleware) nextHandler(handler HandlerFunc) HandlerFunc {
	return func(ctx context.Context, req *Request) *Response {
		return handler(ctx, req)
	}
}
