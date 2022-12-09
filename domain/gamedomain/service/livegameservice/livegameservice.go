package livegameservice

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
)

type LiveGameService interface {
	CreateLiveGame(game gamemodel.Game) (livegamemodel.LiveGameId, error)
	AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error)
	RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error)
	AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) (livegamemodel.LiveGame, error)
	RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error)
	ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) (livegamemodel.LiveGame, error)
}

type LiveGameServe struct {
	liveGameRepository livegamemodel.LiveGameRepository
}

func NewLiveGameService(liveGameRepository livegamemodel.LiveGameRepository) *LiveGameServe {
	return &LiveGameServe{
		liveGameRepository: liveGameRepository,
	}
}

func (serve *LiveGameServe) CreateLiveGame(game gamemodel.Game) (livegamemodel.LiveGameId, error) {
	newLiveGame := livegamemodel.NewLiveGame(livegamemodel.NewLiveGameId(game.GetId().GetId()), game.GetUnitBlock())
	serve.liveGameRepository.Add(newLiveGame)
	return newLiveGame.GetId(), nil
}

func (serce *LiveGameServe) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error) {
	unlocker, err := serce.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := serce.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gameLive.AddPlayer(playerId)
	serce.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (serce *LiveGameServe) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error) {
	unlocker, err := serce.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := serce.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gameLive.RemovePlayer(playerId)
	serce.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (serce *LiveGameServe) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) (livegamemodel.LiveGame, error) {
	unlocker, err := serce.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := serce.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	err = gameLive.AddZoomedArea(playerId, area)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	serce.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (serce *LiveGameServe) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (livegamemodel.LiveGame, error) {
	unlocker, err := serce.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := serce.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	gameLive.RemoveZoomedArea(playerId)
	serce.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}

func (serce *LiveGameServe) ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) (livegamemodel.LiveGame, error) {
	unlocker, err := serce.liveGameRepository.LockAccess(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}
	defer unlocker()

	gameLive, err := serce.liveGameRepository.Get(liveGameId)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	err = gameLive.ReviveUnits(coordinates)
	if err != nil {
		return livegamemodel.LiveGame{}, err
	}

	serce.liveGameRepository.Update(liveGameId, gameLive)

	return gameLive, nil
}
