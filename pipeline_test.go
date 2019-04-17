package pipeline

import (
	"testing"
)

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
