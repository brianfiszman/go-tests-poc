package interfaces

type Builder[T any] interface {
	Build() *T
}
