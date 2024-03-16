package embedunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
)

type Service interface {
	CreateEmbedUnit(CreateEmbedUnitCommand) error
	RemoveEmbedUnit(RemoveEmbedUnitCommand) error

	GetEmbedUnitEmbedCode(GetEmbedUnitEmbedCodeQuery) (string, error)
}

type serve struct {
	embedUnitRepo    embedunitmodel.EmbedUnitRepo
	embedUnitService service.EmbedUnitService
}

func NewService(
	embedUnitRepo embedunitmodel.EmbedUnitRepo,
	embedUnitService service.EmbedUnitService,
) Service {
	return &serve{
		embedUnitRepo:    embedUnitRepo,
		embedUnitService: embedUnitService,
	}
}

func (serve *serve) CreateEmbedUnit(command CreateEmbedUnitCommand) error {
	embedCode, err := worldcommonmodel.NewEmbedCode(command.EmbedCode)
	if err != nil {
		return err
	}

	return serve.embedUnitService.CreateEmbedUnit(
		embedunitmodel.NewEmbedUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
		command.Label,
		embedCode,
	)
}

func (serve *serve) RemoveEmbedUnit(command RemoveEmbedUnitCommand) error {
	return serve.embedUnitService.RemoveEmbedUnit(embedunitmodel.NewEmbedUnitId(command.Id))
}

func (serve *serve) GetEmbedUnitEmbedCode(query GetEmbedUnitEmbedCodeQuery) (string, error) {
	embedUnit, err := serve.embedUnitRepo.Get(embedunitmodel.NewEmbedUnitId(query.Id))
	if err != nil {
		return "", err
	}

	return embedUnit.GetEmbedCode().String(), nil
}
