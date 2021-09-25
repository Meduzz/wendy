package main

import (
	"github.com/Meduzz/wendy"
	"github.com/Meduzz/wendy/example/service"
	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()
	logic := wendy.NewLocal(service.ServiceModule())

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
		res := logic.Handle(req)

		// start dealing with the response
		if res.Headers != nil {
			for k, v := range res.Headers {
				ctx.Header(k, v)
			}
		}

		if res.Body != nil {
			ctx.JSON(res.Code, res.Body)
		} else {
			ctx.Status(res.Code)
		}
	})

	srv.Run(":8080")
}
