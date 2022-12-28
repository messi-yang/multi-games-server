package livegamememoryrepository

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type memoryRepository struct {
	records       map[livegamemodel.LiveGameId]*livegamemodel.LiveGame
	recordLockers map[livegamemodel.LiveGameId]*sync.RWMutex
}

var singleton *memoryRepository

func New() livegamemodel.Repository {
	if singleton == nil {
		singleton = &memoryRepository{
			records:       make(map[livegamemodel.LiveGameId]*livegamemodel.LiveGame),
			recordLockers: make(map[livegamemodel.LiveGameId]*sync.RWMutex),
		}
		return singleton
	}
	return singleton
}

func (m *memoryRepository) Get(id livegamemodel.LiveGameId) (livegamemodel.LiveGame, error) {
	record, exists := m.records[id]
	if !exists {
		return livegamemodel.LiveGame{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *memoryRepository) Update(id livegamemodel.LiveGameId, liveGame livegamemodel.LiveGame) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &liveGame

	return nil
}

func (m *memoryRepository) GetAll() []livegamemodel.LiveGame {
	liveGames := make([]livegamemodel.LiveGame, 0)
	for _, record := range m.records {
		liveGames = append(liveGames, *record)
	}
	return liveGames
}

func (m *memoryRepository) Add(liveGame livegamemodel.LiveGame) error {
	m.records[liveGame.GetId()] = &liveGame

	return nil
}

func (m *memoryRepository) ReadLockAccess(liveGameId livegamemodel.LiveGameId) (unlock func()) {
	_, exists := m.recordLockers[liveGameId]
	if !exists {
		m.recordLockers[liveGameId] = &sync.RWMutex{}
	}

	m.recordLockers[liveGameId].RLock()
	return m.recordLockers[liveGameId].RUnlock
}

func (m *memoryRepository) LockAccess(liveGameId livegamemodel.LiveGameId) (unlock func()) {
	_, exists := m.recordLockers[liveGameId]
	if !exists {
		m.recordLockers[liveGameId] = &sync.RWMutex{}
	}

	m.recordLockers[liveGameId].Lock()
	return m.recordLockers[liveGameId].Unlock
}
