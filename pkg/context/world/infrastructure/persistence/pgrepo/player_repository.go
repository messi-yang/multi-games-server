package pgrepo

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"

	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"

	"github.com/samber/lo"
)

func newPlayerModel(player playermodel.Player) pgmodel.PlayerModel {
	return pgmodel.PlayerModel{
		Id: player.GetId().Uuid(),
		UserId: lo.TernaryF(
			player.GetUserId() == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer((*player.GetUserId()).Uuid()) },
		),
		WorldId:   player.GetWorldId().Uuid(),
		Name:      player.GetName(),
		PosX:      player.GetPosition().GetX(),
		PosZ:      player.GetPosition().GetZ(),
		Direction: player.GetDirection().Int8(),
		HeldItemId: lo.TernaryF(
			player.GetHeldItemId() == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer((*player.GetHeldItemId()).Uuid()) },
		),
		CreatedAt: player.GetCreatedAt(),
		UpdatedAt: player.GetUpdatedAt(),
	}
}

func parsePlayerModel(playerModel pgmodel.PlayerModel) playermodel.Player {
	return playermodel.LoadPlayer(
		playermodel.NewPlayerId(playerModel.Id),
		sharedkernelmodel.NewWorldId(playerModel.WorldId),
		lo.TernaryF(
			playerModel.UserId == nil,
			func() *sharedkernelmodel.UserId { return nil },
			func() *sharedkernelmodel.UserId {
				return commonutil.ToPointer(sharedkernelmodel.NewUserId(*playerModel.UserId))
			},
		),
		lo.TernaryF(
			playerModel.User == nil,
			func() string { return "Untitled" },
			func() string { return playerModel.User.Username },
		),
		commonmodel.NewPosition(playerModel.PosX, playerModel.PosZ),
		commonmodel.NewDirection(playerModel.Direction),
		lo.TernaryF(
			playerModel.HeldItemId == nil,
			func() *commonmodel.ItemId { return nil },
			func() *commonmodel.ItemId {
				return commonutil.ToPointer(commonmodel.NewItemId(*playerModel.HeldItemId))
			},
		),
		playerModel.CreatedAt,
		playerModel.UpdatedAt,
	)
}

type playerRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewPlayerRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository playermodel.PlayerRepo) {
	return &playerRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *playerRepo) Add(player playermodel.Player) error {
	playerModel := newPlayerModel(player)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&playerModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Update(player playermodel.Player) error {
	playerModel := newPlayerModel(player)
	playerModel.UpdatedAt = time.Now()
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&playerModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Delete(player playermodel.Player) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.PlayerModel{Id: player.GetId().Uuid()}).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Get(_ sharedkernelmodel.WorldId, playerId playermodel.PlayerId) (player playermodel.Player, err error) {
	playerModel := pgmodel.PlayerModel{Id: playerId.Uuid()}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Joins("User").First(&playerModel).Error
	}); err != nil {
		return player, err
	}

	return parsePlayerModel(playerModel), nil
}

func (repo *playerRepo) GetPlayersOfWorld(worldId sharedkernelmodel.WorldId) (players []playermodel.Player, err error) {
	playerModels := []pgmodel.PlayerModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Joins("User").Find(
			&playerModels,
			pgmodel.PlayerModel{
				WorldId: worldId.Uuid(),
			},
		).Error
	}); err != nil {
		return players, err
	}

	return lo.Map(playerModels, func(playerModel pgmodel.PlayerModel, _ int) playermodel.Player {
		return parsePlayerModel(playerModel)
	}), nil
}
