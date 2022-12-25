package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
)

type Game struct {
	id                 GameId
	unitBlockDimension commonmodel.Dimension
	unitBlock          commonmodel.UnitBlock
}

func NewGame(id GameId, unitBlock commonmodel.UnitBlock) Game {
	return Game{
		id:                 id,
		unitBlockDimension: unitBlock.GetDimension(),
		unitBlock:          unitBlock,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetUnitBlock() commonmodel.UnitBlock {
	return game.unitBlock
}

func (game *Game) GetUnitBlockDimension() commonmodel.Dimension {
	return game.unitBlockDimension
}
