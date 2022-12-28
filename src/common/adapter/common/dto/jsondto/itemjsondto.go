package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/itemmodel"
)

type ItemJsonDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewItemJsonDtos(items []itemmodel.Item) []ItemJsonDto {
	newItemJsonDtos := make([]ItemJsonDto, 0)
	for _, item := range items {
		newItemJsonDtos = append(newItemJsonDtos, ItemJsonDto{
			Id:   item.GetId().ToString(),
			Name: item.GetName(),
		})
	}
	return newItemJsonDtos
}
