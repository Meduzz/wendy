package wendy

import (
	"context"
	"time"
)

type wendyLocal struct {
	modules []*Module
}

func (w *wendyLocal) Handle(ctx context.Context, req *Request) *Response {
	deadline, ok := ctx.Deadline()

	for _, m := range w.modules {
		if ok && deadline.Before(time.Now()) {
			return &Response{503, nil, nil}
		}

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
