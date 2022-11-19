package main

import (
	"github.com/Meduzz/wendy"
	"github.com/Meduzz/wendy/example/service"
	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()
	// try out serving modules locally
	logic := wendy.NewLocal(service.ServiceModule())

	/*
		// or try out the remote one (dont forget to start remote/remote.go)
		conn, _ := nuts.Connect()
		rpcSrv := rpc.NewRpc(conn)
		logic := wendy.NewProxy(rpcSrv)
	*/

	// add static paths and what else is needed for the app

	// register wendy api path
	srv.POST("/api", func(ctx *gin.Context) {
		// start by binding request
		req := &wendy.Request{}
		err := ctx.BindJSON(req)

		if err != nil {
			// Somebody obviously sent us a bad body
			ctx.AbortWithStatus(400)
			return
		}

		// call wendy
		res := logic.Handle(ctx, req)

		// start dealing with the response
		if res.Headers != nil {
			for k, v := range res.Headers {
				ctx.Header(k, v)
			}
		}

		if res.Body != nil {
			ctx.Data(res.Code, res.Body.Type, res.Body.Data)
		} else {
			ctx.Status(res.Code)
		}
	})

	srv.Run(":8080")
}
