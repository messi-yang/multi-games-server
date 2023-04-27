package valueobject

type ValueObject[T any] interface {
	IsEqual(T) bool
}
