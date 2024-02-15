package domain

type Aggregate[T any] interface {
	GetId() T
}
