package pipeline

import (
	"testing"
)

type echoFilter struct{}

func (ef echoFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		m := <-in
		out <- m
	}()
	return out
}

type reverseFilter struct{}

func (rf reverseFilter) reverse(b []byte) []byte {
	r := make([]byte, len(b))
	copy(r, b)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func (rf reverseFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		m := <-in
		out <- rf.reverse(m)
	}()
	return out
}

func TestEmptyPipelineReturnsSameMessage(t *testing.T) {
	p := NewPipeline()
	want := []byte("Hello World")
	got := p.Process(want)
	assertBytes(t, got, want)
}

func TestPipelineProcessEchoFilter(t *testing.T) {
	ef := echoFilter{}
	p := NewPipeline()
	p.Add(ef)
	want := []byte("Hello World")
	got := p.Process(want)
	assertBytes(t, got, want)
}

func TestPipelineProcessMultipleFilters(t *testing.T) {
	ef := echoFilter{}
	rf := reverseFilter{}
	p := NewPipeline()
	p.Add(ef)
	p.Add(rf)
	got := p.Process([]byte("Hello World"))
	want := []byte("dlroW olleH")
	assertBytes(t, got, want)
}
