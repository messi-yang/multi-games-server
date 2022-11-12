package postgres

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/repository"
	gameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type gamePostgres struct {
	recordLocker *sync.RWMutex
}

var gamePostgresInstance *gamePostgres

func NewGamePostgres() repository.GameRepository {
	if gamePostgresInstance == nil {
		gamePostgresInstance = &gamePostgres{}
		return gamePostgresInstance
	}
	return gamePostgresInstance
}

func (m *gamePostgres) Get(id gameValueObject.GameId) (aggregate.Game, error) {
	return aggregate.Game{}, nil
}

func (m *gamePostgres) Update(id gameValueObject.GameId, liveGame aggregate.Game) error {
	return nil
}

func (m *gamePostgres) GetAll() []aggregate.Game {
	return []aggregate.Game{}
}

func (m *gamePostgres) Add(liveGame aggregate.Game) error {
	return nil
}

func (m *gamePostgres) ReadLockAccess(gameId gameValueObject.GameId) (func(), error) {
	m.recordLocker.RLock()
	return m.recordLocker.RUnlock, nil
}

func (m *gamePostgres) LockAccess(gameId gameValueObject.GameId) (func(), error) {
	m.recordLocker.Lock()
	return m.recordLocker.Unlock, nil
}
