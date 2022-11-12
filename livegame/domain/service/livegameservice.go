package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/infrastructure/memory"
	"github.com/google/uuid"
)

type LiveGameService struct {
	liveGameRepository repository.LiveGameRepository
	hardcodedGameIdId  valueobject.GameId
}

type liveGameServiceConfiguration func(service *LiveGameService) error

func NewLiveGameService(cfgs ...liveGameServiceConfiguration) (*LiveGameService, error) {
	hardcodedGameIdId, _ := uuid.Parse("1a53a474-ebbd-49e4-a2c1-dde5aa5759bc")
	t := &LiveGameService{
		hardcodedGameIdId: valueobject.NewGameId(hardcodedGameIdId),
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

func (gs *LiveGameService) GetAllLiveGameIds() []valueobject.GameId {
	return []valueobject.GameId{gs.hardcodedGameIdId}
}

func (gs *LiveGameService) CreateLiveGame(dimension valueobject.Dimension) (valueobject.GameId, error) {
	unitBlock := make([][]valueobject.Unit, dimension.GetWidth())
	for i := 0; i < dimension.GetWidth(); i += 1 {
		unitBlock[i] = make([]valueobject.Unit, dimension.GetHeight())
		for j := 0; j < dimension.GetHeight(); j += 1 {
			unitBlock[i][j] = valueobject.NewUnit(false, valueobject.ItemTypeEmpty)
		}
	}
	newLiveGame := aggregate.NewLiveGame(gs.hardcodedGameIdId, valueobject.NewUnitBlock(unitBlock))
	gs.liveGameRepository.Add(newLiveGame)
	return newLiveGame.GetId(), nil
}

func (service *LiveGameService) GetLiveGame(id valueobject.GameId) (aggregate.LiveGame, error) {
	liveGame, err := service.liveGameRepository.Get(id)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	return liveGame, nil
}

func (gs *LiveGameService) AddPlayerToLiveGame(gameId valueobject.GameId, playerId valueobject.PlayerId) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gameLive.AddPlayer(playerId)
	gs.liveGameRepository.Update(gameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) RemovePlayerFromLiveGame(gameId valueobject.GameId, playerId valueobject.PlayerId) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gameLive.RemovePlayer(playerId)
	gs.liveGameRepository.Update(gameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) AddZoomedAreaToLiveGame(gameId valueobject.GameId, playerId valueobject.PlayerId, area valueobject.Area) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	err = gameLive.AddZoomedArea(playerId, area)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gs.liveGameRepository.Update(gameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) RemoveZoomedAreaFromLiveGame(gameId valueobject.GameId, playerId valueobject.PlayerId) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gameLive.RemoveZoomedArea(playerId)
	gs.liveGameRepository.Update(gameId, gameLive)

	return gameLive, nil
}

func (gs *LiveGameService) ReviveUnitsInLiveGame(gameId valueobject.GameId, coordinates []valueobject.Coordinate) (aggregate.LiveGame, error) {
	unlocker, err := gs.liveGameRepository.LockAccess(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := gs.liveGameRepository.Get(gameId)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	err = gameLive.ReviveUnits(coordinates)
	if err != nil {
		return aggregate.LiveGame{}, err
	}

	gs.liveGameRepository.Update(gameId, gameLive)

	return gameLive, nil
}
