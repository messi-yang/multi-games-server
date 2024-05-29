package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type GetWorldUseCase struct {
	worldRepo worldmodel.WorldRepo
}

func NewGetWorldUseCase(worldRepo worldmodel.WorldRepo) GetWorldUseCase {
	return GetWorldUseCase{worldRepo}
}

func ProvideGetWorldUseCase(uow pguow.Uow) GetWorldUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)

	return NewGetWorldUseCase(worldRepo)
}

func (useCase *GetWorldUseCase) Execute(worldIdDto uuid.UUID) (worldDto dto.WorldDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)
	world, err := useCase.worldRepo.Get(worldId)
	if err != nil {
		return worldDto, err
	}

	return dto.NewWorldDto(world), nil
}
