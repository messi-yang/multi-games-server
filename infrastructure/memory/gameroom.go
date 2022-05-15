package memory

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type GameRoomMemoryRepository interface {
	Add(aggregate.GameRoom) error
	UpdateGameUnit(uuid.UUID, valueobject.Coordinate, valueobject.GameUnit) error
	UpdateGameUnitMatrix(uuid.UUID, [][]valueobject.GameUnit) error
	Get(uuid.UUID) (aggregate.GameRoom, error)
}

type gameRoomMemoryRepositoryImpl struct {
	gameRoomMap map[uuid.UUID]aggregate.GameRoom
}

var gameRoomMemoryRepository GameRoomMemoryRepository

func NewGameRoomMemoryRepository() GameRoomMemoryRepository {
	if gameRoomMemoryRepository == nil {
		gameRoomMemoryRepository = &gameRoomMemoryRepositoryImpl{
			gameRoomMap: make(map[uuid.UUID]aggregate.GameRoom),
		}
		return gameRoomMemoryRepository
	} else {
		return gameRoomMemoryRepository
	}
}

func (gmi *gameRoomMemoryRepositoryImpl) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	return gmi.gameRoomMap[id], nil
}

func (gmi *gameRoomMemoryRepositoryImpl) UpdateGameUnit(gameId uuid.UUID, coordinate valueobject.Coordinate, gameUnit valueobject.GameUnit) error {
	gameRoom := gmi.gameRoomMap[gameId]
	gameRoom.UpdateGameUnit(coordinate, gameUnit)

	return nil
}

func (gmi *gameRoomMemoryRepositoryImpl) UpdateGameUnitMatrix(gameId uuid.UUID, gameUnitMatrix [][]valueobject.GameUnit) error {
	gameRoom := gmi.gameRoomMap[gameId]
	gameRoom.UpdateGameUnitMatrix(gameUnitMatrix)

	return nil
}

func (gmi *gameRoomMemoryRepositoryImpl) Add(gameRoom aggregate.GameRoom) error {
	gmi.gameRoomMap[gameRoom.GetGameId()] = gameRoom

	return nil
}
