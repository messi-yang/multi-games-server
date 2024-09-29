package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type RemoveFenceUnitUseCase struct {
	fenceUnitService service.FenceUnitService
}

func NewRemoveFenceUnitUseCase(fenceUnitService service.FenceUnitService) RemoveFenceUnitUseCase {
	return RemoveFenceUnitUseCase{fenceUnitService}
}

func ProvideRemoveFenceUnitUseCase(uow pguow.Uow) RemoveFenceUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepo := pgrepo.NewFenceUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepoUnitService := service.NewFenceUnitService(unitRepo, fenceUnitRepo, itemRepo)
	return NewRemoveFenceUnitUseCase(fenceUnitRepoUnitService)
}

func (useCase *RemoveFenceUnitUseCase) Execute(idDto uuid.UUID) error {
	return useCase.fenceUnitService.RemoveFenceUnit(fenceunitmodel.NewFenceUnitId(idDto))
}
