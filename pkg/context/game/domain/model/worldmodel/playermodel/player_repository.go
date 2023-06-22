package playermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type PlayerRepo interface {
	Add(Player) error
	Update(Player) error
	Delete(Player) error
	Get(PlayerId) (Player, error)
	FindPlayersAt(sharedkernelmodel.WorldId, commonmodel.Position) (players []Player, found bool, err error)
	GetPlayersOfWorld(sharedkernelmodel.WorldId) ([]Player, error)
}
