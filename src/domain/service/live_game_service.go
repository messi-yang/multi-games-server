package service

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type LiveGameService interface {
	MovePlayer(liveGameId livegamemodel.LiveGameIdVo, playerId livegamemodel.PlayerIdVo, direction livegamemodel.DirectionVo) error
}

type liveGameServe struct {
	liveGameRepo livegamemodel.Repo
	itemRepo     itemmodel.Repo
}

func NewLiveGameService(liveGameRepo livegamemodel.Repo, itemRepo itemmodel.Repo) LiveGameService {
	return &liveGameServe{liveGameRepo: liveGameRepo, itemRepo: itemRepo}
}

func (serve *liveGameServe) MovePlayer(liveGameId livegamemodel.LiveGameIdVo, playerId livegamemodel.PlayerIdVo, direction livegamemodel.DirectionVo) error {
	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	player, err := liveGame.GetPlayer(playerId)
	if err != nil {
		return err
	}

	newLocation := player.GetLocation()
	if direction.IsUp() {
		newLocation = newLocation.Shift(0, -1)
	} else if direction.IsRight() {
		newLocation = newLocation.Shift(1, 0)
	} else if direction.IsDown() {
		newLocation = newLocation.Shift(0, 1)
	} else if direction.IsLeft() {
		newLocation = newLocation.Shift(-1, 0)
	}

	if !liveGame.GetMapSize().CoversLocation(newLocation) {
		return errors.New("player location is out of map")
	}

	itemId := liveGame.GetUnit(newLocation).GetItemId()
	if !itemId.IsEmpty() {
		item, _ := serve.itemRepo.Get(itemId)
		if !item.IsTraversable() {
			return errors.New("this item is not traversable")
		}
	}

	player.SetLocation(newLocation)
	liveGame.UpdatePlayer(player)
	serve.liveGameRepo.Update(liveGameId, liveGame)

	return nil
}
