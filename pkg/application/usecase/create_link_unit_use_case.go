package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreateLinkUnitUseCase struct {
	linkUnitService service.LinkUnitService
}

func NewCreateLinkUnitUseCase(linkUnitService service.LinkUnitService) CreateLinkUnitUseCase {
	return CreateLinkUnitUseCase{linkUnitService}
}

func ProvideCreateLinkUnitUseCase(uow pguow.Uow) CreateLinkUnitUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	linkUnitRepo := pgrepo.NewLinkUnitRepo(uow, domainEventDispatcher)
	linkUnitRepoUnitService := service.NewLinkUnitService(worldRepo, unitRepo, linkUnitRepo, itemRepo)
	return NewCreateLinkUnitUseCase(linkUnitRepoUnitService)
}

func (useCase *CreateLinkUnitUseCase) Execute(idDto uuid.UUID, worldIdDto uuid.UUID,
	itemIdDto uuid.UUID, positionDto dto.PositionDto, directionDto int8, labelDto *string, urlDto string) (err error) {
	url, err := globalcommonmodel.NewUrl(urlDto)
	if err != nil {
		return err
	}
	return useCase.linkUnitService.CreateLinkUnit(
		linkunitmodel.NewLinkUnitId(idDto),
		globalcommonmodel.NewWorldId(worldIdDto),
		worldcommonmodel.NewItemId(itemIdDto),
		worldcommonmodel.NewPosition(positionDto.X, positionDto.Z),
		worldcommonmodel.NewDirection(directionDto),
		labelDto,
		url,
	)
}
