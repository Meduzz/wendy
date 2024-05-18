package wendy

type (
	Module struct {
		name    string
		methods map[string]Handler
	}
)

func NewModule(name string) *Module {
	methods := make(map[string]Handler)
	return &Module{name, methods}
}

func NewModuleNoPrefix(name string) *Module {
	methods := make(map[string]Handler)
	return &Module{name, methods}
}

func (m *Module) Name() string {
	return m.name
}

func (m *Module) WithHandler(method string, handler Handler) *Module {
	m.methods[method] = handler
	return m
}

func (m *Module) CanHandle(req *Request) bool {
	if m.name == req.Module {
		_, ok := m.methods[req.Method]

		return ok
	}

	return false
}

func (m *Module) Handle(req *Request) *Response {
	method, ok := m.methods[req.Method]

	if ok {
		return method(req)
	}

	return NotFound()
}
