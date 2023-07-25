package interfaces

type ListBuilder[T any] interface {
	WithCustomFunc(customFunc func(*Builder[T])) *ListBuilder[T]
	Build() []*T
}
