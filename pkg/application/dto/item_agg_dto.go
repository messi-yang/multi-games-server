package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ItemAggDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Traversable bool      `json:"traversable"`
	AssetSrc    string    `json:"thumbnailSrc"`
	ModelSrc    string    `json:"modelSrc"`
}

func NewItemAggDto(item itemmodel.ItemAgg) ItemAggDto {
	return ItemAggDto{
		Id:          item.GetId().Uuid(),
		Name:        item.GetName(),
		Traversable: item.GetTraversable(),
		AssetSrc:    item.GetThumbnailSrc(),
		ModelSrc:    item.GetModelSrc(),
	}
}
