package gameroommemory

import (
	"sync"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type record struct {
	unitMap  valueobject.UnitMap
	tickedAt time.Time
}

type recordCollection struct {
	records       map[uuid.UUID]*record
	recordLockers map[uuid.UUID]*sync.RWMutex
}

var recordCollectionInstance *recordCollection

func GetRepository() gameroomrepository.Repository {
	if recordCollectionInstance == nil {
		recordCollectionInstance = &recordCollection{
			records:       make(map[uuid.UUID]*record),
			recordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return recordCollectionInstance
	} else {
		return recordCollectionInstance
	}
}

func (gmi *recordCollection) Get(id uuid.UUID) (aggregate.GameRoom, time.Time, error) {
	record, exists := gmi.records[id]
	if !exists {
		return aggregate.GameRoom{}, time.Time{}, gameroomrepository.ErrGameRoomNotFound
	}

	game := entity.NewGameFromExistingEntity(id, record.unitMap)
	gameRoom := aggregate.NewGameRoomWithLastTickedAt(game, record.tickedAt)
	return gameRoom, time.Now(), nil
}

func (gmi *recordCollection) GetAll() []aggregate.GameRoom {
	gameRoom := make([]aggregate.GameRoom, 0)
	for gameId, record := range gmi.records {
		game := entity.NewGameFromExistingEntity(gameId, record.unitMap)
		newGameRoom := aggregate.NewGameRoomWithLastTickedAt(game, record.tickedAt)
		gameRoom = append(gameRoom, newGameRoom)
	}
	return gameRoom
}

func (gmi *recordCollection) GetLastTickedAt(id uuid.UUID) (time.Time, error) {
	record, exists := gmi.records[id]
	if !exists {
		return time.Time{}, gameroomrepository.ErrGameRoomNotFound
	}

	return record.tickedAt, nil
}

func (gmi *recordCollection) UpdateUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate, units []valueobject.Unit) error {
	record, exists := gmi.records[gameId]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	for coordIdx, coord := range coordinates {
		record.unitMap.SetUnit(coord, units[coordIdx])
	}

	return nil
}

func (gmi *recordCollection) UpdateUnitMap(gameId uuid.UUID, unitMap valueobject.UnitMap) error {
	record, exists := gmi.records[gameId]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	record.unitMap = unitMap

	return nil
}

func (gmi *recordCollection) UpdateLastTickedAt(id uuid.UUID, tickedAt time.Time) error {
	record, exists := gmi.records[id]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	record.tickedAt = tickedAt

	return nil
}

func (gmi *recordCollection) Add(gameRoom aggregate.GameRoom) error {
	gameUnitMap := gameRoom.GetUnitMap()
	gmi.records[gameRoom.GetGameId()] = &record{
		unitMap: gameUnitMap,
	}
	gmi.recordLockers[gameRoom.GetGameId()] = &sync.RWMutex{}

	return nil
}

func (gmi *recordCollection) ReadLockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := gmi.recordLockers[gameId]
	if !exists {
		return nil, gameroomrepository.ErrGameRoomLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (gmi *recordCollection) LockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := gmi.recordLockers[gameId]
	if !exists {
		return nil, gameroomrepository.ErrGameRoomLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
