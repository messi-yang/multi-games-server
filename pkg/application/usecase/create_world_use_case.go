package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	world_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreateWorldUseCase struct {
	worldService    service.WorldService
	worldMemberRepo worldaccessmodel.WorldMemberRepo
}

func NewCreateWorldUseCase(worldService service.WorldService, worldMemberRepo worldaccessmodel.WorldMemberRepo) CreateWorldUseCase {
	return CreateWorldUseCase{worldService, worldMemberRepo}
}

func ProvideCreateWorldUseCase(uow pguow.Uow) CreateWorldUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldAccountRepo := world_pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	worldRepo := world_pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	itemRepo := world_pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := world_pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	staticUnitRepo := world_pgrepo.NewStaticUnitRepo(uow, domainEventDispatcher)
	worldService := service.NewWorldService(worldAccountRepo, worldRepo, unitRepo, staticUnitRepo, itemRepo)

	worldMemberRepo := iam_pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)

	return NewCreateWorldUseCase(worldService, worldMemberRepo)
}

func (useCase *CreateWorldUseCase) Execute(useIdDto uuid.UUID, name string) (worldDto dto.WorldDto, err error) {
	userId := globalcommonmodel.NewUserId(useIdDto)
	newWorld, err := useCase.worldService.CreateWorld(userId, name)
	if err != nil {
		return worldDto, err
	}

	worldRole, err := globalcommonmodel.NewWorldRole("owner")
	if err != nil {
		return worldDto, err
	}
	newWorldMember := worldaccessmodel.NewWorldMember(
		newWorld.GetId(),
		userId,
		worldRole,
	)
	if err := useCase.worldMemberRepo.Add(newWorldMember); err != nil {
		return worldDto, err
	}

	return dto.NewWorldDto(newWorld), nil
}
