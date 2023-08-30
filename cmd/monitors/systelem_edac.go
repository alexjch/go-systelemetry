// SPDX-License-Identifier: MIT
package main

import (
	"os"
	"os/signal"

	"github.com/alexjch/go-systelemetry/internal/pkg/monitors"
)

func main() {
	var err error
	var edac *monitors.EdacMonitor
	// Handle shutdown
	shutdown := make(chan os.Signal, 1)
	// Instantiate monitor
	if edac, err = monitors.NewEdacMonitor(); err != nil {
		panic(err)
	}
	// Start listening for events
	if err := edac.Listen(); err != nil {
		panic(err)
	}
	// Close when done
	defer edac.Close()
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
}
