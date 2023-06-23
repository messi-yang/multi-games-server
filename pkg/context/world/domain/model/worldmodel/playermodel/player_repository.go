package playermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
)

type PlayerRepo interface {
	Add(Player) error
	Update(Player) error
	Delete(Player) error
	Get(PlayerId) (Player, error)
	FindPlayersAt(sharedkernelmodel.WorldId, commonmodel.Position) (players []Player, found bool, err error)
	GetPlayersOfWorld(sharedkernelmodel.WorldId) ([]Player, error)
}
