package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
)

type ItemVm struct {
	Id          int16  `json:"id"`
	Name        string `json:"name"`
	Traversable bool   `json:"traversable"`
	AssetSrc    string `json:"assetSrc"`
}

func NewItemVm(item itemmodel.ItemAgg) ItemVm {
	return ItemVm{
		Id:          item.GetId().ToInt16(),
		Name:        item.GetName(),
		Traversable: item.IsTraversable(),
		AssetSrc:    item.GetAssetSrc(),
	}
}
