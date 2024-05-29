package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetMyWorldsUseCase struct {
	worldRepo worldmodel.WorldRepo
}

func NewGetMyWorldsUseCase(worldRepo worldmodel.WorldRepo) GetMyWorldsUseCase {
	return GetMyWorldsUseCase{worldRepo}
}

func ProvideGetMyWorldsUseCase(uow pguow.Uow) GetMyWorldsUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)

	return NewGetMyWorldsUseCase(worldRepo)
}

func (useCase *GetMyWorldsUseCase) Execute(useIdDto uuid.UUID) (worldDtos []dto.WorldDto, err error) {
	userId := globalcommonmodel.NewUserId(useIdDto)
	myWorlds, err := useCase.worldRepo.GetWorldsOfUser(userId)
	if err != nil {
		return worldDtos, err
	}

	myWorldDtos := lo.Map(myWorlds, func(world worldmodel.World, _ int) dto.WorldDto {
		return dto.NewWorldDto(world)
	})

	return myWorldDtos, nil
}
