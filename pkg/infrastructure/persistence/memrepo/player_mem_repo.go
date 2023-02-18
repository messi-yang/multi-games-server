package memrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
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
func (repo *playerMemRepo) GetPlayerAt(location commonmodel.LocationVo) (playermodel.PlayerAgg, error) {
	foundPlayer, found := lo.Find(repo.players, func(player playermodel.PlayerAgg) bool {
		return player.GetLocation().IsEqual(location)
	})
	if !found {
		return playermodel.PlayerAgg{}, playermodel.ErrPlayerNotFound
	}
	return foundPlayer, nil
}
func (repo *playerMemRepo) Update(updatedPlayer playermodel.PlayerAgg) error {
	lo.Map(repo.players, func(player playermodel.PlayerAgg, _ int) playermodel.PlayerAgg {
		if player.GetId().IsEqual(updatedPlayer.GetId()) {
			return updatedPlayer
		}
		return player
	})
	return nil
}
func (repo *playerMemRepo) GetAll() []playermodel.PlayerAgg {
	return repo.players
}
func (repo *playerMemRepo) Delete(playerId playermodel.PlayerIdVo) {
	repo.players = lo.Filter(repo.players, func(player playermodel.PlayerAgg, _ int) bool {
		return !player.GetId().IsEqual(playerId)
	})
}
func (repo *playerMemRepo) ReadLockAccess(playermodel.PlayerIdVo) (rUnlocker func()) {
	return nil
}
func (repo *playerMemRepo) LockAccess(playermodel.PlayerIdVo) (unlocker func()) {
	return nil
}
