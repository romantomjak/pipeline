package pipeline

import (
	"bytes"
	"testing"
)

func TestEchoFilterProcess(t *testing.T) {
	ef := EchoFilter{}
	in := make(chan []byte)
	go func() {
		in <- []byte("Hello World")
	}()
	got := <- ef.Process(in)
	want := []byte("Hello World")
	if bytes.Compare(got, want) != 0 {
		t.Errorf("Echo filter roundtrip error: want='%s', got='%s'", want, got)
	}
}

func TestReverseFilterProcess(t *testing.T) {
	rf := ReverseFilter{}
	in := make(chan []byte)
	go func() {
		in <- []byte("Hello World")
	}()
	want := []byte("dlroW olleH")
	got := <- rf.Process(in)
	if bytes.Compare(got, want) != 0 {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}