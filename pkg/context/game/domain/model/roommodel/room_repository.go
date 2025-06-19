package roommodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"

type RoomRepo interface {
	Add(Room) error
	Update(Room) error
	Delete(Room) error
	Get(globalcommonmodel.RoomId) (Room, error)
	GetRoomsOfUser(globalcommonmodel.UserId) ([]Room, error)
	Query(limit int, offset int) ([]Room, error)
}
