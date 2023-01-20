package appservice

type Presenter interface {
	OnSuccess(jsonObj any)
}
