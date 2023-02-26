package memrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/samber/lo"
)

type playerMemRepo struct {
	players []playermodel.PlayerAgg
}

var playerMemRepoSingleton *playerMemRepo

func NewPlayerMemRepo() playermodel.Repo {
	if playerMemRepoSingleton == nil {
		playerMemRepoSingleton = &playerMemRepo{
			players: make([]playermodel.PlayerAgg, 0),
		}
		return playerMemRepoSingleton
	}
	return playerMemRepoSingleton
}

func (repo *playerMemRepo) Add(newPlayer playermodel.PlayerAgg) error {
	newPlayers := append(repo.players, newPlayer)
	repo.players = lo.UniqBy(newPlayers, func(player playermodel.PlayerAgg) string {
		return player.GetId().ToString()
	})
	return nil
}

func (repo *playerMemRepo) Get(playerId playermodel.PlayerIdVo) (playermodel.PlayerAgg, error) {
	foundPlayer, found := lo.Find(repo.players, func(player playermodel.PlayerAgg) bool {
		return player.GetId().IsEqual(playerId)
	})
	if !found {
		return playermodel.PlayerAgg{}, playermodel.ErrPlayerNotFound
	}
	return foundPlayer, nil
}

func (repo *playerMemRepo) GetPlayerAt(gameId gamemodel.GameIdVo, position commonmodel.PositionVo) (playermodel.PlayerAgg, bool, error) {
	foundPlayer, found := lo.Find(repo.players, func(player playermodel.PlayerAgg) bool {
		return player.GetGameId().IsEqual(gameId) && player.GetPosition().IsEqual(position)
	})
	if !found {
		return playermodel.PlayerAgg{}, false, nil
	}
	return foundPlayer, true, nil
}

func (repo *playerMemRepo) GetPlayersAround(gameId gamemodel.GameIdVo, position commonmodel.PositionVo) ([]playermodel.PlayerAgg, error) {
	return lo.Filter(repo.players, func(player playermodel.PlayerAgg, _ int) bool {
		return player.GetGameId().IsEqual(gameId) && player.CanSeeAnyPositions([]commonmodel.PositionVo{position})
	}), nil
	return nil, nil
}

func (repo *playerMemRepo) Update(updatedPlayer playermodel.PlayerAgg) error {
	repo.players = lo.Map(repo.players, func(player playermodel.PlayerAgg, _ int) playermodel.PlayerAgg {
		if player.GetId().IsEqual(updatedPlayer.GetId()) {
			return updatedPlayer
		}
		return player
	})
	return nil
}

func (repo *playerMemRepo) GetAll(gameId gamemodel.GameIdVo) []playermodel.PlayerAgg {
	return lo.Filter(repo.players, func(player playermodel.PlayerAgg, _ int) bool {
		return player.GetGameId().IsEqual(gameId)
	})
}

func (repo *playerMemRepo) Delete(playerId playermodel.PlayerIdVo) error {
	repo.players = lo.Filter(repo.players, func(player playermodel.PlayerAgg, _ int) bool {
		return !player.GetId().IsEqual(playerId)
	})
	return nil
}
