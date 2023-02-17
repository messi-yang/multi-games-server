package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
)

type ItemDto struct {
	Id          int16  `json:"id"`
	Name        string `json:"name"`
	Traversable bool   `json:"traversable"`
	AssetSrc    string `json:"assetSrc"`
	ModelSrc    string `json:"modelSrc"`
}

func NewItemDto(item itemmodel.ItemAgg) ItemDto {
	return ItemDto{
		Id:          item.GetId().ToInt16(),
		Name:        item.GetName(),
		Traversable: item.IsTraversable(),
		AssetSrc:    item.GetAssetSrc(),
		ModelSrc:    item.GetModelSrc(),
	}
}
