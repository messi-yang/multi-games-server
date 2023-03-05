package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
)

type ItemDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Traversable bool   `json:"traversable"`
	AssetSrc    string `json:"thumbnailSrc"`
	ModelSrc    string `json:"modelSrc"`
}

func NewItemDto(item itemmodel.ItemAgg) ItemDto {
	return ItemDto{
		Id:          item.GetId().String(),
		Name:        item.GetName(),
		Traversable: item.GetTraversable(),
		AssetSrc:    item.GetThumbnailSrc(),
		ModelSrc:    item.GetModelSrc(),
	}
}
