package worldappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetWorld(GetWorldQuery) (dto.WorldDto, error)
	GetMyWorlds(GetMyWorldsQuery) ([]dto.WorldDto, error)
	QueryWorlds(QueryWorldsQuery) ([]dto.WorldDto, error)
	CreateWorld(CreateWorldCommand) (uuid.UUID, error)
	UpdateWorld(UpdateWorldCommand) error
}

type serve struct {
	worldRepo          worldmodel.WorldRepo
	worldDomainService service.WorldService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	worldDomainService service.WorldService,
) Service {
	return &serve{
		worldRepo:          worldRepo,
		worldDomainService: worldDomainService,
	}
}

func (serve *serve) GetWorld(query GetWorldQuery) (worldDto dto.WorldDto, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	world, err := serve.worldRepo.Get(worldId)
	if err != nil {
		return worldDto, err
	}
	return dto.NewWorldDto(world), nil
}

func (serve *serve) GetMyWorlds(query GetMyWorldsQuery) (worldDtos []dto.WorldDto, err error) {
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worlds, err := serve.worldRepo.GetWorldsOfUser(userId)
	if err != nil {
		return worldDtos, err
	}

	return lo.Map(worlds, func(world worldmodel.World, _ int) dto.WorldDto {
		return dto.NewWorldDto(world)
	}), nil
}

func (serve *serve) QueryWorlds(query QueryWorldsQuery) (worldDtos []dto.WorldDto, err error) {
	worlds, err := serve.worldRepo.Query(query.Limit, query.Offset)
	if err != nil {
		return worldDtos, err
	}

	return lo.Map(worlds, func(world worldmodel.World, _ int) dto.WorldDto {
		return dto.NewWorldDto(world)
	}), nil
}

func (serve *serve) CreateWorld(command CreateWorldCommand) (newWorldIdDto uuid.UUID, err error) {
	userId := sharedkernelmodel.NewUserId(command.UserId)
	newWorldId, err := serve.worldDomainService.CreateWorld(userId, command.Name)
	if err != nil {
		return newWorldIdDto, err
	}

	return newWorldId.Uuid(), nil
}

func (serve *serve) UpdateWorld(command UpdateWorldCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	userId := sharedkernelmodel.NewUserId(command.UserId)
	world, err := serve.worldRepo.Get(worldId)
	if err != nil {
		return err
	}
	if !world.GetUserId().IsEqual(userId) {
		return fmt.Errorf("the world with id of %s do not belong to gamer with id of %s", worldId.Uuid().String(), userId.Uuid().String())
	}
	world.ChangeName(command.Name)
	return serve.worldRepo.Update(world)
}
