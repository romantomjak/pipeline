package pipeline

// Pipeline keeps references to the start and end of pipeline
type Pipeline struct {
	head chan []byte
	tail chan []byte
}

// NewPipeline creates a new execution pipeline
func NewPipeline() *Pipeline {
	return &Pipeline{}
}

// Add adds a new pipeline step
func (p *Pipeline) Add(filter Filter) {
	if p.tail == nil {
		p.head = make(chan []byte)
		p.tail = filter.Process(p.head)
	} else {
		p.tail = filter.Process(p.tail)
	}
}

// Process executes the pipeline
func (p *Pipeline) Process(in []byte) (out []byte) {
	if p.head == nil {
		return in
	}
	defer close(p.head)
	p.head <- in
	return <-p.tail
}
