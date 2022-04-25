package gameservice

type errMissingGameDAODependency struct {
}

func (e *errMissingGameDAODependency) Error() string {
	return "The dependency gamedao.GameDAO is missing in \"gameservice\"."
}

type errGameIsNotInitialized struct {
}

func (e *errGameIsNotInitialized) Error() string {
	return "Game is not initialized."
}
