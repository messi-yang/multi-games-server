package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type RemoveEmbedUnitUseCase struct {
	embedUnitService service.EmbedUnitService
}

func NewRemoveEmbedUnitUseCase(embedUnitService service.EmbedUnitService) RemoveEmbedUnitUseCase {
	return RemoveEmbedUnitUseCase{embedUnitService}
}

func ProvideRemoveEmbedUnitUseCase(uow pguow.Uow) RemoveEmbedUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	embedUnitRepo := pgrepo.NewEmbedUnitRepo(uow, domainEventDispatcher)
	embedUnitRepoUnitService := service.NewEmbedUnitService(unitRepo, embedUnitRepo, itemRepo)
	return NewRemoveEmbedUnitUseCase(embedUnitRepoUnitService)
}

func (useCase *RemoveEmbedUnitUseCase) Execute(idDto uuid.UUID) error {
	return useCase.embedUnitService.RemoveEmbedUnit(embedunitmodel.NewEmbedUnitId(idDto))
}
