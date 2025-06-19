package playermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type PlayerRepo interface {
	Add(Player) error
	Update(Player) error
	Delete(Player) error
	Get(globalcommonmodel.RoomId, PlayerId) (Player, error)
	GetPlayersOfRoom(globalcommonmodel.RoomId) ([]Player, error)
}
