package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"
)

type ItemJsonDto struct {
	Id   ItemIdJsonDto `json:"id"`
	Name string        `json:"name"`
}

func NewItemJsonDtos(items []itemmodel.Item) []ItemJsonDto {
	newItemJsonDtos := make([]ItemJsonDto, 0)
	for _, item := range items {
		newItemJsonDtos = append(newItemJsonDtos, ItemJsonDto{
			Id:   NewItemIdJsonDto(item.GetId()),
			Name: item.GetName(),
		})
	}
	return newItemJsonDtos
}
