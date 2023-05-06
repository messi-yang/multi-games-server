package worldappsrv

import (
	"fmt"
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetWorld(GetWorldQuery) (dto.WorldDto, error)
	QueryWorlds(QueryWorldsQuery) ([]dto.WorldDto, error)
	CreateWorld(CreateWorldCommand) (uuid.UUID, error)
	UpdateWorld(UpdateWorldCommand) error
}

type serve struct {
	worldRepo             worldmodel.Repo
	unitRepo              unitmodel.Repo
	itemRepo              itemmodel.Repo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewService(
	worldRepo worldmodel.Repo,
	unitRepo unitmodel.Repo,
	itemRepo itemmodel.Repo,
	domainEventDispatcher domain.DomainEventDispatcher,
) Service {
	return &serve{
		worldRepo:             worldRepo,
		unitRepo:              unitRepo,
		itemRepo:              itemRepo,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (serve *serve) GetWorld(query GetWorldQuery) (worldDto dto.WorldDto, err error) {
	worldId := commonmodel.NewWorldId(query.WorldId)
	world, err := serve.worldRepo.Get(worldId)
	if err != nil {
		return worldDto, err
	}
	return dto.NewWorldDto(world), nil
}

func (serve *serve) QueryWorlds(query QueryWorldsQuery) (worldDtos []dto.WorldDto, err error) {
	worlds, err := serve.worldRepo.GetAll()
	if err != nil {
		return worldDtos, err
	}

	return lo.Map(worlds, func(world worldmodel.World, _ int) dto.WorldDto {
		return dto.NewWorldDto(world)
	}), nil
}
func (serve *serve) CreateWorld(command CreateWorldCommand) (newWorldIdDto uuid.UUID, err error) {
	worldId := commonmodel.NewWorldId(uuid.New())
	gamerId := commonmodel.NewGamerId(command.GamerId)
	newWorld := worldmodel.NewWorld(worldId, gamerId, "Hello World")

	if err = serve.worldRepo.Add(newWorld); err != nil {
		return newWorldIdDto, err
	}
	if err = serve.domainEventDispatcher.Dispatch(&newWorld); err != nil {
		return newWorldIdDto, err
	}

	items, err := serve.itemRepo.GetAll()
	if err != nil {
		return newWorldIdDto, err
	}

	commonutil.RangeMatrix(100, 100, func(x int, z int) error {
		randomInt := rand.Intn(40)
		position := commonmodel.NewPosition(x-50, z-50)
		if randomInt < 3 {
			newUnit := unitmodel.NewUnit(worldId, position, items[randomInt].GetId(), commonmodel.NewDownDirection())
			if err = serve.unitRepo.Add(newUnit); err != nil {
				return err
			}
			if err = serve.domainEventDispatcher.Dispatch(&newUnit); err != nil {
				return err
			}
		}
		return nil
	})

	return newWorld.GetId().Uuid(), nil
}

func (serve *serve) UpdateWorld(command UpdateWorldCommand) error {
	worldId := commonmodel.NewWorldId(command.WorldId)
	gamerId := commonmodel.NewGamerId(command.GamerId)
	world, err := serve.worldRepo.Get(worldId)
	if err != nil {
		return err
	}
	if !world.GetGamerId().IsEqual(gamerId) {
		return fmt.Errorf("the world with id of %s do not belong to gamer with id of %s", worldId.Uuid().String(), gamerId.Uuid().String())
	}
	world.ChangeName(command.Name)
	if err = serve.worldRepo.Update(world); err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&world)
}
