package playermodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

var (
	ErrPlayerNotFound   = errors.New("player of the given id not found")
	ErrSomethinHappened = errors.New("some unexpected error happened")
)

type Repo interface {
	Add(Player) error
	Update(Player) error
	Delete(Player) error
	Get(commonmodel.PlayerId) (Player, error)
	FindPlayersAt(sharedkernelmodel.WorldId, commonmodel.Position) (players []Player, found bool, err error)
	GetPlayersAround(sharedkernelmodel.WorldId, commonmodel.Position) ([]Player, error)
}
