package linkunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
)

type Service interface {
	CreateLinkUnit(CreateLinkUnitCommand) error
	RemoveLinkUnit(RemoveLinkUnitCommand) error

	GetLinkUnitUrl(GetLinkUnitUrlQuery) (string, error)
}

type serve struct {
	linkUnitRepo    linkunitmodel.LinkUnitRepo
	linkUnitService service.LinkUnitService
}

func NewService(
	linkUnitRepo linkunitmodel.LinkUnitRepo,
	linkUnitService service.LinkUnitService,
) Service {
	return &serve{
		linkUnitRepo:    linkUnitRepo,
		linkUnitService: linkUnitService,
	}
}

func (serve *serve) CreateLinkUnit(command CreateLinkUnitCommand) error {
	url, err := globalcommonmodel.NewUrl(command.Url)
	if err != nil {
		return err
	}

	return serve.linkUnitService.CreateLinkUnit(
		linkunitmodel.NewLinkUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
		url,
	)
}

func (serve *serve) RemoveLinkUnit(command RemoveLinkUnitCommand) error {
	return serve.linkUnitService.RemoveLinkUnit(linkunitmodel.NewLinkUnitId(command.Id))
}

func (serve *serve) GetLinkUnitUrl(query GetLinkUnitUrlQuery) (string, error) {
	linkUnit, err := serve.linkUnitRepo.Get(linkunitmodel.NewLinkUnitId(query.Id))
	if err != nil {
		return "", err
	}

	return linkUnit.GetUrl().String(), nil
}
