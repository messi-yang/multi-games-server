package worldappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

var (
	ErrNotPermitted = fmt.Errorf("not permitted to perform this action")
)

type Service interface {
	GetWorld(GetWorldQuery) (dto.WorldDto, error)
	GetMyWorlds(GetMyWorldsQuery) ([]dto.WorldDto, error)
	QueryWorlds(QueryWorldsQuery) ([]dto.WorldDto, error)
	CreateWorld(CreateWorldCommand) (uuid.UUID, error)
	UpdateWorld(UpdateWorldCommand) error
	DeleteWorld(DeleteWorldCommand) error
}

type serve struct {
	worldRepo    worldmodel.WorldRepo
	worldService service.WorldService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	worldService service.WorldService,
) Service {
	return &serve{
		worldRepo:    worldRepo,
		worldService: worldService,
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
	newWorldId, err := serve.worldService.CreateWorld(userId, command.Name)
	if err != nil {
		return newWorldIdDto, err
	}

	return newWorldId.Uuid(), nil
}

func (serve *serve) UpdateWorld(command UpdateWorldCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	world, err := serve.worldRepo.Get(worldId)
	if err != nil {
		return err
	}
	world.ChangeName(command.Name)
	return serve.worldRepo.Update(world)
}

func (serve *serve) DeleteWorld(command DeleteWorldCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	return serve.worldService.DeleteWorld(worldId)
}
