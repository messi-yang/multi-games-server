package gameservice

type errMissingGameRoomMemoryRepositoryDependency struct {
}

func (e *errMissingGameRoomMemoryRepositoryDependency) Error() string {
	return "The dependency repository.GameRoomMemoryRepository is missing in \"gameservice\"."
}

type errGameIsNotInitialized struct {
}

func (e *errGameIsNotInitialized) Error() string {
	return "Game is not initialized."
}
