package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type Item struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AssetSrc string `json:"assetSrc"`
}

func NewItem(item itemmodel.Item) Item {
	return Item{
		Id:       item.GetId().ToString(),
		Name:     item.GetName(),
		AssetSrc: item.GetAssetSrc(),
	}
}
