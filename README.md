# pipeline

Pipes and Filters Pattern implemented in Go. Heavily inspired by https://blog.golang.org/pipelines

---

Pipeline is a series of steps (_filters_) connected by channels (_pipes_), where each step is a goroutine.

Each filter exposes a very simple interface: it receives messages on the inbound pipe, processes the message, and publishes the results to the outbound pipe. The pipe connects one filter to the next, sending output messages from one filter to the next.

## Example

```go
package main

import (
    "fmt"
    "github.com/romantomjak/pipeline"
)

func main()  {
    rf := pipeline.ReverseFilter{}

    p := pipeline.NewPipeline()
    p.Enqueue(rf)

    m := "Hello World"

    out := p.Process(m)
    fmt.Printf("Pipeline result message: %s", out)
}
```

## License

Mozilla Public License 2.0
