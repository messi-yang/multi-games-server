package portalunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
)

type Service interface {
	CreatePortalUnit(CreatePortalUnitCommand) error
	RemovePortalUnit(RemovePortalUnitCommand) error
}

type serve struct {
	portalUnitRepo    portalunitmodel.PortalUnitRepo
	portalUnitService service.PortalUnitService
}

func NewService(
	portalUnitRepo portalunitmodel.PortalUnitRepo,
	portalUnitService service.PortalUnitService,
) Service {
	return &serve{
		portalUnitRepo:    portalUnitRepo,
		portalUnitService: portalUnitService,
	}
}

func (serve *serve) CreatePortalUnit(command CreatePortalUnitCommand) error {
	return serve.portalUnitService.CreatePortalUnit(
		portalunitmodel.NewPortalUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemovePortalUnit(command RemovePortalUnitCommand) error {
	return serve.portalUnitService.RemovePortalUnit(portalunitmodel.NewPortalUnitId(command.Id))
}
