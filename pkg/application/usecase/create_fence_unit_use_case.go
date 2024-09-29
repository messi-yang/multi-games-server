package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreateFenceUnitUseCase struct {
	fenceUnitService service.FenceUnitService
}

func NewCreateFenceUnitUseCase(fenceUnitService service.FenceUnitService) CreateFenceUnitUseCase {
	return CreateFenceUnitUseCase{fenceUnitService}
}

func ProvideCreateFenceUnitUseCase(uow pguow.Uow) CreateFenceUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepo := pgrepo.NewFenceUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepoUnitService := service.NewFenceUnitService(unitRepo, fenceUnitRepo, itemRepo)
	return NewCreateFenceUnitUseCase(fenceUnitRepoUnitService)
}

func (useCase *CreateFenceUnitUseCase) Execute(idDto uuid.UUID, worldIdDto uuid.UUID,
	itemIdDto uuid.UUID, positionDto dto.PositionDto, directionDto int8) (err error) {
	return useCase.fenceUnitService.CreateFenceUnit(
		fenceunitmodel.NewFenceUnitId(idDto),
		globalcommonmodel.NewWorldId(worldIdDto),
		worldcommonmodel.NewItemId(itemIdDto),
		worldcommonmodel.NewPosition(positionDto.X, positionDto.Z),
		worldcommonmodel.NewDirection(directionDto),
	)
}
