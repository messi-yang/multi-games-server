package gamemodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"

type GameRepo interface {
	Add(Game) error
	Update(Game) error
	GetSelectedGameByRoomId(globalcommonmodel.RoomId) (Game, error)
	GetSelectedGamesByRoomId(globalcommonmodel.RoomId) ([]Game, error)
}
