package itemappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/itemmodel"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(QueryItemsQuery) ([]dto.ItemDto, error)
}

type serve struct {
	itemRepo itemmodel.ItemRepo
}

func NewService(itemRepo itemmodel.ItemRepo) Service {
	return &serve{
		itemRepo: itemRepo,
	}
}

func (serve *serve) QueryItems(query QueryItemsQuery) (itemDtos []dto.ItemDto, err error) {
	items, err := serve.itemRepo.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(items, func(item itemmodel.Item, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	}), nil
}
