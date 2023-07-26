package main

import (
	"net/rpc"
	"os"
	"os/signal"

	"github.com/alexjch/go-systelemetry/internal/pkg/comm"
)

func start(s *comm.SysTelemetry) {
	l := comm.EventListener{}
	rpc.Register(&l)
	rpc.Accept(s)
}

func main() {
	shutdown := make(chan os.Signal, 1)
	s, err := comm.NewSysTelemetry("/var/run/systelem.sock")
	if err != nil {
		panic(err)
	}
	defer s.Close()
	go start(s)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
}
