package service

import (
	commonValueObject "github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	gameAggregate "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/repository"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/infrastructure/memoryrepository"
)

type LiveGameService struct {
	liveGameRepository repository.LiveGameRepository
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
	liveGameMemoryRepository := memoryrepository.NewLiveGameMemoryRepository()
	return func(service *LiveGameService) error {
		service.liveGameRepository = liveGameMemoryRepository
		return nil
	}
}

func (gs *LiveGameService) CreateLiveGame(game gameAggregate.Game) (liveGameValueObject.LiveGameId, error) {
	newLiveGame := aggregate.NewLiveGame(liveGameValueObject.NewLiveGameId(game.GetId().GetId()), game.GetUnitBlock())
	gs.liveGameRepository.Add(newLiveGame)
	return newLiveGame.GetId(), nil
}

func (service *LiveGameService) GetLiveGame(id liveGameValueObject.LiveGameId) (aggregate.LiveGame, error) {
	liveGame, err := service.liveGameRepository.Get(id)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	return liveGame, nil
}

func (gs *LiveGameService) AddPlayerToLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId commonValueObject.PlayerId) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gameLive.AddPlayer(playerId)
	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) RemovePlayerFromLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId commonValueObject.PlayerId) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gameLive.RemovePlayer(playerId)
	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) AddZoomedAreaToLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId commonValueObject.PlayerId, area commonValueObject.Area) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	err = gameLive.AddZoomedArea(playerId, area)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) RemoveZoomedAreaFromLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId commonValueObject.PlayerId) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gameLive.RemoveZoomedArea(playerId)
	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) ReviveUnitsInLiveGame(liveGameId liveGameValueObject.LiveGameId, coordinates []commonValueObject.Coordinate) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(liveGameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	err = gameLive.ReviveUnits(coordinates)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gs.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}
