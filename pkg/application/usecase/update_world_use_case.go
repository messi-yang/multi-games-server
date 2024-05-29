package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	world_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type UpdateWorldUseCase struct {
	worldRepo       worldmodel.WorldRepo
	worldService    service.WorldService
	worldMemberRepo worldaccessmodel.WorldMemberRepo
}

func NewUpdateWorldUseCase(worldRepo worldmodel.WorldRepo, worldService service.WorldService, worldMemberRepo worldaccessmodel.WorldMemberRepo) UpdateWorldUseCase {
	return UpdateWorldUseCase{worldRepo, worldService, worldMemberRepo}
}

func ProvideUpdateWorldUseCase(uow pguow.Uow) UpdateWorldUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldAccountRepo := world_pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	worldRepo := world_pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	itemRepo := world_pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := world_pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	staticUnitRepo := world_pgrepo.NewStaticUnitRepo(uow, domainEventDispatcher)
	worldService := service.NewWorldService(worldAccountRepo, worldRepo, unitRepo, staticUnitRepo, itemRepo)

	worldMemberRepo := iam_pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)

	return NewUpdateWorldUseCase(worldRepo, worldService, worldMemberRepo)
}

func (useCase *UpdateWorldUseCase) Execute(useIdDto uuid.UUID, worldIdDto uuid.UUID, worldName string) (
	updatedWorldDto dto.WorldDto, err error,
) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)
	userId := globalcommonmodel.NewUserId(useIdDto)

	worldMember, err := useCase.worldMemberRepo.GetWorldMemberOfUser(worldId, userId)
	if err != nil {
		return updatedWorldDto, err
	}

	if worldMember == nil {
		return updatedWorldDto, fmt.Errorf("you're not permitted to do this")
	}

	worldPermission := worldaccessmodel.NewWorldPermission(worldMember.GetRole())
	if !worldPermission.CanUpdateWorld() {
		return updatedWorldDto, fmt.Errorf("you're not permitted to do this")
	}

	world, err := useCase.worldRepo.Get(worldId)
	if err != nil {
		return updatedWorldDto, err
	}
	world.ChangeName(worldName)
	if err = useCase.worldRepo.Update(world); err != nil {
		return updatedWorldDto, err
	}

	updatedWorld, err := useCase.worldRepo.Get(worldId)
	if err != nil {
		return updatedWorldDto, err
	}

	return dto.NewWorldDto(updatedWorld), nil
}
