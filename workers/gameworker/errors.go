package gameworker

type errGameWorkHasBeenCreated struct {
}

func (e *errGameWorkHasBeenCreated) Error() string {
	return "Game worker has been created."
}
