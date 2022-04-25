package gameprovider

type errGameProviderHasBeenCreated struct {
}

func (e *errGameProviderHasBeenCreated) Error() string {
	return "Game provider has been created."
}
