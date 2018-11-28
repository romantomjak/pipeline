package pipeline

import "testing"

type DummyFilter struct{}

func TestCanCreatePipelineWithNoFilters(t *testing.T) {
	p := NewPipeline()
	if p.filters != nil {
		t.Error("Expected filters to be empty")
	}
}

func TestCanEnqueueDummyFilterIntoPipeline(t *testing.T) {
	f := &DummyFilter{}
	p := NewPipeline()
	p.Enqueue(f)
	if p.filters == nil {
		t.Error("Expected filters to contain DummyFilter")
	}
}