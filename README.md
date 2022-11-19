# wendy
Another web framework, how many have I created by now?

This one is slightly different however. It has heavy focus on http rpc, so much that all rpc can go through a single endpoint. All requests are expected to follow a simple base protocol. It still leaves up to the implementor to wire things together.

But that's the kind of boilerplate you generate... right?

Example wiring:

    import (
        "github.com/Meduzz/wendy"
        "github.com/gin-gonic/gin"
    )

    func main() {
        srv := gin.Default()
        logic := wendy.NewLocal(<modules>)|wendy.NewProxy(rpc)
        
        srv.POST("/api", func(ctx *gin.Context) {
            req := &wendy.Request{Context: &wendy.Context{}}
            ctx.BindJSON(req)

            res := logic.Handle(req) // this is wendy

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

## Modules

A module encapsulates a bunch of logic that is somehow connected. So in a CRUD app the module would be named after the entity and then have a bunch of methods to interact with that entity.

On the module you register RPC handler functions.

For an example module have a look at the [example service](example/service/service.go)

## Run modes

You can select between 2 run modes, the most basic one is `wendy.Local` which means you never leave the web service. The second mode is `wendy.Proxy` which turns the web service into a thin proxy, sending request over nats to well known topics `<app>.<module>.<method>`. This forces you to move your logic over into a RPC service but allows you to break it appart even further into RPC methods if needed later.

## The imagined use case

The way I imagine this framework will be used is like this:

1. You start up PoCing a service, you put all your logic in the service together with any statics you might like. You'd use the `wendy.Local` method, which takes all your modules and mount the handler it returns on `/api`.

2. As your service grow, or you iterate it, you'd break it out into a web service and a number RPC services (one per module), listening on nats on topic `<module>.>`. You would switch from the `wendy.Local` handler to the `wendy.Proxy` handler. Inside the RPC service, you can still use `wendy.Module` to handle your "routing". This lets you scale each module individually. This also opens up for writing RPC services in other languages than go.

3. As your service continue to grow, you break your modules into individual RPC functions, listening on topic `<app>.<module>.<method>`. Your web service would continue using `wendy.Proxy`. This lets you scale each function individually.

## Extension points

Since you are in charge of wiring the web service/proxy together. You can use middlewares there to achieve things.

Further down the stack, the RPC handler function is pretty simple and could simply be wrapped.
