package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
)

type Game struct {
	id                 GameId
	unitBlockDimension gamecommonmodel.Dimension
	unitBlock          gamecommonmodel.UnitBlock
}

func NewGame(id GameId, unitBlock gamecommonmodel.UnitBlock) Game {
	return Game{
		id:                 id,
		unitBlockDimension: unitBlock.GetDimension(),
		unitBlock:          unitBlock,
	}
}

func (game *Game) GetId() GameId {
	return game.id
}

func (game *Game) GetUnitBlock() gamecommonmodel.UnitBlock {
	return game.unitBlock
}

func (game *Game) GetUnitBlockDimension() gamecommonmodel.Dimension {
	return game.unitBlockDimension
}
