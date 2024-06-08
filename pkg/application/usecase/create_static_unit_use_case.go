package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreateStaticUnitUseCase struct {
	staticUnitService service.StaticUnitService
}

func NewCreateStaticUnitUseCase(staticUnitService service.StaticUnitService) CreateStaticUnitUseCase {
	return CreateStaticUnitUseCase{staticUnitService}
}

func ProvideCreateStaticUnitUseCase(uow pguow.Uow) CreateStaticUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	staticUnitRepo := pgrepo.NewStaticUnitRepo(uow, domainEventDispatcher)
	staticUnitRepoUnitService := service.NewStaticUnitService(worldRepo, unitRepo, staticUnitRepo, itemRepo)
	return NewCreateStaticUnitUseCase(staticUnitRepoUnitService)
}

func (useCase *CreateStaticUnitUseCase) Execute(idDto uuid.UUID, worldIdDto uuid.UUID,
	itemIdDto uuid.UUID, positionDto dto.PositionDto, directionDto int8) (err error) {
	return useCase.staticUnitService.CreateStaticUnit(
		staticunitmodel.NewStaticUnitId(idDto),
		globalcommonmodel.NewWorldId(worldIdDto),
		worldcommonmodel.NewItemId(itemIdDto),
		worldcommonmodel.NewPosition(positionDto.X, positionDto.Z),
		worldcommonmodel.NewDirection(directionDto),
	)
}
