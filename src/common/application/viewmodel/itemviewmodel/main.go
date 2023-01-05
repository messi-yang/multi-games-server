package itemviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type ViewModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AssetSrc string `json:"assetSrc"`
}

func New(item itemmodel.Item) ViewModel {
	return ViewModel{
		Id:       item.GetId().ToString(),
		Name:     item.GetName(),
		AssetSrc: item.GetAssetSrc(),
	}
}
