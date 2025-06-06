package shared

type Event interface {
	Name() string
	Data() any
}
