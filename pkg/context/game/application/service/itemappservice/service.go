package itemappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(QueryItemsQuery) ([]jsondto.ItemAggDto, error)
}

type serve struct {
	itemRepository itemmodel.Repository
}

func NewService(itemRepository itemmodel.Repository) Service {
	return &serve{
		itemRepository: itemRepository,
	}
}

func (serve *serve) QueryItems(query QueryItemsQuery) (itemDtos []jsondto.ItemAggDto, err error) {
	items, err := serve.itemRepository.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(items, func(item itemmodel.ItemAgg, _ int) jsondto.ItemAggDto {
		return jsondto.NewItemAggDto(item)
	}), nil
}
