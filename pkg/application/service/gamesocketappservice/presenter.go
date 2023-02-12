package gamesocketappservice

type Presenter interface {
	OnMessage(jsonObj any)
}
