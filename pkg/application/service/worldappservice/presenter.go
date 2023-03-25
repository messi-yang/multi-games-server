package worldappservice

type Presenter interface {
	OnSuccess(jsonObj any)
	OnError(error)
}
