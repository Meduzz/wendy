package wendy

import (
	"fmt"
	"log"

	"github.com/Meduzz/rpc"
	"github.com/nats-io/nats.go"
)

type (
	wendyProxy struct {
		srv *rpc.RPC
	}
)

func (w *wendyProxy) Handle(req *Request) *Response {
	topic := fmt.Sprintf("%s.%s.%s", req.App, req.Module, req.Method)
	resCtx, err := w.srv.Request(topic, req, 10)

	if err != nil {
		if err == nats.ErrTimeout {
			log.Printf("Request to %s timed out\n", topic)
			return &Response{503, nil, nil}
		} else {
			log.Printf("Request to %s threw error: %v\n", topic, err)
			return Error(nil)
		}
	}

	res := &Response{}
	err = resCtx.Bind(res)

	if err != nil {
		log.Printf("Parsing response threw error: %v\n", err)
		return Error(nil)
	}

	return res
}

func proxy(srv *rpc.RPC) Wendy {
	return &wendyProxy{srv}
}
