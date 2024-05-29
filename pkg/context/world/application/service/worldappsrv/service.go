package worldappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/samber/lo"
)

var (
	ErrNotPermitted = fmt.Errorf("not permitted to perform this action")
)

type Service interface {
	GetWorld(GetWorldQuery) (dto.WorldDto, error)
	GetMyWorlds(GetMyWorldsQuery) ([]dto.WorldDto, error)
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
	worldId := globalcommonmodel.NewWorldId(query.WorldId)
	world, err := serve.worldRepo.Get(worldId)
	if err != nil {
		return worldDto, err
	}
	return dto.NewWorldDto(world), nil
}

func (serve *serve) GetMyWorlds(query GetMyWorldsQuery) (worldDtos []dto.WorldDto, err error) {
	userId := globalcommonmodel.NewUserId(query.UserId)
	worlds, err := serve.worldRepo.GetWorldsOfUser(userId)
	if err != nil {
		return worldDtos, err
	}

	return lo.Map(worlds, func(world worldmodel.World, _ int) dto.WorldDto {
		return dto.NewWorldDto(world)
	}), nil
}
