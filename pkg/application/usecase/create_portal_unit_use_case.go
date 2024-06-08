package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreatePortalUnitUseCase struct {
	portalUnitService service.PortalUnitService
}

func NewCreatePortalUnitUseCase(portalUnitService service.PortalUnitService) CreatePortalUnitUseCase {
	return CreatePortalUnitUseCase{portalUnitService}
}

func ProvideCreatePortalUnitUseCase(uow pguow.Uow) CreatePortalUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	portalUnitRepoUnitService := service.NewPortalUnitService(worldRepo, unitRepo, portalUnitRepo, itemRepo)
	return NewCreatePortalUnitUseCase(portalUnitRepoUnitService)
}

func (useCase *CreatePortalUnitUseCase) Execute(idDto uuid.UUID, worldIdDto uuid.UUID,
	itemIdDto uuid.UUID, positionDto dto.PositionDto, directionDto int8) (err error) {
	return useCase.portalUnitService.CreatePortalUnit(
		portalunitmodel.NewPortalUnitId(idDto),
		globalcommonmodel.NewWorldId(worldIdDto),
		worldcommonmodel.NewItemId(itemIdDto),
		worldcommonmodel.NewPosition(positionDto.X, positionDto.Z),
		worldcommonmodel.NewDirection(directionDto),
	)
}
