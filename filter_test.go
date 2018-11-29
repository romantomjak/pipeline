package pipeline

import (
	"testing"
)

func TestEchoFilterProcess(t *testing.T) {
	ef := EchoFilter{}
	in := make(chan Message)
	want := Message{
		Body: "Hello World",
	}
	go func() {
		in <- want
	}()
	got := <- ef.Process(in)
	if want.Body != got.Body {
		t.Errorf("Echo filter roundtrip error: want='%s', got='%s'", want, got)
	}
}