package wendy

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

func (m *Module) App() string {
	return m.app
}

func (m *Module) Name() string {
	return m.name
}

func (m *Module) Method(name string) Handler {
	return m.methods[name]
}

func (m *Module) WithHandler(method string, handler Handler) *Module {
	m.methods[method] = handler
	return m
}
