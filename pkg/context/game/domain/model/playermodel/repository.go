package playermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

var (
	ErrPlayerNotFound   = errors.New("player of the given id not found")
	ErrSomethinHappened = errors.New("some unexpected error happened")
)

type Repo interface {
	Add(Player) error
	Get(commonmodel.PlayerId) (Player, error)
	Update(Player) error
	Delete(commonmodel.PlayerId) error
	FindPlayerAt(commonmodel.WorldId, commonmodel.Position) (player Player, found bool, err error)
	GetPlayersAround(commonmodel.WorldId, commonmodel.Position) ([]Player, error)
}
