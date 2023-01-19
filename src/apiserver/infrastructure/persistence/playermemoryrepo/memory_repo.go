package playermemoryrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/samber/lo"
)

type memoryRepo struct {
	players map[playermodel.PlayerIdVo]playermodel.PlayerAgg
}

var singleton *memoryRepo

func New() playermodel.Repo {
	if singleton == nil {
		singleton = &memoryRepo{
			players: make(map[playermodel.PlayerIdVo]playermodel.PlayerAgg),
		}
		return singleton
	}
	return singleton
}

func (repo *memoryRepo) Add(player playermodel.PlayerAgg) {
	repo.players[player.GetId()] = player
}

func (repo *memoryRepo) GetAll() []playermodel.PlayerAgg {
	return lo.Values(repo.players)
}

func (repo *memoryRepo) Remove(playerId playermodel.PlayerIdVo) {
	delete(repo.players, playerId)
}
