package repository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/google/uuid"
)

type GameRepository interface {
	Add(entity.Game) error
	Get(gameId uuid.UUID) (game entity.Game, err error)
	GetFirstGameId() (gameId uuid.UUID, err error)
	Update(gameId uuid.UUID, game entity.Game) error
}
