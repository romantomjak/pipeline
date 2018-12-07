package pipeline

type Pipeline struct {
	head chan []byte
	tail chan []byte
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Enqueue(filter Filter) {
	if p.tail == nil {
		p.head = make(chan []byte)
		p.tail = filter.Process(p.head)
	} else {
		p.tail = filter.Process(p.tail)
	}
}

func (p *Pipeline) Process(in []byte) (out []byte) {
	if p.head == nil {
		return in
	}
	defer close(p.head)
	p.head <- in
	return <- p.tail
}
