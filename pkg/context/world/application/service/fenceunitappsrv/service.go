package fenceunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
)

type Service interface {
	CreateFenceUnit(CreateFenceUnitCommand) error
	RemoveFenceUnit(RemoveFenceUnitCommand) error
}

type serve struct {
	fenceUnitRepo    fenceunitmodel.FenceUnitRepo
	fenceUnitService service.FenceUnitService
}

func NewService(
	fenceUnitRepo fenceunitmodel.FenceUnitRepo,
	fenceUnitService service.FenceUnitService,
) Service {
	return &serve{
		fenceUnitRepo:    fenceUnitRepo,
		fenceUnitService: fenceUnitService,
	}
}

func (serve *serve) CreateFenceUnit(command CreateFenceUnitCommand) error {
	return serve.fenceUnitService.CreateFenceUnit(
		fenceunitmodel.NewFenceUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemoveFenceUnit(command RemoveFenceUnitCommand) error {
	return serve.fenceUnitService.RemoveFenceUnit(fenceunitmodel.NewFenceUnitId(command.Id))
}
