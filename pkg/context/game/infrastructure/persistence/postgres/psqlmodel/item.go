package psqlmodel

import (
	"fmt"
	"os"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ItemModel struct {
	Id           uuid.UUID `gorm:"primaryKey;unique"`
	Name         string    `gorm:"not null"`
	Traversable  bool      `gorm:"not null"`
	ModelSrc     string    `gorm:"not null"`
	ThumbnailSrc string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;not null"`
}

func (ItemModel) TableName() string {
	return "items"
}

func NewItemModel(item itemmodel.ItemAgg) ItemModel {
	return ItemModel{
		Id:           item.GetId().Uuid(),
		Name:         item.GetName(),
		Traversable:  item.GetTraversable(),
		ThumbnailSrc: item.GetThumbnailSrc(),
		ModelSrc:     item.GetModelSrc(),
	}
}

func (model ItemModel) ToAggregate() itemmodel.ItemAgg {
	serverUrl := os.Getenv("SERVER_URL")
	return itemmodel.NewItemAgg(
		itemmodel.NewItemIdVo(model.Id),
		model.Name,
		model.Traversable,
		fmt.Sprintf("%s%s", serverUrl, model.ThumbnailSrc),
		fmt.Sprintf("%s%s", serverUrl, model.ModelSrc),
	)
}
