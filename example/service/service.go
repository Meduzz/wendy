package service

import (
	"context"
	"math/rand"

	"github.com/Meduzz/wendy"
)

type (
	Service struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
		Host string   `json:"host"`
		Port int      `json:"port"`
	}
)

func ServiceModule() *wendy.Module {
	module := wendy.NewModule("example", "service")
	return module.
		WithHandler("add", addService).
		WithHandler("remove", removeService).
		WithHandler("list", listServices).
		WithHandler("find", findService)
}

var (
	registry = make(map[string][]*Service)
)

func addService(ctx context.Context, req *wendy.Request) *wendy.Response {
	svc := &Service{}
	req.Body.Bind(svc)

	list, ok := registry[svc.Name]

	if !ok {
		list = []*Service{svc}
	} else {
		list = append(list, svc)
	}

	registry[svc.Name] = list

	resp := wendy.Ok(wendy.Json(len(list)))

	return resp
}

func removeService(ctx context.Context, req *wendy.Request) *wendy.Response {
	svc := &Service{}
	req.Body.Bind(svc)

	list, ok := registry[svc.Name]

	if !ok {
		resp := wendy.Ok(wendy.Json(0))

		return resp
	}

	list = filter(list, svc)
	registry[svc.Name] = list

	resp := wendy.Ok(wendy.Json(len(list)))

	return resp
}

func listServices(ctx context.Context, req *wendy.Request) *wendy.Response {
	name := ""
	req.Body.Bind(&name)

	list, ok := registry[name]

	if !ok {
		resp := wendy.NotFound()

		return resp
	}

	resp := wendy.Ok(wendy.Json(list))

	return resp
}

func findService(ctx context.Context, req *wendy.Request) *wendy.Response {
	name := ""
	req.Body.Bind(&name)

	list, ok := registry[name]

	if !ok {
		resp := wendy.NotFound()

		return resp
	}

	if len(list) == 0 {
		return wendy.NotFound()
	}

	idx := rand.Intn(len(list))
	svc := list[idx]

	resp := wendy.Ok(wendy.Json(svc))

	return resp
}

func filter(list []*Service, svc *Service) []*Service {
	copy := make([]*Service, 0)

	for _, s := range list {
		if s.Host != svc.Host && s.Port != svc.Port {
			copy = append(copy, s)
		}
	}

	return copy
}
