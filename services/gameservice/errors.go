package gameservice

type errMissingGameRepositoryDependency struct {
}

func (e *errMissingGameRepositoryDependency) Error() string {
	return "The dependency repository.GameRepository is missing in \"gameservice\"."
}

type errGameIsNotInitialized struct {
}

func (e *errGameIsNotInitialized) Error() string {
	return "Game is not initialized."
}
