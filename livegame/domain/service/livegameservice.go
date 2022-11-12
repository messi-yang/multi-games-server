package service

import (
	commonValueObject "github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/repository"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/infrastructure/memory"
	"github.com/google/uuid"
)

type LiveGameService struct {
	liveGameRepository  repository.LiveGameRepository
	hardcodedLiveGameId liveGameValueObject.LiveGameId
}

type liveGameServiceConfiguration func(service *LiveGameService) error

func NewLiveGameService(cfgs ...liveGameServiceConfiguration) (*LiveGameService, error) {
	hardcodedLiveGameId, _ := uuid.Parse("1a53a474-ebbd-49e4-a2c1-dde5aa5759bc")
	t := &LiveGameService{
		hardcodedLiveGameId: liveGameValueObject.NewLiveGameId(hardcodedLiveGameId),
	}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithGameMemory() liveGameServiceConfiguration {
	liveGameMemory := memory.NewLiveGameMemory()
	return func(service *LiveGameService) error {
		service.liveGameRepository = liveGameMemory
		return nil
	}
}

func (gs *LiveGameService) GetAllLiveGameIds() []liveGameValueObject.LiveGameId {
	return []liveGameValueObject.LiveGameId{gs.hardcodedLiveGameId}
}

func (gs *LiveGameService) CreateLiveGame(dimension commonValueObject.Dimension) (liveGameValueObject.LiveGameId, error) {
	unitBlock := make([][]commonValueObject.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitBlock[i] = make([]commonValueObject.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitBlock[i][j] = commonValueObject.NewUnit(false, commonValueObject.ItemTypeEmpty)
		}
	}
	newLiveGame := aggregate.NewLiveGame(gs.hardcodedLiveGameId, commonValueObject.NewUnitBlock(unitBlock))
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
