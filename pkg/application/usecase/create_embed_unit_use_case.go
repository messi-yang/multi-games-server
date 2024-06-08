package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreateEmbedUnitUseCase struct {
	embedUnitService service.EmbedUnitService
}

func NewCreateEmbedUnitUseCase(embedUnitService service.EmbedUnitService) CreateEmbedUnitUseCase {
	return CreateEmbedUnitUseCase{embedUnitService}
}

func ProvideCreateEmbedUnitUseCase(uow pguow.Uow) CreateEmbedUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	embedUnitRepo := pgrepo.NewEmbedUnitRepo(uow, domainEventDispatcher)
	embedUnitRepoUnitService := service.NewEmbedUnitService(worldRepo, unitRepo, embedUnitRepo, itemRepo)
	return NewCreateEmbedUnitUseCase(embedUnitRepoUnitService)
}

func (useCase *CreateEmbedUnitUseCase) Execute(idDto uuid.UUID, worldIdDto uuid.UUID,
	itemIdDto uuid.UUID, positionDto dto.PositionDto, directionDto int8, labelDto *string, embedCodeDto string) (err error) {
	embedCode, err := worldcommonmodel.NewEmbedCode(embedCodeDto)
	if err != nil {
		return err
	}

	return useCase.embedUnitService.CreateEmbedUnit(
		embedunitmodel.NewEmbedUnitId(idDto),
		globalcommonmodel.NewWorldId(worldIdDto),
		worldcommonmodel.NewItemId(itemIdDto),
		worldcommonmodel.NewPosition(positionDto.X, positionDto.Z),
		worldcommonmodel.NewDirection(directionDto),
		labelDto,
		embedCode,
	)
}
