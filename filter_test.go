package pipeline

import (
	"testing"
)

func TestEchoFilterRoundtrip(t *testing.T) {
	ef := &EchoFilter{}
	in := make(chan Message)
	out := make(chan Message)
	want := Message{
		Body: "Hello World",
	}

	go func() {
		in <- want
	}()

	go func() {
		ef.Process(in, out)
	}()

	got := <- out
	if want.Body != got.Body {
		t.Errorf("Echo filter roundtrip error: want='%s', got='%s'", want, got)
	}
}