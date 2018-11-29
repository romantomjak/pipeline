package pipeline

type Filter interface {
	Process(in chan Message) chan Message
}

type EchoFilter struct{}

func (ef EchoFilter) Process(in chan Message) chan Message {
	out := make(chan Message)
	go func() {
		defer close(out)
		m := <- in
		out <- m
	}()
	return out
}

type ReverseFilter struct{}

func (rf ReverseFilter) reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (rf ReverseFilter) Process(in chan Message) chan Message {
	out := make(chan Message)
	go func() {
		defer close(out)
		m := <- in
		m.Body = rf.reverse(m.Body)
		out <- m
	}()
	return out
}