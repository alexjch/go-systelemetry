// SPDX-License-Identifier: MIT
package comm

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	/* Systemd socket activation file descriptor */
	LISTEN_FDS_START uintptr = 3
)

type SysTelemetry struct {
	net.Listener
}

func NewSysTelemetry(s string) (*SysTelemetry, error) {
	var l net.Listener
	var err error

	if os.Getenv("LISTEN_PID") == strconv.Itoa(os.Getpid()) {
		// systemd run
		fmt.Println("Running from systemd")
		f := os.NewFile(LISTEN_FDS_START, "")
		l, err = net.FileListener(f)
		if err != nil {
			return nil, err
		}
	} else {
		// manual run
		fmt.Println("Running manually")
		l, err = net.Listen("unix", DEFAULT_SOCKET)
		if err != nil {
			return nil, err
		}
	}
	return &SysTelemetry{Listener: l}, nil
}
