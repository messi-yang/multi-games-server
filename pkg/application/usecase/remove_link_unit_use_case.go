package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type RemoveLinkUnitUseCase struct {
	linkUnitService service.LinkUnitService
}

func NewRemoveLinkUnitUseCase(linkUnitService service.LinkUnitService) RemoveLinkUnitUseCase {
	return RemoveLinkUnitUseCase{linkUnitService}
}

func ProvideRemoveLinkUnitUseCase(uow pguow.Uow) RemoveLinkUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	linkUnitRepo := pgrepo.NewLinkUnitRepo(uow, domainEventDispatcher)
	linkUnitRepoUnitService := service.NewLinkUnitService(unitRepo, linkUnitRepo, itemRepo)
	return NewRemoveLinkUnitUseCase(linkUnitRepoUnitService)
}

func (useCase *RemoveLinkUnitUseCase) Execute(idDto uuid.UUID) error {
	return useCase.linkUnitService.RemoveLinkUnit(linkunitmodel.NewLinkUnitId(idDto))
}
