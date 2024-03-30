package staticunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
)

type Service interface {
	CreateStaticUnit(CreateStaticUnitCommand) error
	RemoveStaticUnit(RemoveStaticUnitCommand) error
}

type serve struct {
	staticUnitRepo    staticunitmodel.StaticUnitRepo
	staticUnitService service.StaticUnitService
}

func NewService(
	staticUnitRepo staticunitmodel.StaticUnitRepo,
	staticUnitService service.StaticUnitService,
) Service {
	return &serve{
		staticUnitRepo:    staticUnitRepo,
		staticUnitService: staticUnitService,
	}
}

func (serve *serve) CreateStaticUnit(command CreateStaticUnitCommand) error {
	return serve.staticUnitService.CreateStaticUnit(
		staticunitmodel.NewStaticUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemoveStaticUnit(command RemoveStaticUnitCommand) error {
	return serve.staticUnitService.RemoveStaticUnit(staticunitmodel.NewStaticUnitId(command.Id))
}
