package main

import (
	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/wendy"
	"github.com/Meduzz/wendy/example/service"
)

func main() {
	conn, err := nuts.Connect()

	if err != nil {
		panic(err)
	}

	err = wendy.ServeModules(conn, "", service.ServiceModule())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	block.Block(func() error {
		return conn.Drain()
	})
}
