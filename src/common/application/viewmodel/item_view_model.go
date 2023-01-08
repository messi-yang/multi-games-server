package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type ItemViewModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AssetSrc string `json:"assetSrc"`
}

func NewItemViewModel(item itemmodel.Item) ItemViewModel {
	return ItemViewModel{
		Id:       item.GetId().ToString(),
		Name:     item.GetName(),
		AssetSrc: item.GetAssetSrc(),
	}
}
