package pgrepo

import (
	"fmt"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
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

func (repo *gameRepo) Get(gameId gamemodel.GameId) (game gamemodel.Game, err error) {
	gameModel := pgmodel.GameModel{Id: gameId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&gameModel).Error
	}); err != nil {
		return game, err
	}
	return pgmodel.ParseGameModel(gameModel), nil
}

func (repo *gameRepo) Add(game gamemodel.Game) error {
	gameModel := pgmodel.NewGameModel(game)
	fmt.Println("Game model", gameModel)
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
		return transaction.Model(&pgmodel.GameModel{}).Where(
			"id = ?",
			game.GetId().Uuid(),
		).Select("*").Updates(&gameModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&game)
}
