package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/alexjch/go-systelemetry/internal/pkg/types"
)

type Client struct{}

func (c *Client) Request() {
	/* Make this a parameter */
	client, err := rpc.Dial("unix", "/var/run/systelem.sock")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &types.Args{}
	var reply types.Response
	err = client.Call("Plugin.Setup", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Response: %s", reply)
}
