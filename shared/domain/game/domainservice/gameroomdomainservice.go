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
	LoadGameRoom(game entity.Game) error
	GetAllGameRooms() []aggregate.GameRoom

	AddPlayerToGameRoom(gameId uuid.UUID, player entity.Player) (aggregate.GameRoom, error)
	RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error)

	AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (aggregate.GameRoom, error)
	RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error)

	TickUnitMapInGame(gameId uuid.UUID) (aggregate.GameRoom, error)
	ReviveUnitsInGame(gameId uuid.UUID, coords []valueobject.Coordinate) (aggregate.GameRoom, error)
}

type gameRoomDomainServiceImplement struct {
	gameRoomRepository repository.GameRoomRepository
}

type GameRoomDomainServiceConfiguration struct {
	GameRoomRepository repository.GameRoomRepository
}

func NewGameRoomDomainService(coniguration GameRoomDomainServiceConfiguration) GameRoomDomainService {
	return &gameRoomDomainServiceImplement{
		gameRoomRepository: coniguration.GameRoomRepository,
	}
}

func (gsi *gameRoomDomainServiceImplement) GetAllGameRooms() []aggregate.GameRoom {
	return gsi.gameRoomRepository.GetAll()
}

func (gsi *gameRoomDomainServiceImplement) LoadGameRoom(game entity.Game) error {
	gameRoom := aggregate.NewGameRoom(game, time.Second.Microseconds())
	gsi.gameRoomRepository.Add(gameRoom)
	return nil
}

func (gsi *gameRoomDomainServiceImplement) AddPlayerToGameRoom(gameId uuid.UUID, player entity.Player) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.AddPlayer(player)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.RemovePlayer(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.AddZoomedArea(playerId, area)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gameRoom.RemoveZoomedArea(playerId)
	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) TickUnitMapInGame(gameId uuid.UUID) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.TickUnitMap()
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}

func (gsi *gameRoomDomainServiceImplement) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) (aggregate.GameRoom, error) {
	unlocker, err := gsi.gameRoomRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}
	defer unlocker()

	gameRoom, err := gsi.gameRoomRepository.Get(gameId)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	err = gameRoom.ReviveUnits(coordinates)
	if err != nil {
		return aggregate.GameRoom{}, err
	}

	gsi.gameRoomRepository.Update(gameId, gameRoom)

	return gameRoom, nil
}
