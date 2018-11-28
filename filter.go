package pipeline

type Filter interface {
	Process(in chan Message, out chan Message)
}

type EchoFilter struct{}

func (ef *EchoFilter) Process(in chan Message, out chan Message) {
	m := <- in
	out <- m
}
