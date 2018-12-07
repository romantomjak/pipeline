package pipeline

import (
	"bytes"
	"testing"
)

func TestEmptyPipelineReturnsSameMessage(t *testing.T) {
	p := NewPipeline()
	want := []byte("Hello World")
	got := p.Process(want)
	if bytes.Compare(got, want) != 0 {
		t.Errorf("Pipeline unexpectedly modified message: want=%s, got=%s", want, got)
	}
}


func TestPipelineProcessEchoFilter(t *testing.T) {
	ef := EchoFilter{}
	p := NewPipeline()
	p.Enqueue(ef)
	want := []byte("Hello World")
	got := p.Process(want)
	if bytes.Compare(got, want) != 0 {
		t.Errorf("Pipeline process error: want=%s, got=%s", want, got)
	}
}

func TestPipelineProcessMultipleFilters(t *testing.T) {
	ef := EchoFilter{}
	rf := ReverseFilter{}
	p := NewPipeline()
	p.Enqueue(ef)
	p.Enqueue(rf)
	got := p.Process([]byte("Hello World"))
	want := []byte("dlroW olleH")
	if bytes.Compare(got, want) != 0 {
		t.Errorf("Pipeline process error: want=%s, got=%s", want, got)
	}
}