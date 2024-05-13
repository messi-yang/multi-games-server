package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/samber/lo"
)

type QueryItemsUseCase struct {
	itemRepo itemmodel.ItemRepo
}

func NewQueryItemsUseCase(itemRepo itemmodel.ItemRepo) QueryItemsUseCase {
	return QueryItemsUseCase{itemRepo}
}

func ProvideQueryItemsUseCase(uow pguow.Uow) QueryItemsUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)

	return NewQueryItemsUseCase(itemRepo)
}

func (useCase *QueryItemsUseCase) Execute() (itemDtos []dto.ItemDto, err error) {
	items, err := useCase.itemRepo.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(items, func(item itemmodel.Item, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	}), nil
}
