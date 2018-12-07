package pipeline

import (
	"bytes"
	"testing"
)

func assertByteBuffers(t *testing.T, got, want []byte) {
	t.Helper()
	if bytes.Compare(got, want) != 0 {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestEchoFilterProcess(t *testing.T) {
	ef := EchoFilter{}
	in := make(chan []byte)
	go func() {
		in <- []byte("Hello World")
	}()
	got := <- ef.Process(in)
	want := []byte("Hello World")
	assertByteBuffers(t, got, want)
}

func TestReverseFilterProcess(t *testing.T) {
	rf := ReverseFilter{}
	in := make(chan []byte)
	go func() {
		in <- []byte("Hello World")
	}()
	want := []byte("dlroW olleH")
	got := <- rf.Process(in)
	assertByteBuffers(t, got, want)
}