package gameroomservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type Service interface {
	CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error)
	GetAllRooms() []aggregate.GameRoom
	GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error)

	GetUnitMap(gameId uuid.UUID) (*valueobject.UnitMap, error)
	GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (*valueobject.UnitMap, error)

	AddPlayer(gameId uuid.UUID, player entity.Player) error
	RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error

	AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error
	RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error

	TickUnitMap(gameId uuid.UUID) error
	ReviveUnits(gameId uuid.UUID, coords []valueobject.Coordinate) error
}

type serviceImplement struct {
	gameRoomRepository gameroomrepository.Repository
}

var serviceInstance Service = nil

func NewService(gameRoomRepository gameroomrepository.Repository) Service {
	if serviceInstance == nil {
		serviceInstance = &serviceImplement{
			gameRoomRepository: gameRoomRepository,
		}
	}
	return serviceInstance
}

func (gsi *serviceImplement) CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error) {
	unitMap := valueobject.NewUnitMap(mapSize)
	game := entity.NewGame(unitMap)
	gameRoom := aggregate.NewGameRoom(game, make(map[uuid.UUID]entity.Player), make(map[uuid.UUID]valueobject.Area))
	gsi.gameRoomRepository.Add(gameRoom)

	return gameRoom, nil
}

func (gsi *serviceImplement) GetAllRooms() []aggregate.GameRoom {
	return gsi.gameRoomRepository.GetAll()
}

func (gsi *serviceImplement) GetRoom(gameId uuid.UUID) (aggregate.GameRoom, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer rUnlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	return gameRoom, nil
}

func (gsi *serviceImplement) GetUnitMap(gameId uuid.UUID) (*valueobject.UnitMap, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return &valueobject.UnitMap{}, err
	}
	defer rUnlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return &valueobject.UnitMap{}, err
	}

	unitMap := gameRoom.GetUnitMap()
	if err != nil {
		return &valueobject.UnitMap{}, err
	}

	return unitMap, nil
}

func (gsi *serviceImplement) GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (*valueobject.UnitMap, error) {
	rUnlocker, err := gsi.gameRoomRepository.ReadLockAccess(gameId)
	if err != nil {
		return &valueobject.UnitMap{}, err
	}
	defer rUnlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)

	if err != nil {
		return &valueobject.UnitMap{}, err
	}

	unitMap, err := gameRoom.GetUnitMapByArea(area)
	if err != nil {
		return &valueobject.UnitMap{}, err
	}

	return unitMap, nil
}

func (gsi *serviceImplement) AddPlayer(gameId uuid.UUID, player entity.Player) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	err = gameRoom.AddPlayer(player)
	if err != nil {
		return err
	}

	err = gsi.gameRoomRepository.AddPlayer(gameId, player)
	if err != nil {
		return err
	}

	return nil
}

func (gsi *serviceImplement) RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameRoom.RemovePlayer(playerId)

	err = gsi.gameRoomRepository.RemovePlayer(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (gsi *serviceImplement) AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameRoom.AddZoomedArea(playerId, area)

	err = gsi.gameRoomRepository.AddZoomedArea(gameId, playerId, area)
	if err != nil {
		return err
	}

	return nil
}

func (gsi *serviceImplement) RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	gameRoom.RemoveZoomedArea(playerId)

	err = gsi.gameRoomRepository.RemoveZoomedArea(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (gsi *serviceImplement) TickUnitMap(gameId uuid.UUID) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	err = gameRoom.TickUnitMap()
	if err != nil {
		return err
	}

	updatedUnitMap := gameRoom.GetUnitMap()

	err = gsi.gameRoomRepository.UpdateUnitMap(gameId, updatedUnitMap)
	if err != nil {
		return err
	}

	err = gsi.gameRoomRepository.UpdateLastTickedAt(gameId, gameRoom.GetLastTickedAt())
	if err != nil {
		return err
	}

	return nil
}

func (gsi *serviceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return err
	}

	err = gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return err
	}

	updatedUnits, err := gameRoom.GetUnitsWithCoordinates(coordinates)
	if err != nil {
		return err
	}

	err = gsi.gameRoomRepository.UpdateUnits(gameId, coordinates, updatedUnits)
	if err != nil {
		return err
	}

	return nil
}
