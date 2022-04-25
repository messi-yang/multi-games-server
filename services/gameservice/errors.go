package gameservice

type errGameServiceHasBeenCreated struct {
}

func (e *errGameServiceHasBeenCreated) Error() string {
	return "Game service has been created."
}
