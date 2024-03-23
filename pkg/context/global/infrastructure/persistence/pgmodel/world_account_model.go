package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/google/uuid"
)

type WorldAccountModel struct {
	Id               uuid.UUID `gorm:"primaryKey"`
	UserId           uuid.UUID `gorm:"unique;not null"`
	User             UserModel `gorm:"foreignKey:UserId;references:Id"`
	WorldsCount      int8      `gorm:"not null"`
	WorldsCountLimit int8      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;not null"`
}

func (WorldAccountModel) TableName() string {
	return "world_accounts"
}

func NewWorldAccountModel(worldAccount worldaccountmodel.WorldAccount) WorldAccountModel {
	return WorldAccountModel{
		Id:               worldAccount.GetId().Uuid(),
		UserId:           worldAccount.GetUserId().Uuid(),
		WorldsCount:      worldAccount.GetWorldsCount(),
		WorldsCountLimit: worldAccount.GetWorldsCountLimit(),
		CreatedAt:        worldAccount.GetCreatedAt(),
		UpdatedAt:        worldAccount.GetUpdatedAt(),
	}
}

func ParseWorldAccountModel(worldAccountModel WorldAccountModel) worldaccountmodel.WorldAccount {
	return worldaccountmodel.LoadWorldAccount(
		worldaccountmodel.NewWorldAccountId(worldAccountModel.Id),
		globalcommonmodel.NewUserId(worldAccountModel.UserId),
		worldAccountModel.WorldsCount,
		worldAccountModel.WorldsCountLimit,
		worldAccountModel.CreatedAt,
		worldAccountModel.UpdatedAt,
	)
}
