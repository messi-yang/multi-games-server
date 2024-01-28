package itemappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(QueryItemsQuery) ([]dto.ItemDto, error)
	GetItemsOfIds(GetItemsOfIdsQuery) ([]dto.ItemDto, error)
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

func (serve *serve) GetItemsOfIds(query GetItemsOfIdsQuery) (itemDtos []dto.ItemDto, err error) {
	itemIds := lo.Map(query.ItemIds, func(itemIdDto uuid.UUID, _ int) worldcommonmodel.ItemId {
		return worldcommonmodel.NewItemId(itemIdDto)
	})

	items, err := serve.itemRepo.GetItemsOfIds(itemIds)
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(items, func(item itemmodel.Item, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	}), nil
}
