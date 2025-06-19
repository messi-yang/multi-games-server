package gameaccountmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type GameAccount struct {
	id              GameAccountId
	userId          globalcommonmodel.UserId
	roomsCount      int8
	roomsCountLimit int8
	createdAt       time.Time
	updatedAt       time.Time
}

// Interface Implementation Check
var _ domain.Aggregate = (*GameAccount)(nil)

func NewGameAccount(
	userId globalcommonmodel.UserId,
) GameAccount {
	return GameAccount{
		id:              NewGameAccountId(uuid.New()),
		userId:          userId,
		roomsCount:      0,
		roomsCountLimit: 1,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}
}

func LoadGameAccount(
	id GameAccountId,
	userId globalcommonmodel.UserId,
	roomsCount int8,
	roomsCountLimit int8,
	createdAt time.Time,
	updatedAt time.Time,
) GameAccount {
	return GameAccount{
		id:              id,
		userId:          userId,
		roomsCount:      roomsCount,
		roomsCountLimit: roomsCountLimit,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}
}

func (gameAccount *GameAccount) GetId() GameAccountId {
	return gameAccount.id
}

func (gameAccount *GameAccount) GetUserId() globalcommonmodel.UserId {
	return gameAccount.userId
}

func (gameAccount *GameAccount) GetRoomsCount() int8 {
	return gameAccount.roomsCount
}

func (gameAccount *GameAccount) AddRoomsCount() {
	gameAccount.roomsCount += 1
}

func (gameAccount *GameAccount) SubtractRoomsCount() {
	gameAccount.roomsCount -= 1
}

func (gameAccount *GameAccount) GetRoomsCountLimit() int8 {
	return gameAccount.roomsCountLimit
}

func (gameAccount *GameAccount) CanAddNewRoom() bool {
	return gameAccount.GetRoomsCount() < gameAccount.GetRoomsCountLimit()
}

func (gameAccount *GameAccount) GetCreatedAt() time.Time {
	return gameAccount.createdAt
}

func (gameAccount *GameAccount) GetUpdatedAt() time.Time {
	return gameAccount.updatedAt
}
