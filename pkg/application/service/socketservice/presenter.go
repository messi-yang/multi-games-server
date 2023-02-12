package socketservice

type Presenter interface {
	OnMessage(jsonObj any)
}
