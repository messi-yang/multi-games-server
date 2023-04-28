package domainmodel

type ValueObject[T any] interface {
	IsEqual(T) bool
}
