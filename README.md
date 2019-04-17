# pipeline

Pipes and Filters Pattern implemented in Go. Heavily inspired by https://blog.golang.org/pipelines

---

Pipeline is a series of steps (_filters_) connected by channels (_pipes_), where each step is a goroutine.

Each filter exposes a very simple interface: it receives messages on the inbound pipe, processes the message, and publishes the results to the outbound pipe. The pipe connects one filter to the next, sending output messages from one filter to the next.

## Example

To get you started, here's a filter that will reverse incoming message:

```go
package main

import (
	"fmt"

	"github.com/romantomjak/pipeline"
)

type ReverseFilter struct{}

func (rf ReverseFilter) reverse(b []byte) []byte {
	r := make([]byte, len(b))
	copy(r, b)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func (rf ReverseFilter) Process(in chan []byte) chan []byte {
	out := make(chan []byte)
	go func() {
		for m := range in {
			out <- rf.reverse(m)
		}
		close(out)
	}()
	return out
}

func main() {
	rf := ReverseFilter{}

	p := pipeline.NewPipeline()
	p.Add(rf)

	out := p.Process([]byte("Hello World"))
	fmt.Printf("Pipeline result: %s\n", out)
}
```

Output:

```sh
Pipeline result: dlroW olleH
```

## License

Mozilla Public License 2.0
