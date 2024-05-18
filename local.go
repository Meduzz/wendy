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

	if ok && deadline.Before(time.Now()) {
		return &Response{503, nil, nil}
	}

	for _, m := range w.modules {
		if m.CanHandle(req) {
			return m.Handle(req)
		}
	}

	return NotFound()
}

func local(modules []*Module) Wendy {
	return &wendyLocal{modules}
}
