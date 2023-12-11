package wendy

import (
	"context"
)

type Wendy interface {
	Handle(context.Context, *Request) *Response
}

// FromModules - boot wendy in local mode with an app prefix
func FromModules(app string, modules ...*Module) Wendy {
	if app != "" {
		for _, it := range modules {
			it.SetApp(app)
		}
	}

	// Handle the request locally
	return local(app, modules)
}

// FromModulesNoApp - boot wendy in local mode without an app prefix
func FromModulesNoApp(modules ...*Module) Wendy {
	return local("", modules)
}
