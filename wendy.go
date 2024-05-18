package wendy

import (
	"context"
)

type Wendy interface {
	Handle(context.Context, *Request) *Response
}

// FromModules - boot wendy in local mode with an app prefix
func FromModules(modules ...*Module) Wendy {
	// Handle the request locally
	return local(modules)
}
