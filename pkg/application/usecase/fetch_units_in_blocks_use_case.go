package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type FetchUnitsInBlocksUseCase struct {
	unitRepo unitmodel.UnitRepo
}

func NewFetchUnitsInBlocksUseCase(unitRepo unitmodel.UnitRepo) FetchUnitsInBlocksUseCase {
	return FetchUnitsInBlocksUseCase{unitRepo}
}

func ProvideFetchUnitsInBlocksUseCase(uow pguow.Uow) FetchUnitsInBlocksUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)

	return NewFetchUnitsInBlocksUseCase(unitRepo)
}

func (useCase *FetchUnitsInBlocksUseCase) Execute(worldIdDto uuid.UUID, blockDtos []dto.BlockDto) (
	unitDtos []dto.UnitDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)

	units := []unitmodel.Unit{}

	blocks := lo.Map(blockDtos, func(blockDto dto.BlockDto, _ int) worldcommonmodel.Block {
		return blockDto.ToValueObject()
	})

	for _, block := range blocks {
		unitsInBlock, err := useCase.unitRepo.GetUnitsInBlock(worldId, block)
		if err != nil {
			return unitDtos, err
		}
		units = append(units, unitsInBlock...)
	}

	unitDtos = lo.Map(units, func(unit unitmodel.Unit, _ int) dto.UnitDto {
		return dto.NewUnitDto(unit)
	})

	return unitDtos, nil
}
