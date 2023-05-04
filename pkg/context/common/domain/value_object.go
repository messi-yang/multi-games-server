package domain

type ValueObject[T any] interface {
	IsEqual(T) bool
}
