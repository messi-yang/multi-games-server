package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Game struct {
	id             GameId
	gameMapMapSize commonmodel.MapSize
	gameMap        commonmodel.GameMap
}

func NewGame(id GameId, gameMap commonmodel.GameMap) Game {
	return Game{
		id:             id,
		gameMapMapSize: gameMap.GetMapSize(),
		gameMap:        gameMap,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetGameMap() commonmodel.GameMap {
	return game.gameMap
}

func (game *Game) GetGameMapMapSize() commonmodel.MapSize {
	return game.gameMapMapSize
}
