package unitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
)

type Service interface {
	RotateUnit(RotateUnitCommand) error
}

type serve struct {
	worldRepo         worldmodel.WorldRepo
	unitRepo          unitmodel.UnitRepo
	itemRepo          itemmodel.ItemRepo
	staticUnitService service.StaticUnitService
	fenceUnitService  service.FenceUnitService
	portalUnitService service.PortalUnitService
	linkUnitService   service.LinkUnitService
	embedUnitService  service.EmbedUnitService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	staticUnitService service.StaticUnitService,
	fenceUnitService service.FenceUnitService,
	portalUnitService service.PortalUnitService,
	linkUnitService service.LinkUnitService,
	embedUnitService service.EmbedUnitService,
) Service {
	return &serve{
		worldRepo:         worldRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		staticUnitService: staticUnitService,
		fenceUnitService:  fenceUnitService,
		portalUnitService: portalUnitService,
		linkUnitService:   linkUnitService,
		embedUnitService:  embedUnitService,
	}
}

func (serve *serve) RotateUnit(command RotateUnitCommand) error {
	unitId := unitmodel.NewUnitId(command.Id)
	unit, err := serve.unitRepo.Get(unitId)
	if err != nil {
		return err
	}

	if unit.GetType().IsPortal() {
		return serve.portalUnitService.RotatePortalUnit(portalunitmodel.NewPortalUnitId(command.Id))
	} else if unit.GetType().IsStatic() {
		return serve.staticUnitService.RotateStaticUnit(staticunitmodel.NewStaticUnitId(command.Id))
	} else if unit.GetType().IsFence() {
		return serve.fenceUnitService.RotateFenceUnit(fenceunitmodel.NewFenceUnitId(command.Id))
	} else if unit.GetType().IsLink() {
		return serve.linkUnitService.RotateLinkUnit(linkunitmodel.NewLinkUnitId(command.Id))
	} else if unit.GetType().IsEmbed() {
		return serve.embedUnitService.RotateEmbedUnit(embedunitmodel.NewEmbedUnitId(command.Id))
	}

	return nil
}
