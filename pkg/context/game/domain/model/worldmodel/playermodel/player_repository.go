package playermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type PlayerRepo interface {
	Add(Player) error
	Update(Player) error
	Delete(Player) error
	Get(PlayerId) (Player, error)
	FindPlayersAt(sharedkernelmodel.WorldId, commonmodel.Position) (players []Player, found bool, err error)
	GetPlayersAround(sharedkernelmodel.WorldId, commonmodel.Position) ([]Player, error)
}
