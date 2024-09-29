package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type RemoveStaticUnitUseCase struct {
	staticUnitService service.StaticUnitService
}

func NewRemoveStaticUnitUseCase(staticUnitService service.StaticUnitService) RemoveStaticUnitUseCase {
	return RemoveStaticUnitUseCase{staticUnitService}
}

func ProvideRemoveStaticUnitUseCase(uow pguow.Uow) RemoveStaticUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	staticUnitRepo := pgrepo.NewStaticUnitRepo(uow, domainEventDispatcher)
	staticUnitRepoUnitService := service.NewStaticUnitService(unitRepo, staticUnitRepo, itemRepo)
	return NewRemoveStaticUnitUseCase(staticUnitRepoUnitService)
}

func (useCase *RemoveStaticUnitUseCase) Execute(idDto uuid.UUID) error {
	return useCase.staticUnitService.RemoveStaticUnit(staticunitmodel.NewStaticUnitId(idDto))
}
