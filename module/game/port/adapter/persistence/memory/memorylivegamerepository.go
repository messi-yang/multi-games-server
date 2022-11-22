package memory

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type memoryLiveGameRepository struct {
	records       map[livegamemodel.LiveGameId]*livegamemodel.LiveGame
	recordLockers map[livegamemodel.LiveGameId]*sync.RWMutex
}

var memoryLiveGameRepositoryInstance *memoryLiveGameRepository

func NewMemoryLiveGameRepository() livegamemodel.LiveGameRepository {
	if memoryLiveGameRepositoryInstance == nil {
		memoryLiveGameRepositoryInstance = &memoryLiveGameRepository{
			records:       make(map[livegamemodel.LiveGameId]*livegamemodel.LiveGame),
			recordLockers: make(map[livegamemodel.LiveGameId]*sync.RWMutex),
		}
		return memoryLiveGameRepositoryInstance
	}
	return memoryLiveGameRepositoryInstance
}

func (m *memoryLiveGameRepository) Get(id livegamemodel.LiveGameId) (livegamemodel.LiveGame, error) {
	record, exists := m.records[id]
	if !exists {
		return livegamemodel.LiveGame{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *memoryLiveGameRepository) Update(id livegamemodel.LiveGameId, liveGame livegamemodel.LiveGame) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &liveGame

	return nil
}

func (m *memoryLiveGameRepository) GetAll() []livegamemodel.LiveGame {
	liveGames := make([]livegamemodel.LiveGame, 0)
	for _, record := range m.records {
		liveGames = append(liveGames, *record)
	}
	return liveGames
}

func (m *memoryLiveGameRepository) Add(liveGame livegamemodel.LiveGame) error {
	m.records[liveGame.GetId()] = &liveGame
	m.recordLockers[liveGame.GetId()] = &sync.RWMutex{}

	return nil
}

func (m *memoryLiveGameRepository) ReadLockAccess(liveGameId livegamemodel.LiveGameId) (func(), error) {
	recordLocker, exists := m.recordLockers[liveGameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *memoryLiveGameRepository) LockAccess(liveGameId livegamemodel.LiveGameId) (func(), error) {
	recordLocker, exists := m.recordLockers[liveGameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
