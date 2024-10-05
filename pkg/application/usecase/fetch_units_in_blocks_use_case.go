package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/blockmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
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

func (useCase *FetchUnitsInBlocksUseCase) Execute(worldIdDto uuid.UUID, blockIdDtos []dto.BlockIdDto) (
	unitDtos []dto.UnitDto, blockDtos []dto.BlockDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)

	units := []unitmodel.Unit{}

	blocks := lo.Map(blockIdDtos, func(blockIdDto dto.BlockIdDto, _ int) blockmodel.Block {
		return blockmodel.LoadBlock(blockIdDto.ToValueObject())
	})

	for _, block := range blocks {
		unitsInBlock, err := useCase.unitRepo.GetUnitsInBlock(worldId, block)
		if err != nil {
			return unitDtos, blockDtos, err
		}
		units = append(units, unitsInBlock...)
	}

	unitDtos = lo.Map(units, func(unit unitmodel.Unit, _ int) dto.UnitDto {
		return dto.NewUnitDto(unit)
	})

	blockDtos = lo.Map(blocks, func(block blockmodel.Block, _ int) dto.BlockDto {
		return dto.NewBlockDto(block)
	})

	return unitDtos, blockDtos, nil
}
