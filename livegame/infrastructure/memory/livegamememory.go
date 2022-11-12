package memory

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/repository"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type gameMemory struct {
	records       map[liveGameValueObject.LiveGameId]*aggregate.LiveGame
	recordLockers map[liveGameValueObject.LiveGameId]*sync.RWMutex
}

var gameMemoryInstance *gameMemory

func NewLiveGameMemory() repository.LiveGameRepository {
	if gameMemoryInstance == nil {
		gameMemoryInstance = &gameMemory{
			records:       make(map[liveGameValueObject.LiveGameId]*aggregate.LiveGame),
			recordLockers: make(map[liveGameValueObject.LiveGameId]*sync.RWMutex),
		}
		return gameMemoryInstance
	}
	return gameMemoryInstance
}

func (m *gameMemory) Get(id liveGameValueObject.LiveGameId) (aggregate.LiveGame, error) {
	record, exists := m.records[id]
	if !exists {
		return aggregate.LiveGame{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *gameMemory) Update(id liveGameValueObject.LiveGameId, liveGame aggregate.LiveGame) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &liveGame

	return nil
}

func (m *gameMemory) GetAll() []aggregate.LiveGame {
	liveGames := make([]aggregate.LiveGame, 0)
	for _, record := range m.records {
		liveGames = append(liveGames, *record)
	}
	return liveGames
}

func (m *gameMemory) Add(liveGame aggregate.LiveGame) error {
	m.records[liveGame.GetId()] = &liveGame
	m.recordLockers[liveGame.GetId()] = &sync.RWMutex{}

	return nil
}

func (m *gameMemory) ReadLockAccess(liveGameId liveGameValueObject.LiveGameId) (func(), error) {
	recordLocker, exists := m.recordLockers[liveGameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *gameMemory) LockAccess(liveGameId liveGameValueObject.LiveGameId) (func(), error) {
	recordLocker, exists := m.recordLockers[liveGameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
