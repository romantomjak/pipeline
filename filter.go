package pipeline

type Filter interface {
	Process(in chan Message) chan Message
}

type EchoFilter struct{}

func (ef EchoFilter) Process(in chan Message) chan Message {
	out := make(chan Message)
	go func() {
		m := <- in
		out <- m
	}()
	return out
}
