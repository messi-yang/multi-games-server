package game

import (
	"github.com/google/uuid"
)

type GameRepository interface {
	Add(Game) error
	Get(gameId uuid.UUID) (game Game, err error)
	Update(gameId uuid.UUID, game Game) error
	GetAll() []Game

	ReadLockAccess(gameId uuid.UUID) (rUnlocker func(), err error)
	LockAccess(gameId uuid.UUID) (unlocker func(), err error)
}
