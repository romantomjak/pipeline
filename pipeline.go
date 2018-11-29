package pipeline

type Pipeline struct {
	head chan Message
	tail chan Message
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Enqueue(filter Filter) {
	if p.tail == nil {
		p.head = make(chan Message)
		p.tail = filter.Process(p.head)
	} else {
		p.tail = filter.Process(p.tail)
	}
}

func (p *Pipeline) Process(in Message) (out Message) {
	if p.head == nil {
		return in
	}
	defer close(p.head)
	p.head <- in
	return <- p.tail
}
