package gameappservice

type Presenter interface {
	OnMessage(jsonObj any) error
}
