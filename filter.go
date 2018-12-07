package pipeline

type Filter interface {
	Process(in chan []byte) chan []byte
}

type EchoFilter struct{}

func (ef EchoFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		m := <- in
		out <- m
	}()
	return out
}

type ReverseFilter struct{}

func (rf ReverseFilter) reverse(b []byte) []byte {
	r := make([]byte, len(b))
	copy(r, b)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func (rf ReverseFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		m := <- in
		out <- rf.reverse(m)
	}()
	return out
}