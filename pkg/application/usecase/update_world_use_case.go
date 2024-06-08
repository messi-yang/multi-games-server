package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	world_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type UpdateWorldUseCase struct {
	worldRepo       worldmodel.WorldRepo
	worldMemberRepo worldaccessmodel.WorldMemberRepo
}

func NewUpdateWorldUseCase(worldRepo worldmodel.WorldRepo, worldMemberRepo worldaccessmodel.WorldMemberRepo) UpdateWorldUseCase {
	return UpdateWorldUseCase{worldRepo, worldMemberRepo}
}

func ProvideUpdateWorldUseCase(uow pguow.Uow) UpdateWorldUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldRepo := world_pgrepo.NewWorldRepo(uow, domainEventDispatcher)

	worldMemberRepo := iam_pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)

	return NewUpdateWorldUseCase(worldRepo, worldMemberRepo)
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
