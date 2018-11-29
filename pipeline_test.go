package pipeline

import "testing"

func TestEmptyPipelineReturnsSameMessage(t *testing.T) {
	want := Message{
		Body: "Hello World",
	}
	p := NewPipeline()
	got := p.Process(want)
	if want != got {
		t.Errorf("Pipeline unexpectedly modified message: want=%s, got=%s", want, got)
	}
}


func TestPipelineProcessEchoFilter(t *testing.T) {
	ef := EchoFilter{}
	p := NewPipeline()
	p.Enqueue(ef)
	want := Message{
		Body: "Hello World",
	}
	got := p.Process(want)
	if want != got {
		t.Errorf("Pipeline process error: want=%s, got=%s", want, got)
	}
}