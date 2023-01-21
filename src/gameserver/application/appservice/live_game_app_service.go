package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
)

type LiveGameAppService interface {
	LoadGame(gameIdVm string)
	ChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm)
	BuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	DestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm)
	JoinGame(liveGameIdVm string, playerIdVm string)
	LeaveGame(liveGameIdVm string, playerIdVm string)
}

type liveGameAppServe struct {
	liveGameRepo      livegamemodel.Repo
	gameRepo          gamemodel.GameRepo
	IntEventPublisher intevent.IntEventPublisher
}

func NewLiveGameAppService(
	liveGameRepo livegamemodel.Repo,
	gameRepo gamemodel.GameRepo,
	IntEventPublisher intevent.IntEventPublisher,
) LiveGameAppService {
	return &liveGameAppServe{
		liveGameRepo:      liveGameRepo,
		gameRepo:          gameRepo,
		IntEventPublisher: IntEventPublisher,
	}
}

func (liveGameAppServe *liveGameAppServe) publishViewUpdatedEvents(liveGameId livegamemodel.LiveGameIdVo, location commonmodel.LocationVo) error {
	liveGame, err := liveGameAppServe.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	for _, playerId := range liveGame.GetPlayerIds() {
		if !liveGame.CanPlayerSeeAnyLocations(playerId, []commonmodel.LocationVo{location}) {
			continue
		}

		view, err := liveGame.GetPlayerView(playerId)
		if err != nil {
			continue
		}
		liveGameAppServe.IntEventPublisher.Publish(
			intevent.CreateLiveGameClientChannel(liveGameId.ToString(), playerId.ToString()),
			jsonmarshaller.Marshal(intevent.NewViewUpdatedIntEvent(
				liveGameId.ToString(),
				playerId.ToString(),
				viewmodel.NewViewVm(view),
			)))
	}

	return nil
}

func (liveGameAppServe *liveGameAppServe) LoadGame(gameIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}

	game, err := liveGameAppServe.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	liveGameId, _ := livegamemodel.NewLiveGameIdVo(gameId.ToString())

	liveGameMap := livegamemodel.NewMapVo(game.GetMap().GetUnitMatrix())
	newLiveGame := livegamemodel.NewLiveGameAgg(liveGameId, liveGameMap)

	liveGameAppServe.liveGameRepo.Add(newLiveGame)
}

func (liveGameAppServe *liveGameAppServe) ChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := livegamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	camera, err := cameraVm.ToValueObject()
	if err != nil {
		return
	}

	unlocker := liveGameAppServe.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := liveGameAppServe.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	if err = liveGame.ChangePlayerCamera(playerId, camera); err != nil {
		return
	}

	liveGameAppServe.liveGameRepo.Update(liveGameId, liveGame)

	player, _ := liveGame.GetPlayer(playerId)
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		jsonmarshaller.Marshal(
			intevent.NewPlayerUpdatedIntEvent(liveGameIdVm, viewmodel.NewPlayerVm(player)),
		),
	)

	view, _ := liveGame.GetPlayerView(playerId)
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		jsonmarshaller.Marshal(intevent.NewViewChangedIntEvent(liveGameIdVm, playerIdVm, viewmodel.NewViewVm(view))),
	)
}

func (liveGameAppServe *liveGameAppServe) BuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	itemId, err := itemmodel.NewItemIdVo(itemIdVm)
	if err != nil {
		return
	}
	location, err := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)
	if err != nil {
		return
	}

	unlocker := liveGameAppServe.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := liveGameAppServe.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.BuildItem(location, itemId)
	if err != nil {
		return
	}

	liveGameAppServe.liveGameRepo.Update(liveGameId, liveGame)

	liveGameAppServe.publishViewUpdatedEvents(liveGameId, location)
}

func (liveGameAppServe *liveGameAppServe) DestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	location, err := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)
	if err != nil {
		return
	}

	unlocker := liveGameAppServe.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := liveGameAppServe.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.DestroyItem(location)
	if err != nil {
		return
	}

	liveGameAppServe.liveGameRepo.Update(liveGameId, liveGame)
	liveGameAppServe.publishViewUpdatedEvents(liveGameId, location)
}

func (liveGameAppServe *liveGameAppServe) JoinGame(liveGameIdVm string, playerIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := livegamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := liveGameAppServe.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := liveGameAppServe.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.AddPlayer(playerId)
	if err != nil {
		return
	}

	liveGameAppServe.liveGameRepo.Update(liveGameId, liveGame)

	player, _ := liveGame.GetPlayer(playerId)
	camera, _ := liveGame.GetPlayerCamera(playerId)
	view, _ := liveGame.GetPlayerView(playerId)
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		jsonmarshaller.Marshal(
			intevent.NewGameJoinedIntEvent(
				liveGameIdVm,
				viewmodel.NewPlayerVm(player),
				viewmodel.NewCameraVm(camera),
				viewmodel.NewSizeVm(liveGame.GetMapSize()),
				viewmodel.NewViewVm(view),
			),
		),
	)
}

func (liveGameAppServe *liveGameAppServe) LeaveGame(liveGameIdVm string, playerIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := livegamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := liveGameAppServe.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := liveGameAppServe.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	liveGame.RemovePlayer(playerId)
	liveGameAppServe.liveGameRepo.Update(liveGameId, liveGame)
}
