package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	LISTEN_FDS_START uintptr = 3
	/* Make this a parameter */
	DEFAULT_SOCKET string = "/var/run/systelem.sock"
)

type SysTelem struct {
	net.Listener
}

func NewSysTelem(s string) (*SysTelem, error) {
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
	return &SysTelem{Listener: l}, nil
}
