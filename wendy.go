package wendy

import (
	"context"

	"github.com/Meduzz/rpc"
)

type Wendy interface {
	Handle(context.Context, *Request) *Response
}

// Proxy - boot wendy in proxy mode
func NewProxy(srv *rpc.RPC) Wendy {
	// Proxy request to downstream rpc
	return proxy(srv)
}

// Local - boot wendy in local mode
func NewLocal(modules ...*Module) Wendy {
	// Handle the request locally
	return local(modules)
}
