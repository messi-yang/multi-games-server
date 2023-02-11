package gamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/samber/lo"
)

var (
	ErrSomeLocationsNotIncludedInMap = errors.New("some locations are not included in the unit map")
	ErrLocationHasPlayer             = errors.New("the location has player")
	ErrPlayerNotFound                = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists           = errors.New("the play with the given id already exists")
)

type GameAgg struct {
	id              GameIdVo
	players         map[PlayerIdVo]PlayerEntity
	playerLocations map[commonmodel.LocationVo]bool
}

func NewGameAgg(id GameIdVo) GameAgg {
	return GameAgg{
		id:      id,
		players: make(map[PlayerIdVo]PlayerEntity),
	}
}

func (game *GameAgg) GetId() GameIdVo {
	return game.id
}

func (game *GameAgg) GetPlayerIds() []PlayerIdVo {
	return lo.Keys(game.players)
}

func (game *GameAgg) GetPlayerIdsExcept(playerId PlayerIdVo) []PlayerIdVo {
	playerIds := lo.Keys(game.players)
	return lo.Filter(playerIds, func(pId PlayerIdVo, _ int) bool {
		return !pId.isEqual(playerId)
	})
}

func (game *GameAgg) GetPlayerViewBound(playerId PlayerIdVo) (commonmodel.BoundVo, error) {
	player, exists := game.players[playerId]
	if !exists {
		return commonmodel.BoundVo{}, ErrPlayerNotFound
	}

	playerLocation := player.GetLocation()

	fromX := playerLocation.GetX() - 25
	toX := playerLocation.GetX() + 25

	fromY := playerLocation.GetY() - 25
	toY := playerLocation.GetY() + 25

	from := commonmodel.NewLocationVo(fromX, fromY)
	to := commonmodel.NewLocationVo(toX, toY)
	bound, _ := commonmodel.NewBoundVo(from, to)

	return bound, nil
}

func (game *GameAgg) CanPlayerSeeAnyLocations(playerId PlayerIdVo, locations []commonmodel.LocationVo) bool {
	_, exists := game.players[playerId]
	if !exists {
		return false
	}

	bound, _ := game.GetPlayerViewBound(playerId)
	return bound.CoverAnyLocations(locations)
}

func (game *GameAgg) updatePlayerLocations() {
	players := game.GetPlayers()
	playerLocations := make(map[commonmodel.LocationVo]bool)
	lo.ForEach(players, func(player PlayerEntity, _ int) {
		playerLocations[player.GetLocation()] = true
	})
	game.playerLocations = playerLocations
}

func (game *GameAgg) DoesLocationHavePlayer(location commonmodel.LocationVo) bool {
	found := game.playerLocations[location]
	return found
}

func (game *GameAgg) AddPlayer(playerId PlayerIdVo) error {
	_, exists := game.players[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	playerLocation := commonmodel.NewLocationVo(0, 0)
	newPlayer := NewPlayerEntity(playerId, "Hello World", playerLocation)

	game.players[playerId] = newPlayer
	game.updatePlayerLocations()

	return nil
}

func (game *GameAgg) UpdatePlayer(player PlayerEntity) error {
	_, exists := game.players[player.id]
	if !exists {
		return ErrPlayerNotFound
	}
	game.players[player.id] = player
	game.updatePlayerLocations()
	return nil
}

func (game *GameAgg) GetPlayers() []PlayerEntity {
	return lo.Values(game.players)
}

func (game *GameAgg) GetPlayer(playerId PlayerIdVo) (PlayerEntity, error) {
	player, exists := game.players[playerId]
	if !exists {
		return PlayerEntity{}, ErrPlayerNotFound
	}
	return player, nil
}

func (game *GameAgg) RemovePlayer(playerId PlayerIdVo) {
	delete(game.players, playerId)
	game.updatePlayerLocations()
}
