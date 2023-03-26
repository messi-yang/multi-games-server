package worldapiservice

type Presenter interface {
	OnSuccess(jsonObj any)
	OnError(error)
}
