package memory

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type gameMemory struct {
	records       map[valueobject.GameId]*aggregate.Game
	recordLockers map[valueobject.GameId]*sync.RWMutex
}

var gameMemoryInstance *gameMemory

func NewGameMemory() repository.GameRepository {
	if gameMemoryInstance == nil {
		gameMemoryInstance = &gameMemory{
			records:       make(map[valueobject.GameId]*aggregate.Game),
			recordLockers: make(map[valueobject.GameId]*sync.RWMutex),
		}
		return gameMemoryInstance
	}
	return gameMemoryInstance
}

func (m *gameMemory) Get(id valueobject.GameId) (aggregate.Game, error) {
	record, exists := m.records[id]
	if !exists {
		return aggregate.Game{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *gameMemory) Update(id valueobject.GameId, game aggregate.Game) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &game

	return nil
}

func (m *gameMemory) GetAll() []aggregate.Game {
	games := make([]aggregate.Game, 0)
	for _, record := range m.records {
		games = append(games, *record)
	}
	return games
}

func (m *gameMemory) Add(game aggregate.Game) error {
	m.records[game.GetId()] = &game
	m.recordLockers[game.GetId()] = &sync.RWMutex{}

	return nil
}

func (m *gameMemory) ReadLockAccess(gameId valueobject.GameId) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *gameMemory) LockAccess(gameId valueobject.GameId) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
