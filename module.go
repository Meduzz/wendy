package wendy

import "context"

type (
	Module struct {
		app     string
		name    string
		methods map[string]HandlerFunc
	}

	HandlerFunc = func(context.Context, *Request) *Response
)

func NewModule(app, name string) *Module {
	methods := make(map[string]HandlerFunc)
	return &Module{app, name, methods}
}

func (m *Module) App() string {
	return m.app
}

func (m *Module) Name() string {
	return m.name
}

func (m *Module) Method(name string) HandlerFunc {
	return m.methods[name]
}

func (m *Module) WithHandler(method string, handler HandlerFunc, middlewares ...Middleware) *Module {
	if len(middlewares) > 0 {
		chain := Chain(middlewares, handler)
		m.methods[method] = chain.Handle
	} else {
		m.methods[method] = handler
	}

	return m
}

func (m *Module) Handle(ctx context.Context, req *Request) *Response {
	method := m.Method(req.Method)

	if method == nil {
		return NotFound()
	}

	return method(ctx, req)
}
