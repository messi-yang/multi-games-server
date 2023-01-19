package memrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/samber/lo"
)

type playerMemRepo struct {
	players map[playermodel.PlayerIdVo]playermodel.PlayerAgg
}

var playerMemRepoSingleton *playerMemRepo

func NewPlayerMemRepo() playermodel.Repo {
	if playerMemRepoSingleton == nil {
		playerMemRepoSingleton = &playerMemRepo{
			players: make(map[playermodel.PlayerIdVo]playermodel.PlayerAgg),
		}
		return playerMemRepoSingleton
	}
	return playerMemRepoSingleton
}

func (repo *playerMemRepo) Add(player playermodel.PlayerAgg) {
	repo.players[player.GetId()] = player
}

func (repo *playerMemRepo) GetAll() []playermodel.PlayerAgg {
	return lo.Values(repo.players)
}

func (repo *playerMemRepo) Remove(playerId playermodel.PlayerIdVo) {
	delete(repo.players, playerId)
}
