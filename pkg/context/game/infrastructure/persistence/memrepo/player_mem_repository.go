package memrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type playerMemRepo struct {
	players []playermodel.Player
}

var playerMemRepoSingleton *playerMemRepo

func NewPlayerMemRepo() playermodel.Repo {
	if playerMemRepoSingleton == nil {
		playerMemRepoSingleton = &playerMemRepo{
			players: make([]playermodel.Player, 0),
		}
		return playerMemRepoSingleton
	}
	return playerMemRepoSingleton
}

func (repo *playerMemRepo) Add(newPlayer playermodel.Player) error {
	newPlayers := append(repo.players, newPlayer)
	repo.players = lo.UniqBy(newPlayers, func(player playermodel.Player) uuid.UUID {
		return player.GetId().Uuid()
	})
	return nil
}

func (repo *playerMemRepo) Get(playerId commonmodel.PlayerId) (playermodel.Player, error) {
	foundPlayer, found := lo.Find(repo.players, func(player playermodel.Player) bool {
		return player.GetId().IsEqual(playerId)
	})
	if !found {
		return playermodel.Player{}, playermodel.ErrPlayerNotFound
	}
	return foundPlayer, nil
}

func (repo *playerMemRepo) FindPlayerAt(worldId commonmodel.WorldId, position commonmodel.Position) (playermodel.Player, bool, error) {
	foundPlayer, found := lo.Find(repo.players, func(player playermodel.Player) bool {
		return player.GetWorldId().IsEqual(worldId) && player.GetPosition().IsEqual(position)
	})
	if !found {
		return playermodel.Player{}, false, nil
	}
	return foundPlayer, true, nil
}

func (repo *playerMemRepo) GetPlayersAround(worldId commonmodel.WorldId, position commonmodel.Position) ([]playermodel.Player, error) {
	return lo.Filter(repo.players, func(player playermodel.Player, _ int) bool {
		return player.GetWorldId().IsEqual(worldId) && player.CanSeeAnyPositions([]commonmodel.Position{position})
	}), nil
}

func (repo *playerMemRepo) Update(updatedPlayer playermodel.Player) error {
	repo.players = lo.Map(repo.players, func(player playermodel.Player, _ int) playermodel.Player {
		if player.GetId().IsEqual(updatedPlayer.GetId()) {
			return updatedPlayer
		}
		return player
	})
	return nil
}

func (repo *playerMemRepo) GetAll(worldId commonmodel.WorldId) []playermodel.Player {
	return lo.Filter(repo.players, func(player playermodel.Player, _ int) bool {
		return player.GetWorldId().IsEqual(worldId)
	})
}

func (repo *playerMemRepo) Delete(playerId commonmodel.PlayerId) error {
	repo.players = lo.Filter(repo.players, func(player playermodel.Player, _ int) bool {
		return !player.GetId().IsEqual(playerId)
	})
	return nil
}
