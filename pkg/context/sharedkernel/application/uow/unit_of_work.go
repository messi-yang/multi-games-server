package uow

type Uow[Transaction any] interface {
	GetTransaction() Transaction
	Rollback()
	Commit()
}
