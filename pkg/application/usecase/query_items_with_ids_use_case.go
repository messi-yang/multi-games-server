package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetItemsWithIdsUseCase struct {
	itemRepo itemmodel.ItemRepo
}

func NewGetItemsWithIdsUseCase(itemRepo itemmodel.ItemRepo) GetItemsWithIdsUseCase {
	return GetItemsWithIdsUseCase{itemRepo}
}

func ProvideGetItemsWithIdsUseCase(uow pguow.Uow) GetItemsWithIdsUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)

	return NewGetItemsWithIdsUseCase(itemRepo)
}

func (useCase *GetItemsWithIdsUseCase) Execute(itemIdDtos []uuid.UUID) (itemDtos []dto.ItemDto, err error) {
	itemIds := lo.Map(itemIdDtos, func(itemIdDto uuid.UUID, _ int) worldcommonmodel.ItemId {
		return worldcommonmodel.NewItemId(itemIdDto)
	})

	items, err := useCase.itemRepo.GetItemsWithIds(itemIds)
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(items, func(item itemmodel.Item, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	}), nil
}
