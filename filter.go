package pipeline

// Filter is an interface that all filters must implement
type Filter interface {
	Process(in chan []byte) chan []byte
}
