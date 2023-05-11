package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/google/uuid"

	"github.com/samber/lo"
)

func newPlayerModel(player playermodel.Player) pgmodel.PlayerModel {
	return pgmodel.PlayerModel{
		Id:        player.GetId().Uuid(),
		GamerId:   nil,
		WorldId:   player.GetWorldId().Uuid(),
		Name:      player.GetName(),
		PosX:      player.GetPosition().GetX(),
		PosZ:      player.GetPosition().GetZ(),
		Direction: player.GetDirection().Int8(),
		HeldItemId: lo.TernaryF(
			player.GetHeldItemId() == nil,
			func() *uuid.UUID {
				return nil
			},
			func() *uuid.UUID {
				heldItemUuid := (*player.GetHeldItemId()).Uuid()
				return &heldItemUuid
			},
		),
		CreatedAt: player.GetCreatedAt(),
		UpdatedAt: player.GetUpdatedAt(),
	}
}

func parsePlayerModel(playerModel pgmodel.PlayerModel) playermodel.Player {
	return playermodel.LoadPlayer(
		commonmodel.NewPlayerId(playerModel.Id),
		commonmodel.NewWorldId(playerModel.WorldId),
		playerModel.Name,
		commonmodel.NewPosition(playerModel.PosX, playerModel.PosZ),
		commonmodel.NewDirection(playerModel.Direction),
		lo.TernaryF(
			playerModel.HeldItemId == nil,
			func() *commonmodel.ItemId { return nil },
			func() *commonmodel.ItemId {
				heldItemId := commonmodel.NewItemId(*playerModel.HeldItemId)
				return &heldItemId
			},
		),
		playerModel.CreatedAt,
		playerModel.UpdatedAt,
	)
}

type playerRepo struct {
	uow pguow.Uow
}

func NewPlayerRepo(uow pguow.Uow) (repository playermodel.Repo) {
	return &playerRepo{uow: uow}
}

func (repo *playerRepo) Add(player playermodel.Player) error {
	playerModel := newPlayerModel(player)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&playerModel).Error
	})
}

func (repo *playerRepo) Update(player playermodel.Player) error {
	playerModel := newPlayerModel(player)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&playerModel).Error
	})
}

func (repo *playerRepo) Delete(player playermodel.Player) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.PlayerModel{Id: player.GetId().Uuid()}).Error
	})
}

func (repo *playerRepo) Get(playerId commonmodel.PlayerId) (player playermodel.Player, err error) {
	playerModel := pgmodel.PlayerModel{Id: playerId.Uuid()}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&playerModel).Error
	}); err != nil {
		return player, err
	}

	return parsePlayerModel(playerModel), nil
}

func (repo *playerRepo) FindPlayersAt(worldId commonmodel.WorldId, position commonmodel.Position) (players []playermodel.Player, playersFound bool, err error) {
	var playerModels = []pgmodel.PlayerModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
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

func (repo *playerRepo) GetPlayersAround(worldId commonmodel.WorldId, position commonmodel.Position) (players []playermodel.Player, err error) {
	playerModels := []pgmodel.PlayerModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(
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

func (repo *playerRepo) GetAll(worldId commonmodel.WorldId) []playermodel.Player {
	var playerModels []pgmodel.PlayerModel
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&playerModels).Error
	}); err != nil {
		return []playermodel.Player{}
	}

	return lo.Map(playerModels, func(playerModel pgmodel.PlayerModel, _ int) playermodel.Player {
		return parsePlayerModel(playerModel)
	})
}
