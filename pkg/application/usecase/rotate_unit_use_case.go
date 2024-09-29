package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type RotateUnitUseCase struct {
	unitRepo          unitmodel.UnitRepo
	staticUnitService service.StaticUnitService
	fenceUnitService  service.FenceUnitService
	portalUnitService service.PortalUnitService
	linkUnitService   service.LinkUnitService
	embedUnitService  service.EmbedUnitService
}

func NewRotateUnitUseCase(unitRepo unitmodel.UnitRepo, staticUnitService service.StaticUnitService,
	fenceUnitService service.FenceUnitService, portalUnitService service.PortalUnitService, linkUnitService service.LinkUnitService, embedUnitService service.EmbedUnitService,
) RotateUnitUseCase {
	return RotateUnitUseCase{unitRepo, staticUnitService, fenceUnitService, portalUnitService, linkUnitService, embedUnitService}
}

func ProvideRotateUnitUseCase(uow pguow.Uow) RotateUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	staticUnitRepo := pgrepo.NewStaticUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepo := pgrepo.NewFenceUnitRepo(uow, domainEventDispatcher)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	linkUnitRepo := pgrepo.NewLinkUnitRepo(uow, domainEventDispatcher)
	embedUnitRepo := pgrepo.NewEmbedUnitRepo(uow, domainEventDispatcher)
	staticUnitService := service.NewStaticUnitService(unitRepo, staticUnitRepo, itemRepo)
	fenceUnitService := service.NewFenceUnitService(unitRepo, fenceUnitRepo, itemRepo)
	portalUnitService := service.NewPortalUnitService(unitRepo, portalUnitRepo, itemRepo)
	linkUnitService := service.NewLinkUnitService(unitRepo, linkUnitRepo, itemRepo)
	embedUnitService := service.NewEmbedUnitService(unitRepo, embedUnitRepo, itemRepo)
	return NewRotateUnitUseCase(unitRepo, staticUnitService, fenceUnitService, portalUnitService, linkUnitService, embedUnitService)
}

func (useCase *RotateUnitUseCase) Execute(unitIdDto uuid.UUID) error {
	unitId := unitmodel.NewUnitId(unitIdDto)
	unit, err := useCase.unitRepo.Get(unitId)
	if err != nil {
		return err
	}

	if unit.GetType().IsPortal() {
		return useCase.portalUnitService.RotatePortalUnit(portalunitmodel.NewPortalUnitId(unitIdDto))
	} else if unit.GetType().IsStatic() {
		return useCase.staticUnitService.RotateStaticUnit(staticunitmodel.NewStaticUnitId(unitIdDto))
	} else if unit.GetType().IsFence() {
		return useCase.fenceUnitService.RotateFenceUnit(fenceunitmodel.NewFenceUnitId(unitIdDto))
	} else if unit.GetType().IsLink() {
		return useCase.linkUnitService.RotateLinkUnit(linkunitmodel.NewLinkUnitId(unitIdDto))
	} else if unit.GetType().IsEmbed() {
		return useCase.embedUnitService.RotateEmbedUnit(embedunitmodel.NewEmbedUnitId(unitIdDto))
	}

	return nil
}
