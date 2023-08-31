package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ItemDto struct {
	Id                 uuid.UUID `json:"id"`
	CompatibleUnitType string    `json:"compatibleUnitType"`
	Name               string    `json:"name"`
	Traversable        bool      `json:"traversable"`
	AssetSrc           string    `json:"thumbnailSrc"`
	ModelSrc           string    `json:"modelSrc"`
}

func NewItemDto(item itemmodel.Item) ItemDto {
	return ItemDto{
		Id:                 item.GetId().Uuid(),
		CompatibleUnitType: item.GetCompatibleUnitType().String(),
		Name:               item.GetName(),
		Traversable:        item.GetTraversable(),
		AssetSrc:           item.GetThumbnailSrc(),
		ModelSrc:           item.GetModelSrc(),
	}
}
