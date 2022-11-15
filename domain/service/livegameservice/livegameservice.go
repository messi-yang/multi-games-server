package livegameservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/memoryrepository"
)

type LiveGameService struct {
	liveGameRepository livegamemodel.LiveGameRepository
}

type liveGameServiceConfiguration func(service *LiveGameService) error

func NewLiveGameService(cfgs ...liveGameServiceConfiguration) (*LiveGameService, error) {
	t := &LiveGameService{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithGameMemoryRepository() liveGameServiceConfiguration {
	liveGameMemoryRepository := memoryrepository.NewMemoryLiveGameRepository()
	return func(service *LiveGameService) error {
		service.liveGameRepository = liveGameMemoryRepository
		return nil
	}
}

func (gs *LiveGameService) CreateLiveGame(game gamemodel.Game) (livegamemodel.LiveGameId, error) {
	newLiveGame := livegamemodel.NewLiveGame(livegamemodel.NewLiveGameId(game.GetId().GetId()), game.GetUnitBlock())
	gs.liveGameRepository.Add(newLiveGame)
	return newLiveGame.GetId(), nil
}

func (service *LiveGameService) GetLiveGame(id livegamemodel.LiveGameId) (livegamemodel.LiveGame, error) {
	liveGame, err := service.liveGameRepository.Get(id)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	return liveGame, nil
}

func (gs *LiveGameService) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gameLive.AddPlayer(playerId)
	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gameLive.RemovePlayer(playerId)
	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) (livegamemodel.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	err = gameLive.AddZoomedArea(playerId, area)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gameLive.RemoveZoomedArea(playerId)
	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) (livegamemodel.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	err = gameLive.ReviveUnits(coordinates)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}
