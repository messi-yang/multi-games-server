package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ItemDto struct {
	Id                 uuid.UUID    `json:"id"`
	CompatibleUnitType string       `json:"compatibleUnitType"`
	Name               string       `json:"name"`
	Dimension          DimensionDto `json:"dimension"`
	Traversable        bool         `json:"traversable"`
	AssetSrc           string       `json:"thumbnailSrc"`
	ModelSources       []string     `json:"modelSources"`
}

func NewItemDto(item itemmodel.Item) ItemDto {
	return ItemDto{
		Id:                 item.GetId().Uuid(),
		CompatibleUnitType: item.GetCompatibleUnitType().String(),
		Name:               item.GetName(),
		Dimension:          NewDimensionDto(item.GetDimension()),
		Traversable:        item.GetTraversable(),
		AssetSrc:           item.GetThumbnailSrc(),
		ModelSources:       item.GetModelSources(),
	}
}
