package memrepo

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type gameMemRepo struct {
	records       map[gamemodel.GameIdVo]*gamemodel.GameAgg
	recordLockers map[gamemodel.GameIdVo]*sync.RWMutex
}

var gameMemRepoSingleton *gameMemRepo

func NewGameMemRepo() gamemodel.Repo {
	if gameMemRepoSingleton == nil {
		gameMemRepoSingleton = &gameMemRepo{
			records:       make(map[gamemodel.GameIdVo]*gamemodel.GameAgg),
			recordLockers: make(map[gamemodel.GameIdVo]*sync.RWMutex),
		}
		return gameMemRepoSingleton
	}
	return gameMemRepoSingleton
}

func (m *gameMemRepo) Get(id gamemodel.GameIdVo) (gamemodel.GameAgg, error) {
	record, exists := m.records[id]
	if !exists {
		return gamemodel.GameAgg{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *gameMemRepo) Update(id gamemodel.GameIdVo, game gamemodel.GameAgg) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &game

	return nil
}

func (m *gameMemRepo) GetAll() []gamemodel.GameAgg {
	games := make([]gamemodel.GameAgg, 0)
	for _, record := range m.records {
		games = append(games, *record)
	}
	return games
}

func (m *gameMemRepo) Add(game gamemodel.GameAgg) error {
	m.records[game.GetId()] = &game

	return nil
}

func (m *gameMemRepo) ReadLockAccess(gameId gamemodel.GameIdVo) (unlock func()) {
	_, exists := m.recordLockers[gameId]
	if !exists {
		m.recordLockers[gameId] = &sync.RWMutex{}
	}

	m.recordLockers[gameId].RLock()
	return m.recordLockers[gameId].RUnlock
}

func (m *gameMemRepo) LockAccess(gameId gamemodel.GameIdVo) (unlock func()) {
	_, exists := m.recordLockers[gameId]
	if !exists {
		m.recordLockers[gameId] = &sync.RWMutex{}
	}

	m.recordLockers[gameId].Lock()
	return m.recordLockers[gameId].Unlock
}
