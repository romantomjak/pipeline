package pipeline

import "testing"

func TestPipeline_Enqueue(t *testing.T) {
	ef := EchoFilter{}
	p := NewPipeline()
	p.Enqueue(ef)
	if p.head == nil {
		t.Error("Expected Pipeline head to not be nil")
	}
	if p.tail == nil {
		t.Error("Expected Pipeline tail to not be nil")
	}
}


func TestPipeline_Process(t *testing.T) {
	ef := EchoFilter{}
	p := NewPipeline()
	p.Enqueue(ef)
	want := Message{
		Body: "Hello World",
	}
	got := p.Process(want)
	if want.Body != got.Body {
		t.Errorf("Echo filter roundtrip error: want='%s', got='%s'", want, got)
	}
}