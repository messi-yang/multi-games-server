package domain

type Entity[T any] interface {
	GetId() T
}
