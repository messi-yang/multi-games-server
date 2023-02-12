package presenter

type SocketPresenter interface {
	OnMessage(jsonObj any)
}
