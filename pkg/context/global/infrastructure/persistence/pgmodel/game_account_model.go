package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gameaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type GameAccountModel struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	UserId          uuid.UUID `gorm:"unique;not null"`
	User            UserModel `gorm:"foreignKey:UserId;references:Id"`
	RoomsCount      int8      `gorm:"not null"`
	RoomsCountLimit int8      `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime;not null"`
}

func (GameAccountModel) TableName() string {
	return "game_accounts"
}

func NewGameAccountModel(gameAccount gameaccountmodel.GameAccount) GameAccountModel {
	return GameAccountModel{
		Id:              gameAccount.GetId().Uuid(),
		UserId:          gameAccount.GetUserId().Uuid(),
		RoomsCount:      gameAccount.GetRoomsCount(),
		RoomsCountLimit: gameAccount.GetRoomsCountLimit(),
		CreatedAt:       gameAccount.GetCreatedAt(),
		UpdatedAt:       gameAccount.GetUpdatedAt(),
	}
}

func ParseGameAccountModel(gameAccountModel GameAccountModel) gameaccountmodel.GameAccount {
	return gameaccountmodel.LoadGameAccount(
		gameaccountmodel.NewGameAccountId(gameAccountModel.Id),
		globalcommonmodel.NewUserId(gameAccountModel.UserId),
		gameAccountModel.RoomsCount,
		gameAccountModel.RoomsCountLimit,
		gameAccountModel.CreatedAt,
		gameAccountModel.UpdatedAt,
	)
}
