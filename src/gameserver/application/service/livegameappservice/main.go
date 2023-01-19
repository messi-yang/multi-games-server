package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
)

type Service interface {
	CreateLiveGame(gameIdVm string)
	ChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm)
	BuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	DestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm)
	AddPlayer(liveGameIdVm string, playerIdVm string)
	RemovePlayer(liveGameIdVm string, playerIdVm string)
}

type serve struct {
	liveGameRepo        livegamemodel.Repo
	gameRepo            gamemodel.GameRepo
	intgrEventPublisher intgrevent.IntgrEventPublisher
}

func New(
	liveGameRepo livegamemodel.Repo,
	gameRepo gamemodel.GameRepo,
	intgrEventPublisher intgrevent.IntgrEventPublisher,
) *serve {
	return &serve{
		liveGameRepo:        liveGameRepo,
		gameRepo:            gameRepo,
		intgrEventPublisher: intgrEventPublisher,
	}
}

func (serve *serve) publishViewUpdatedEvents(liveGameId livegamemodel.LiveGameIdVo, location commonmodel.LocationVo) error {
	liveGame, err := serve.liveGameRepo.Get(liveGameId)
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
		serve.intgrEventPublisher.Publish(
			intgrevent.CreateLiveGameClientChannel(liveGameId.ToString(), playerId.ToString()),
			jsonmarshaller.Marshal(intgrevent.NewViewUpdatedIntgrEvent(
				liveGameId.ToString(),
				playerId.ToString(),
				viewmodel.NewViewVm(view),
			)))
	}

	return nil
}

func (serve *serve) CreateLiveGame(gameIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	liveGameId, _ := livegamemodel.NewLiveGameIdVo(gameId.ToString())

	liveGameMap := livegamemodel.NewMapVo(game.GetMap().GetUnitMatrix())
	newLiveGame := livegamemodel.NewLiveGameAgg(liveGameId, liveGameMap)

	serve.liveGameRepo.Add(newLiveGame)
}

func (serve *serve) ChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	camera, err := cameraVm.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	if err = liveGame.ChangePlayerCamera(playerId, camera); err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)

	view, _ := liveGame.GetPlayerView(playerId)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		jsonmarshaller.Marshal(
			intgrevent.NewCameraChangedIntgrEvent(liveGameIdVm, playerIdVm, viewmodel.NewCameraVm(camera)),
		),
	)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		jsonmarshaller.Marshal(intgrevent.NewViewChangedIntgrEvent(liveGameIdVm, playerIdVm, viewmodel.NewViewVm(view))),
	)
}

func (serve *serve) BuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
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

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.BuildItem(location, itemId)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)

	serve.publishViewUpdatedEvents(liveGameId, location)
}

func (serve *serve) DestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	location, err := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.DestroyItem(location)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.publishViewUpdatedEvents(liveGameId, location)
}

func (serve *serve) AddPlayer(liveGameIdVm string, playerIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	liveGame.AddPlayer(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)

	camera, _ := liveGame.GetPlayerCamera(playerId)
	view, _ := liveGame.GetPlayerView(playerId)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		jsonmarshaller.Marshal(
			intgrevent.NewGameJoinedIntgrEvent(
				liveGameIdVm, playerIdVm,
				viewmodel.NewCameraVm(camera),
				viewmodel.NewSizeVm(liveGame.GetMapSize()),
				viewmodel.NewViewVm(view),
			),
		),
	)
}

func (serve *serve) RemovePlayer(liveGameIdVm string, playerIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameIdVo(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	liveGame.RemovePlayer(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)
}
