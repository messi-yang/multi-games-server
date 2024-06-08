package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/redisrepo"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetWorldInformationUseCase struct {
	worldRepo  worldmodel.WorldRepo
	playerRepo playermodel.PlayerRepo
	unitRepo   unitmodel.UnitRepo
}

func NewGetWorldInformationUseCase(worldRepo worldmodel.WorldRepo, playerRepo playermodel.PlayerRepo, unitRepo unitmodel.UnitRepo) GetWorldInformationUseCase {
	return GetWorldInformationUseCase{worldRepo, playerRepo, unitRepo}
}

func ProvideGetWorldInformationUseCase(uow pguow.Uow) GetWorldInformationUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	worldRepo := pgrepo.NewWorldRepo(uow, domainEventDispatcher)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)

	return NewGetWorldInformationUseCase(worldRepo, playerRepo, unitRepo)
}

func (useCase *GetWorldInformationUseCase) Execute(worldIdDto uuid.UUID) (
	worldDto dto.WorldDto, unitDtos []dto.UnitDto, playerDtos []dto.PlayerDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)

	world, err := useCase.worldRepo.Get(worldId)
	if err != nil {
		return worldDto, unitDtos, playerDtos, err
	}
	worldDto = dto.NewWorldDto(world)

	units, err := useCase.unitRepo.GetUnitsOfWorld(worldId)
	if err != nil {
		return worldDto, unitDtos, playerDtos, err
	}
	unitDtos = lo.Map(units, func(unit unitmodel.Unit, _ int) dto.UnitDto {
		return dto.NewUnitDto(unit)
	})

	players, err := useCase.playerRepo.GetPlayersOfWorld(worldId)
	if err != nil {
		return worldDto, unitDtos, playerDtos, err
	}
	playerDtos = lo.Map(players, func(_player playermodel.Player, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(_player)
	})

	return worldDto, unitDtos, playerDtos, nil
}
