package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type ItemVm struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AssetSrc string `json:"assetSrc"`
}

func NewItemVm(item itemmodel.ItemAgr) ItemVm {
	return ItemVm{
		Id:       item.GetId().ToString(),
		Name:     item.GetName(),
		AssetSrc: item.GetAssetSrc(),
	}
}
