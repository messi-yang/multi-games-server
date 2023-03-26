package gameapiservice

type Presenter interface {
	OnMessage(jsonObj any) error
}
