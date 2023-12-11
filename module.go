package wendy

import "context"

type (
	Module struct {
		app     string
		name    string
		methods map[string]Handler
	}
)

func NewModule(app, name string) *Module {
	methods := make(map[string]Handler)
	return &Module{app, name, methods}
}

func NewModuleNoApp(name string) *Module {
	methods := make(map[string]Handler)
	return &Module{"", name, methods}
}

func (m *Module) App() string {
	return m.app
}

func (m *Module) SetApp(app string) {
	m.app = app
}

func (m *Module) Name() string {
	return m.name
}

func (m *Module) WithHandler(method string, handler Handler) *Module {
	m.methods[method] = handler
	return m
}

func (m *Module) CanHandle(req *Request) bool {
	if m.app == req.App && m.name == req.Method {
		_, ok := m.methods[req.Method]

		return ok
	}

	return false
}

func (m *Module) Handle(ctx context.Context, req *Request) *Response {
	method, ok := m.methods[req.Method]

	if ok {
		return method(req)
	}

	return NotFound()
}
