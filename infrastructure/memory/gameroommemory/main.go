package gameroommemory

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type gameRoomMemory struct {
	gameRoomMap map[uuid.UUID]aggregate.GameRoom
}

var gameRoomMemoryInstance *gameRoomMemory

func GetGameRoomMemory() gameroomrepository.GameRoomRepository {
	if gameRoomMemoryInstance == nil {
		gameRoomMemoryInstance = &gameRoomMemory{
			gameRoomMap: make(map[uuid.UUID]aggregate.GameRoom),
		}
		return gameRoomMemoryInstance
	} else {
		return gameRoomMemoryInstance
	}
}

func (gmi *gameRoomMemory) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	gameRoom, exists := gmi.gameRoomMap[id]
	if !exists {
		return aggregate.GameRoom{}, gameroomrepository.ErrGameRoomNotFound
	}
	return gameRoom, nil
}

func (gmi *gameRoomMemory) GetAll() []aggregate.GameRoom {
	gameRooms := make([]aggregate.GameRoom, 0)
	for _, gameRoom := range gmi.gameRoomMap {
		gameRooms = append(gameRooms, gameRoom)
	}
	return gameRooms
}

func (gmi *gameRoomMemory) UpdateUnit(gameId uuid.UUID, coordinate valueobject.Coordinate, unit valueobject.Unit) error {
	gameRoom := gmi.gameRoomMap[gameId]
	gameRoom.UpdateUnit(coordinate, unit)

	return nil
}

func (gmi *gameRoomMemory) UpdateUnitMap(gameId uuid.UUID, unitMap [][]valueobject.Unit) error {
	gameRoom := gmi.gameRoomMap[gameId]
	gameRoom.UpdateUnitMap(unitMap)

	return nil
}

func (gmi *gameRoomMemory) Add(gameRoom aggregate.GameRoom) error {
	gmi.gameRoomMap[gameRoom.GetGameId()] = gameRoom

	return nil
}
