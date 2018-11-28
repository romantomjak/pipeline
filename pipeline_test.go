package pipeline

import "testing"

func TestCanCreatePipelineWithNoFilters(t *testing.T) {
	p := NewPipeline()
	if p.filters != nil {
		t.Error("Expected filters to be empty")
	}
}
