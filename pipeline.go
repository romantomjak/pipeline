package pipeline

type Pipeline struct {
	filters []Filter
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		filters: nil,
	}
}

func (p *Pipeline) Enqueue(filter Filter) {
	p.filters = append(p.filters, filter)
}
