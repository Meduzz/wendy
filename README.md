# Wendy

Wendy is a lightweight meta http framework. At the base is a familiar `Handler`-func `func(context.Context, *wendy.Request) *wendy.Response`. What differs is the routing, instead of urls, a `Handler`-func is bound to a module and a name.

You take care of the routing from what ever transport you're using into wendy.

Dont forget to have a look at the premade transports:
* RPC https://github.com/Meduzz/wendy-rpc
* GIN https://github.com/Meduzz/wendy-gin

## Modules

Think of modules as a container for a bunch of logic that is somehow connected. Ex. in a CRUD app the module would be named after the entity and then have a bunch of methods to interact with that entity.

On the module you register your `Handler`-funcs.

For an example module have a look at the [example service](example/service/service.go)

## Request

The request is modeled roughtly after a http request, exept for the routing ofc. It has `Module` and `Method` properties for routing. It has a `Header` property (`map[string]string`). It has a `Body` property, that in turn have its own format.

### Body

The body have 2 properties, `Type` which matches perfectly with the http `Content-Type` and `Data` which is a byte array.

#### Helpers

There are helpers for the most common body types. (Json, Form, Text & Static)

## Response

The response too is modeled after its http counter part. So it has a `Code` property, which matches the http status code. It has `Headers` property, again a `map[string]string`. And it has a `Body` property, of the same type as the request.

### Helpers

There are helpers for the most common responses. (Ok, BadRequest, Error, Forbidden, NotAllowed, NotFound, Invalid & Redirect)