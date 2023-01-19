package memrepo

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

type liveGameMemRepo struct {
	records       map[livegamemodel.LiveGameIdVo]*livegamemodel.LiveGameAgg
	recordLockers map[livegamemodel.LiveGameIdVo]*sync.RWMutex
}

var liveGameMemRepoSingleton *liveGameMemRepo

func NewLiveGameMemRepo() livegamemodel.Repo {
	if liveGameMemRepoSingleton == nil {
		liveGameMemRepoSingleton = &liveGameMemRepo{
			records:       make(map[livegamemodel.LiveGameIdVo]*livegamemodel.LiveGameAgg),
			recordLockers: make(map[livegamemodel.LiveGameIdVo]*sync.RWMutex),
		}
		return liveGameMemRepoSingleton
	}
	return liveGameMemRepoSingleton
}

func (m *liveGameMemRepo) Get(id livegamemodel.LiveGameIdVo) (livegamemodel.LiveGameAgg, error) {
	record, exists := m.records[id]
	if !exists {
		return livegamemodel.LiveGameAgg{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *liveGameMemRepo) Update(id livegamemodel.LiveGameIdVo, liveGame livegamemodel.LiveGameAgg) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &liveGame

	return nil
}

func (m *liveGameMemRepo) GetAll() []livegamemodel.LiveGameAgg {
	liveGames := make([]livegamemodel.LiveGameAgg, 0)
	for _, record := range m.records {
		liveGames = append(liveGames, *record)
	}
	return liveGames
}

func (m *liveGameMemRepo) Add(liveGame livegamemodel.LiveGameAgg) error {
	m.records[liveGame.GetId()] = &liveGame

	return nil
}

func (m *liveGameMemRepo) ReadLockAccess(liveGameId livegamemodel.LiveGameIdVo) (unlock func()) {
	_, exists := m.recordLockers[liveGameId]
	if !exists {
		m.recordLockers[liveGameId] = &sync.RWMutex{}
	}

	m.recordLockers[liveGameId].RLock()
	return m.recordLockers[liveGameId].RUnlock
}

func (m *liveGameMemRepo) LockAccess(liveGameId livegamemodel.LiveGameIdVo) (unlock func()) {
	_, exists := m.recordLockers[liveGameId]
	if !exists {
		m.recordLockers[liveGameId] = &sync.RWMutex{}
	}

	m.recordLockers[liveGameId].Lock()
	return m.recordLockers[liveGameId].Unlock
}
