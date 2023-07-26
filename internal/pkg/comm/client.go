package comm

import (
	"net/rpc"

	"github.com/alexjch/go-dmesg/pkg/dmesg"
)

type Client struct{}

func (c *Client) Notify(e *dmesg.Record) (*Response, error) {
	/* Make this a parameter */
	client, err := rpc.Dial("unix", DEFAULT_SOCKET)
	if err != nil {
		return nil, err
	}
	// Synchronous call
	args := &Args{Event: e}
	var reply Response
	err = client.Call(EVENT_LISTENER_NOTIFY, args, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
