package itemviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type ItemViewModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func New(items []itemmodel.Item) []ItemViewModel {
	newItemViewModels := make([]ItemViewModel, 0)
	for _, item := range items {
		newItemViewModels = append(newItemViewModels, ItemViewModel{
			Id:   item.GetId().ToString(),
			Name: item.GetName(),
		})
	}
	return newItemViewModels
}
