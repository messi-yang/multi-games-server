package pgrepo

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type gameRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewGameRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository gamemodel.GameRepo) {
	return &gameRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *gameRepo) Add(game gamemodel.Game) error {
	gameModel := pgmodel.NewGameModel(game)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&gameModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&game)
}

func (repo *gameRepo) Update(game gamemodel.Game) error {
	gameModel := pgmodel.NewGameModel(game)
	gameModel.UpdatedAt = time.Now()
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.RoomModel{}).Where(
			"id = ?",
			game.GetId().Uuid(),
		).Select("*").Updates(&gameModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&game)
}

func (repo *gameRepo) GetSelectedGameByRoomId(roomId globalcommonmodel.RoomId) (game gamemodel.Game, err error) {
	gameModel := pgmodel.GameModel{RoomId: roomId.Uuid(), Selected: true}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&gameModel).Error
	}); err != nil {
		return game, err
	}
	return pgmodel.ParseGameModel(gameModel), nil
}

func (repo *gameRepo) GetSelectedGamesByRoomId(roomId globalcommonmodel.RoomId) (games []gamemodel.Game, err error) {
	gameModels := []pgmodel.GameModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"room_id = ?",
			roomId.Uuid(),
		).Where(
			"selected = ?",
			true,
		).Find(&gameModels).Error
	}); err != nil {
		return games, err
	}
	return commonutil.MapWithError[pgmodel.GameModel](gameModels, func(_ int, gameModel pgmodel.GameModel) (gamemodel.Game, error) {
		return pgmodel.ParseGameModel(gameModel), nil
	})
}
