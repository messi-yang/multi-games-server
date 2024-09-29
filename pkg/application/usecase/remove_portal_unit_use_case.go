package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type RemovePortalUnitUseCase struct {
	portalUnitService service.PortalUnitService
}

func NewRemovePortalUnitUseCase(portalUnitService service.PortalUnitService) RemovePortalUnitUseCase {
	return RemovePortalUnitUseCase{portalUnitService}
}

func ProvideRemovePortalUnitUseCase(uow pguow.Uow) RemovePortalUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	portalUnitRepoUnitService := service.NewPortalUnitService(unitRepo, portalUnitRepo, itemRepo)
	return NewRemovePortalUnitUseCase(portalUnitRepoUnitService)
}

func (useCase *RemovePortalUnitUseCase) Execute(idDto uuid.UUID) error {
	return useCase.portalUnitService.RemovePortalUnit(portalunitmodel.NewPortalUnitId(idDto))
}
