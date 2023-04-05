package worldappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/commonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetWorld(GetWorldQuery) (jsondto.WorldAggDto, error)
	GetWorlds(GetWorldsQuery) ([]jsondto.WorldAggDto, error)
	CreateWorld(CreateWorldCommand) (uuid.UUID, error)
}

type serve struct {
	worldRepository worldmodel.Repository
	unitRepository  unitmodel.Repository
	itemRepository  itemmodel.Repository
}

func NewService(worldRepository worldmodel.Repository, unitRepository unitmodel.Repository, itemRepository itemmodel.Repository) Service {
	return &serve{
		worldRepository: worldRepository,
		unitRepository:  unitRepository,
		itemRepository:  itemRepository,
	}
}

func (serve *serve) GetWorld(query GetWorldQuery) (worldDto jsondto.WorldAggDto, err error) {
	worldId := worldmodel.NewWorldIdVo(query.WorldId)
	world, err := serve.worldRepository.Get(worldId)
	if err != nil {
		return worldDto, err
	}
	return jsondto.NewWorldAggDto(world), nil
}

func (serve *serve) GetWorlds(query GetWorldsQuery) (worldDtos []jsondto.WorldAggDto, err error) {
	worlds, err := serve.worldRepository.GetAll()
	if err != nil {
		return worldDtos, err
	}

	return lo.Map(worlds, func(world worldmodel.WorldAgg, _ int) jsondto.WorldAggDto {
		return jsondto.NewWorldAggDto(world)
	}), nil
}

func (serve *serve) CreateWorld(command CreateWorldCommand) (newWorldIdDto uuid.UUID, err error) {
	worldId := worldmodel.NewWorldIdVo(uuid.New())
	userId := usermodel.NewUserIdVo(command.UserId)
	newWorld := worldmodel.NewWorldAgg(worldId, userId)

	err = serve.worldRepository.Add(newWorld)
	if err != nil {
		return newWorldIdDto, err
	}

	items, err := serve.itemRepository.GetAll()
	if err != nil {
		return newWorldIdDto, err
	}

	commonutil.RangeMatrix(100, 100, func(x int, z int) error {
		randomInt := rand.Intn(40)
		position := commonmodel.NewPositionVo(x-50, z-50)
		if randomInt < 3 {
			newUnit := unitmodel.NewUnitAgg(worldId, position, items[randomInt].GetId(), commonmodel.NewDownDirectionVo())
			err = serve.unitRepository.Add(newUnit)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return newWorld.GetId().Uuid(), nil
}
