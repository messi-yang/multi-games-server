package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type Game struct {
	id                 GameId
	unitBlockDimension common.Dimension
	unitBlock          common.UnitBlock
}

func NewGame(id GameId, unitBlock common.UnitBlock) Game {
	return Game{
		id:                 id,
		unitBlockDimension: unitBlock.GetDimension(),
		unitBlock:          unitBlock,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetUnitBlock() common.UnitBlock {
	return game.unitBlock
}

func (game *Game) GetUnitBlockDimension() common.Dimension {
	return game.unitBlockDimension
}
