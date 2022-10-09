package domainservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameRoomDomainService interface {
	CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error)
	GetAllGameRooms() []aggregate.GameRoom
	GetGameRoom(gameId uuid.UUID) (aggregate.GameRoom, error)

	AddPlayerToGameRoom(gameId uuid.UUID, player entity.Player) (aggregate.GameRoom, error)
	RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error)

	AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (aggregate.GameRoom, error)
	RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error)

	TickUnitMapInGame(gameId uuid.UUID) (aggregate.GameRoom, error)
	ReviveUnitsInGame(gameId uuid.UUID, coords []valueobject.Coordinate) (aggregate.GameRoom, error)
}

type gameRoomDomainServiceImplement struct {
	gameRoomRealtimeRepository repository.GameRoomRealtimeRepository
}

type GameRoomDomainServiceConfiguration struct {
	GameRoomRealtimeRepository repository.GameRoomRealtimeRepository
}

func NewGameRoomDomainService(coniguration GameRoomDomainServiceConfiguration) GameRoomDomainService {
	return &gameRoomDomainServiceImplement{
		gameRoomRealtimeRepository: coniguration.GameRoomRealtimeRepository,
	}
}

func (gsi *gameRoomDomainServiceImplement) CreateGameRoom(mapSize valueobject.MapSize) (aggregate.GameRoom, error) {
	unitMap := valueobject.NewUnitMap(mapSize)
	game := entity.NewGame(unitMap, time.Second.Microseconds())
	gameRoom := aggregate.NewGameRoom(game, make(map[uuid.UUID]entity.Player), make(map[uuid.UUID]valueobject.Area))
	gsi.gameRoomRealtimeRepository.Add(gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) GetAllGameRooms() []aggregate.GameRoom {
	return gsi.gameRoomRealtimeRepository.GetAll()
}

func (gsi *gameRoomDomainServiceImplement) GetGameRoom(gameId uuid.UUID) (aggregate.GameRoom, error) {
	rUnlocker, err := gsi.gameRoomRealtimeRepository.ReadLockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer rUnlocker()

	gameRoom, err := gsi.gameRoomRealtimeRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) AddPlayerToGameRoom(gameId uuid.UUID, player entity.Player) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRealtimeRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRealtimeRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.AddPlayer(player)
	gsi.gameRoomRealtimeRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRealtimeRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRealtimeRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.RemovePlayer(playerId)
	gsi.gameRoomRealtimeRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRealtimeRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRealtimeRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.AddZoomedArea(playerId, area)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRealtimeRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRealtimeRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRealtimeRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.RemoveZoomedArea(playerId)
	gsi.gameRoomRealtimeRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) TickUnitMapInGame(gameId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRealtimeRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRealtimeRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.TickUnitMap()
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRealtimeRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRealtimeRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRealtimeRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRealtimeRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}
