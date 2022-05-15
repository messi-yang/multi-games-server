package gameservice

type errGameIsNotInitialized struct {
}

func (e *errGameIsNotInitialized) Error() string {
	return "Game is not initialized."
}
