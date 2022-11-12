package aggregate

import (
	commonValueObject "github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	gameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

type Game struct {
	id        gameValueObject.GameId
	unitBlock commonValueObject.UnitBlock
}

func NewGame(id gameValueObject.GameId, unitBlock commonValueObject.UnitBlock) Game {
	return Game{
		id:        id,
		unitBlock: unitBlock,
	}
}

func (game *Game) GetId() gameValueObject.GameId {
	return game.id
}
