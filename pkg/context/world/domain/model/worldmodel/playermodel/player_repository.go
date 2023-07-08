package playermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type PlayerRepo interface {
	Add(Player) error
	Update(Player) error
	Delete(Player) error
	Get(sharedkernelmodel.WorldId, PlayerId) (Player, error)
	GetPlayersOfWorld(sharedkernelmodel.WorldId) ([]Player, error)
}
