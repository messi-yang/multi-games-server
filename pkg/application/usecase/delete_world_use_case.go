package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	world_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type DeleteWorldUseCase struct {
	worldService    service.WorldService
	worldMemberRepo worldaccessmodel.WorldMemberRepo
}

func NewDeleteWorldUseCase(worldService service.WorldService, worldMemberRepo worldaccessmodel.WorldMemberRepo) DeleteWorldUseCase {
	return DeleteWorldUseCase{worldService, worldMemberRepo}
}

func ProvideDeleteWorldUseCase(uow pguow.Uow) DeleteWorldUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	worldAccountRepo := world_pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	worldRepo := world_pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	worldService := service.NewWorldService(worldAccountRepo, worldRepo)

	worldMemberRepo := iam_pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)

	return NewDeleteWorldUseCase(worldService, worldMemberRepo)
}

func (useCase *DeleteWorldUseCase) Execute(useIdDto uuid.UUID, worldIdDto uuid.UUID) (err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)
	userId := globalcommonmodel.NewUserId(useIdDto)
	worldMember, err := useCase.worldMemberRepo.GetWorldMemberOfUser(worldId, userId)
	if err != nil {
		return err
	}

	if worldMember == nil {
		return fmt.Errorf("you're not permitted to do this")
	}

	worldPermission := worldaccessmodel.NewWorldPermission(worldMember.GetRole())
	if !worldPermission.CanDeleteWorld() {
		return fmt.Errorf("you're not permitted to do this")
	}

	// TODO - handle this side effects by using integration events
	worldMembersInWorld, err := useCase.worldMemberRepo.GetWorldMembersInWorld(worldId)
	if err != nil {
		return err
	}
	for _, worldMember := range worldMembersInWorld {
		if err = useCase.worldMemberRepo.Delete(worldMember); err != nil {
			return err
		}
	}

	err = useCase.worldService.DeleteWorld(worldId)
	if err != nil {
		return err
	}

	return nil
}
