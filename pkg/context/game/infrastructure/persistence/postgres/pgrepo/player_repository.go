package pgrepo

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"

	"github.com/samber/lo"
)

func newPlayerModel(player playermodel.Player) pgmodel.PlayerModel {
	return pgmodel.PlayerModel{
		Id: player.GetId().Uuid(),
		UserId: lo.Ternary(
			player.GetUserId() == nil,
			nil,
			commonutil.ToPointer((*player.GetUserId()).Uuid()),
		),
		WorldId:   player.GetWorldId().Uuid(),
		Name:      player.GetName(),
		PosX:      player.GetPosition().GetX(),
		PosZ:      player.GetPosition().GetZ(),
		Direction: player.GetDirection().Int8(),
		HeldItemId: lo.Ternary(
			player.GetHeldItemId() == nil,
			nil,
			commonutil.ToPointer((*player.GetHeldItemId()).Uuid()),
		),
		CreatedAt: player.GetCreatedAt(),
		UpdatedAt: player.GetUpdatedAt(),
	}
}

func parsePlayerModel(playerModel pgmodel.PlayerModel) playermodel.Player {
	return playermodel.LoadPlayer(
		playermodel.NewPlayerId(playerModel.Id),
		sharedkernelmodel.NewWorldId(playerModel.WorldId),
		lo.Ternary(
			playerModel.UserId == nil,
			nil,
			commonutil.ToPointer(sharedkernelmodel.NewUserId(*playerModel.UserId)),
		),
		lo.Ternary(
			playerModel.User == nil,
			"Untitled",
			playerModel.User.Username,
		),
		commonmodel.NewPosition(playerModel.PosX, playerModel.PosZ),
		commonmodel.NewDirection(playerModel.Direction),
		lo.Ternary(
			playerModel.HeldItemId == nil,
			nil,
			commonutil.ToPointer(commonmodel.NewItemId(*playerModel.HeldItemId)),
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

func (repo *playerRepo) Get(playerId playermodel.PlayerId) (player playermodel.Player, err error) {
	playerModel := pgmodel.PlayerModel{Id: playerId.Uuid()}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Joins("User").First(&playerModel).Error
	}); err != nil {
		return player, err
	}

	return parsePlayerModel(playerModel), nil
}

func (repo *playerRepo) FindPlayersAt(worldId sharedkernelmodel.WorldId, position commonmodel.Position) (players []playermodel.Player, playersFound bool, err error) {
	var playerModels = []pgmodel.PlayerModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Joins("User").Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			worldId.Uuid(),
			position.GetX(),
			position.GetZ(),
		).Find(&playerModels).Error
	}); err != nil {
		return players, playersFound, err
	}

	playersFound = len(playerModels) >= 1

	return lo.Map(playerModels, func(playerModel pgmodel.PlayerModel, _ int) playermodel.Player {
		return parsePlayerModel(playerModel)
	}), playersFound, nil
}

func (repo *playerRepo) GetPlayersAround(worldId sharedkernelmodel.WorldId, position commonmodel.Position) (players []playermodel.Player, err error) {
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

func (repo *playerRepo) GetAll(worldId sharedkernelmodel.WorldId) []playermodel.Player {
	var playerModels []pgmodel.PlayerModel
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Joins("User").Find(&playerModels).Error
	}); err != nil {
		return []playermodel.Player{}
	}

	return lo.Map(playerModels, func(playerModel pgmodel.PlayerModel, _ int) playermodel.Player {
		return parsePlayerModel(playerModel)
	})
}
