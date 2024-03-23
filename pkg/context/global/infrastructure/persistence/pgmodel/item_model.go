package pgmodel

import (
	"fmt"
	"os"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	pq "github.com/lib/pq"
	"github.com/samber/lo"
)

type ItemModel struct {
	Id                 uuid.UUID      `gorm:"primaryKey"`
	CompatibleUnitType UnitTypeEnum   `gorm:"not null"`
	Name               string         `gorm:"not null"`
	DimensionWidth     int8           `gorm:"not null"`
	DimensionDepth     int8           `gorm:"not null"`
	Traversable        bool           `gorm:"not null"`
	ModelSources       pq.StringArray `gorm:"not null;type:text[]"`
	ThumbnailSrc       string         `gorm:"not null"`
	CreatedAt          time.Time      `gorm:"autoCreateTime;not null"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime;not null"`
}

func (ItemModel) TableName() string {
	return "items"
}

func NewItemModel(item itemmodel.Item) ItemModel {
	return ItemModel{
		Id:                 item.GetId().Uuid(),
		CompatibleUnitType: UnitTypeEnum(item.GetCompatibleUnitType().String()),
		Name:               item.GetName(),
		DimensionWidth:     item.GetDimension().GetWidth(),
		DimensionDepth:     item.GetDimension().GetDepth(),
		Traversable:        item.GetTraversable(),
		ThumbnailSrc:       item.GetThumbnailSrc(),
		ModelSources:       item.GetModelSources(),
	}
}

func ParseItemModel(itemModel ItemModel) (item itemmodel.Item, err error) {
	serverUrl := os.Getenv("SERVER_URL")
	compatibleUnitType, err := worldcommonmodel.NewUnitType(string(itemModel.CompatibleUnitType))
	if err != nil {
		return item, err
	}
	dimension, err := worldcommonmodel.NewDimension(itemModel.DimensionWidth, itemModel.DimensionDepth)
	if err != nil {
		return item, err
	}

	return itemmodel.LoadItem(
		worldcommonmodel.NewItemId(itemModel.Id),
		compatibleUnitType,
		itemModel.Name,
		dimension,
		itemModel.Traversable,
		fmt.Sprintf("%s%s", serverUrl, itemModel.ThumbnailSrc),
		lo.Map(itemModel.ModelSources, func(modelSource string, _ int) string {
			return fmt.Sprintf("%s%s", serverUrl, modelSource)
		}),
	), nil
}
