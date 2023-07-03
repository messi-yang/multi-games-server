package playermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
)

type PlayerRepo interface {
	Add(Player) error
	Update(Player) error
	Delete(Player) error
	Get(sharedkernelmodel.WorldId, PlayerId) (Player, error)
	GetPlayersAt(sharedkernelmodel.WorldId, commonmodel.Position) (players []Player, err error)
	GetPlayersOfWorld(sharedkernelmodel.WorldId) ([]Player, error)
}
