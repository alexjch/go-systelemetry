package types

type Args struct{}
type Response string
type Result []string

type Plugin struct {
}

func (p Plugin) Setup(args Args, reply *Response) error {
	*reply = "Hello"
	return nil
}

func (p Plugin) Process(args Args, reply *Result) error {
	return nil
}
