package itemviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type ViewModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	AssetSrc string `json:"assetSrc"`
}

func BatchNew(items []itemmodel.Item) []ViewModel {
	newViewModels := make([]ViewModel, 0)
	for _, item := range items {
		newViewModel := ViewModel{
			Id:       item.GetId().ToString(),
			Name:     item.GetName(),
			AssetSrc: item.GetAssetSrc(),
		}
		newViewModels = append(newViewModels, newViewModel)
	}
	return newViewModels
}
