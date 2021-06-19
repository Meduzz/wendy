package wendy

type wendyLocal struct {
	modules []*Module
}

func (w *wendyLocal) Handle(req *Request) *Response {
	for _, m := range w.modules {
		if m.Name() == req.Module {
			method := m.Method(req.Method)

			if method != nil {
				return method(req)
			}
		}
	}

	return NotFound()
}

func local(modules []*Module) Wendy {
	return &wendyLocal{modules}
}
