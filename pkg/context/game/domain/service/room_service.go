package service

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gameaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

var (
	ErrRoomsCountReachLimit = errors.New("rooms count has reached the limit")
	ErrDeleteNotWorking     = errors.New("room delete is not working now")
)

type RoomService interface {
	CreateRoom(userId globalcommonmodel.UserId, name string) (roommodel.Room, error)
	StartGame(roomId globalcommonmodel.RoomId, gameId gamemodel.GameId, gameState map[string]interface{}) (gamemodel.Game, error)
	SetupNewGame(roomId globalcommonmodel.RoomId, gameName string) (gamemodel.Game, error)
	DeleteRoom(roomId globalcommonmodel.RoomId) error
}

type roomServe struct {
	gameAccountRepo gameaccountmodel.GameAccountRepo
	roomRepo        roommodel.RoomRepo
	gameRepo        gamemodel.GameRepo
}

func NewRoomService(
	gameAccountRepo gameaccountmodel.GameAccountRepo,
	roomRepo roommodel.RoomRepo,
	gameRepo gamemodel.GameRepo,
) RoomService {
	return &roomServe{
		gameAccountRepo: gameAccountRepo,
		roomRepo:        roomRepo,
		gameRepo:        gameRepo,
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

	newRoom = roommodel.NewRoom(userId, nil, name)

	if err = roomServe.roomRepo.Add(newRoom); err != nil {
		return newRoom, err
	}

	game := gamemodel.NewGame(newRoom.GetId(), "maze_battle")
	err = roomServe.gameRepo.Add(game)
	if err != nil {
		return newRoom, err
	}
	newRoom.SetCurrentGameId(commonutil.ToPointer(game.GetId()))

	if err = roomServe.roomRepo.Update(newRoom); err != nil {
		return newRoom, err
	}

	return newRoom, nil
}

func (roomServe *roomServe) StartGame(roomId globalcommonmodel.RoomId, gameId gamemodel.GameId, gameState map[string]interface{}) (game gamemodel.Game, err error) {
	room, err := roomServe.roomRepo.Get(roomId)
	if err != nil {
		return game, err
	}

	roomCurrentGameId := room.GetCurrentGameId()
	if roomCurrentGameId == nil || roomCurrentGameId.Uuid() != gameId.Uuid() {
		return game, errors.New("game is not the current game of the room")
	}

	game, err = roomServe.gameRepo.Get(gameId)
	if err != nil {
		return game, err
	}

	if game.GetStarted() {
		return game, errors.New("game already started")
	}

	game.SetStarted(true)
	game.SetState(&gameState)

	err = roomServe.gameRepo.Update(game)
	if err != nil {
		return game, err
	}

	return game, nil
}

func (roomServe *roomServe) SetupNewGame(roomId globalcommonmodel.RoomId, gameName string) (newGame gamemodel.Game, err error) {
	room, err := roomServe.roomRepo.Get(roomId)
	if err != nil {
		return newGame, err
	}

	newGame = gamemodel.NewGame(roomId, gameName)
	room.SetCurrentGameId(commonutil.ToPointer(newGame.GetId()))

	err = roomServe.gameRepo.Add(newGame)
	if err != nil {
		return newGame, err
	}

	err = roomServe.roomRepo.Update(room)
	if err != nil {
		return newGame, err
	}

	return newGame, nil
}

func (roomServe *roomServe) DeleteRoom(roomId globalcommonmodel.RoomId) error {
	room, err := roomServe.roomRepo.Get(roomId)
	if err != nil {
		return err
	}
	room.Delete()
	return roomServe.roomRepo.Delete(room)
}
