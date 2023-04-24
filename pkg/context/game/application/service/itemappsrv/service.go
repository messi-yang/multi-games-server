package itemappsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(QueryItemsQuery) ([]jsondto.ItemAggDto, error)
}

type serve struct {
	itemRepo itemmodel.Repo
}

func NewService(itemRepo itemmodel.Repo) Service {
	return &serve{
		itemRepo: itemRepo,
	}
}

func (serve *serve) QueryItems(query QueryItemsQuery) (itemDtos []jsondto.ItemAggDto, err error) {
	items, err := serve.itemRepo.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(items, func(item itemmodel.ItemAgg, _ int) jsondto.ItemAggDto {
		return jsondto.NewItemAggDto(item)
	}), nil
}
