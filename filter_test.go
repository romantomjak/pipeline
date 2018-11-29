package pipeline

import (
	"testing"
)

func TestEchoFilterProcess(t *testing.T) {
	ef := EchoFilter{}
	in := make(chan Message)
	m := Message{
		Body: "Hello World",
	}
	go func() {
		in <- m
	}()
	want := Message{
		Body: "Hello World",
	}
	got := <- ef.Process(in)
	if want.Body != got.Body {
		t.Errorf("Echo filter roundtrip error: want='%s', got='%s'", want, got)
	}
}

func TestReverseFilterProcess(t *testing.T) {
	rf := ReverseFilter{}
	in := make(chan Message)
	m := Message{
		Body: "Hello World",
	}
	want := Message{
		Body: "dlroW olleH",
	}
	go func() {
		in <- m
	}()
	got := <- rf.Process(in)
	if want.Body != got.Body {
		t.Errorf("Reverse filter roundtrip error: want='%s', got='%s'", want, got)
	}
}