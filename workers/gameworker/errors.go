package gameworker

type errGameStoreIsNotSet struct {
}

func (e *errGameStoreIsNotSet) Error() string {
	return "Game store is not set."
}
