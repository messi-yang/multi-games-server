package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type ItemVm struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Traversable bool   `json:"traversable"`
	AssetSrc    string `json:"assetSrc"`
}

func NewItemVm(item itemmodel.ItemAgg) ItemVm {
	return ItemVm{
		Id:          item.GetId().ToString(),
		Name:        item.GetName(),
		Traversable: item.IsTraversable(),
		AssetSrc:    item.GetAssetSrc(),
	}
}
