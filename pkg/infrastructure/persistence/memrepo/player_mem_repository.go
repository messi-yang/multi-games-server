package memrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type playerMemRepository struct {
	players []playermodel.PlayerAgg
}

var playerMemRepositorySingleton *playerMemRepository

func NewPlayerMemRepository() playermodel.Repository {
	if playerMemRepositorySingleton == nil {
		playerMemRepositorySingleton = &playerMemRepository{
			players: make([]playermodel.PlayerAgg, 0),
		}
		return playerMemRepositorySingleton
	}
	return playerMemRepositorySingleton
}

func (repo *playerMemRepository) Add(newPlayer playermodel.PlayerAgg) error {
	newPlayers := append(repo.players, newPlayer)
	repo.players = lo.UniqBy(newPlayers, func(player playermodel.PlayerAgg) uuid.UUID {
		return player.GetId().Uuid()
	})
	return nil
}

func (repo *playerMemRepository) Get(playerId playermodel.PlayerIdVo) (playermodel.PlayerAgg, error) {
	foundPlayer, found := lo.Find(repo.players, func(player playermodel.PlayerAgg) bool {
		return player.GetId().IsEqual(playerId)
	})
	if !found {
		return playermodel.PlayerAgg{}, playermodel.ErrPlayerNotFound
	}
	return foundPlayer, nil
}

func (repo *playerMemRepository) GetPlayerAt(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) (playermodel.PlayerAgg, bool, error) {
	foundPlayer, found := lo.Find(repo.players, func(player playermodel.PlayerAgg) bool {
		return player.GetWorldId().IsEqual(worldId) && player.GetPosition().IsEqual(position)
	})
	if !found {
		return playermodel.PlayerAgg{}, false, nil
	}
	return foundPlayer, true, nil
}

func (repo *playerMemRepository) GetPlayersAround(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) ([]playermodel.PlayerAgg, error) {
	return lo.Filter(repo.players, func(player playermodel.PlayerAgg, _ int) bool {
		return player.GetWorldId().IsEqual(worldId) && player.CanSeeAnyPositions([]commonmodel.PositionVo{position})
	}), nil
}

func (repo *playerMemRepository) Update(updatedPlayer playermodel.PlayerAgg) error {
	repo.players = lo.Map(repo.players, func(player playermodel.PlayerAgg, _ int) playermodel.PlayerAgg {
		if player.GetId().IsEqual(updatedPlayer.GetId()) {
			return updatedPlayer
		}
		return player
	})
	return nil
}

func (repo *playerMemRepository) GetAll(worldId worldmodel.WorldIdVo) []playermodel.PlayerAgg {
	return lo.Filter(repo.players, func(player playermodel.PlayerAgg, _ int) bool {
		return player.GetWorldId().IsEqual(worldId)
	})
}

func (repo *playerMemRepository) Delete(playerId playermodel.PlayerIdVo) error {
	repo.players = lo.Filter(repo.players, func(player playermodel.PlayerAgg, _ int) bool {
		return !player.GetId().IsEqual(playerId)
	})
	return nil
}
