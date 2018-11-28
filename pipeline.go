package pipeline

type Pipeline struct {
	filters []Filter
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		filters: nil,
	}
}
