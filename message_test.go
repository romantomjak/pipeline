package pipeline

import "testing"

func TestCanCreateEmptyMessage(t *testing.T) {
	m := Message{}
	if m.Body != "" {
		t.Errorf("Expected message body to be empty, but got '%s'", m.Body)
	}
}
