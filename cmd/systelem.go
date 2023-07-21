package main

import (
	"net/rpc"
	"os"
	"os/signal"

	"github.com/alexjch/go-systelemetry/internal/pkg/server"
	"github.com/alexjch/go-systelemetry/internal/pkg/types"
)

func start(s *server.SysTelem) {
	plugin := types.Plugin{}
	rpc.Register(plugin)
	rpc.Accept(s)
}

func main() {
	shutdown := make(chan os.Signal, 1)
	s, err := server.NewSysTelem("/var/run/systelem.sock")
	if err != nil {
		panic(err)
	}
	defer s.Close()
	go start(s)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
}
