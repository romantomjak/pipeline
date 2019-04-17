package pipeline

import (
	"bytes"
	"fmt"
	"testing"
)

type echoFilter struct{}

func (ef echoFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		for m := range in {
			out <- m
		}
		close(out)
	}()
	return out
}

type reverseFilter struct{}

func (rf reverseFilter) reverse(b []byte) []byte {
	r := make([]byte, len(b))
	copy(r, b)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func (rf reverseFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		for m := range in {
			out <- rf.reverse(m)
		}
		close(out)
	}()
	return out
}

func assertBytes(t *testing.T, got, want []byte) {
	t.Helper()
	if bytes.Compare(got, want) != 0 {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestEchoFilterProcess(t *testing.T) {

	t.Run("send once", func(t *testing.T) {
		ef := echoFilter{}
		in := make(chan []byte)
		go func() {
			in <- []byte("Hello World")
			close(in)
		}()
		got := <-ef.Process(in)
		want := []byte("Hello World")
		assertBytes(t, got, want)
	})

	t.Run("send many", func(t *testing.T) {
		ef := echoFilter{}
		in := make(chan []byte)
		expect := 3
		go func() {
			for i := 0; i < expect; i++ {
				in <- []byte(fmt.Sprintf("Hello World #%d", i))
			}
			close(in)
		}()
		out := ef.Process(in)
		for i := 0; i < expect; i++ {
			want := []byte(fmt.Sprintf("Hello World #%d", i))
			got := <-out
			assertBytes(t, got, want)
		}
	})
}

func TestReverseFilterProcess(t *testing.T) {
	rf := reverseFilter{}
	in := make(chan []byte)
	go func() {
		in <- []byte("Hello World")
		close(in)
	}()
	want := []byte("dlroW olleH")
	got := <-rf.Process(in)
	assertBytes(t, got, want)
}
