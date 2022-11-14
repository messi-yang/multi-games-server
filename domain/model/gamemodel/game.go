package gamemodel

import (
	commonValueObject "github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
)

type Game struct {
	id                 GameId
	unitBlockDimension commonValueObject.Dimension
	unitBlock          commonValueObject.UnitBlock
}

func NewGame(id GameId, unitBlock commonValueObject.UnitBlock) Game {
	return Game{
		id:                 id,
		unitBlockDimension: unitBlock.GetDimension(),
		unitBlock:          unitBlock,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetUnitBlock() commonValueObject.UnitBlock {
	return game.unitBlock
}

func (game *Game) GetUnitBlockDimension() commonValueObject.Dimension {
	return game.unitBlockDimension
}
