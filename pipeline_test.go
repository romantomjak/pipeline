package pipeline

import (
	"testing"
)

func TestEmptyPipelineReturnsSameMessage(t *testing.T) {
	p := NewPipeline()
	want := []byte("Hello World")
	got := p.Process(want)
	assertByteBuffers(t, got, want)
}


func TestPipelineProcessEchoFilter(t *testing.T) {
	ef := EchoFilter{}
	p := NewPipeline()
	p.Enqueue(ef)
	want := []byte("Hello World")
	got := p.Process(want)
	assertByteBuffers(t, got, want)
}

func TestPipelineProcessMultipleFilters(t *testing.T) {
	ef := EchoFilter{}
	rf := ReverseFilter{}
	p := NewPipeline()
	p.Enqueue(ef)
	p.Enqueue(rf)
	got := p.Process([]byte("Hello World"))
	want := []byte("dlroW olleH")
	assertByteBuffers(t, got, want)
}