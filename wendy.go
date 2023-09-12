package wendy

import (
	"context"
)

type Wendy interface {
	Handle(context.Context, *Request) *Response
}

// Local - boot wendy in local mode
func NewLocal(app string, modules ...*Module) Wendy {
	// Handle the request locally
	return local(app, modules)
}
