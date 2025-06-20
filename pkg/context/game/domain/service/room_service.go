package service

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gameaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

var (
	ErrRoomsCountReachLimit = errors.New("rooms count has reached the limit")
	ErrDeleteNotWorking     = errors.New("room delete is not working now")
)

type RoomService interface {
	CreateRoom(userId globalcommonmodel.UserId, name string) (roommodel.Room, error)
	DeleteRoom(roomId globalcommonmodel.RoomId) error
}

type roomServe struct {
	gameAccountRepo gameaccountmodel.GameAccountRepo
	roomRepo        roommodel.RoomRepo
}

func NewRoomService(
	gameAccountRepo gameaccountmodel.GameAccountRepo,
	roomRepo roommodel.RoomRepo,
) RoomService {
	return &roomServe{
		gameAccountRepo: gameAccountRepo,
		roomRepo:        roomRepo,
	}
}

func (roomServe *roomServe) CreateRoom(userId globalcommonmodel.UserId, name string) (newRoom roommodel.Room, err error) {
	gameAccount, err := roomServe.gameAccountRepo.GetGameAccountOfUser(userId)
	if err != nil {
		return newRoom, err
	}
	if !gameAccount.CanAddNewRoom() {
		return newRoom, ErrRoomsCountReachLimit
	}

	newRoom = roommodel.NewRoom(userId, name)

	if err = roomServe.roomRepo.Add(newRoom); err != nil {
		return newRoom, err
	}

	return newRoom, nil
}

func (roomServe *roomServe) DeleteRoom(roomId globalcommonmodel.RoomId) error {
	room, err := roomServe.roomRepo.Get(roomId)
	if err != nil {
		return err
	}
	room.Delete()
	return roomServe.roomRepo.Delete(room)
}
