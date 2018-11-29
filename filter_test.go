package pipeline

import (
	"testing"
)

func TestEchoFilterProcess(t *testing.T) {
	ef := EchoFilter{}
	in := make(chan Message)
	go func() {
		in <- "Hello World"
	}()
	want := "Hello World"
	got := <- ef.Process(in)
	if want != got {
		t.Errorf("Echo filter roundtrip error: want='%s', got='%s'", want, got)
	}
}

func TestReverseFilterProcess(t *testing.T) {
	rf := ReverseFilter{}
	in := make(chan Message)
	m := "Hello World"
	want := "dlroW olleH"
	go func() {
		in <- m
	}()
	got := <- rf.Process(in)
	if want != got {
		t.Errorf("Reverse filter roundtrip error: want='%s', got='%s'", want, got)
	}
}