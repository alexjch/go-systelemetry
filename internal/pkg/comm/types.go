package comm

import "github.com/alexjch/go-dmesg/pkg/dmesg"

const (
	EVENT_LISTENER_NOTIFY = "EventListener.Notify"
)

type Args struct {
	Event *dmesg.Record
}

type Response struct {
	Error bool
	Msg   string
}

type EventListener struct{}

func (e *EventListener) Notify(args *Args, response *Response) error {
	*response = Response{
		Error: false,
		Msg:   "OK",
	}
	return nil
}
