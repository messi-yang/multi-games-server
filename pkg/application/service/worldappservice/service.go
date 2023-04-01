package worldappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/commonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type Service interface {
	GetWorld(GetWorldQuery) (worldmodel.WorldAgg, error)
	GetWorlds(GetWorldsQuery) ([]worldmodel.WorldAgg, error)
	CreateWorld(CreateWorldCommand) (worldmodel.WorldIdVo, error)
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

func (serve *serve) GetWorld(query GetWorldQuery) (worldmodel.WorldAgg, error) {
	return serve.worldRepository.Get(query.WorldId)
}

func (serve *serve) GetWorlds(query GetWorldsQuery) (worlds []worldmodel.WorldAgg, err error) {
	return serve.worldRepository.GetAll()
}

func (serve *serve) CreateWorld(command CreateWorldCommand) (newWorldId worldmodel.WorldIdVo, err error) {
	worldId := worldmodel.NewWorldIdVo(uuid.New())
	newWorld := worldmodel.NewWorldAgg(worldId, command.UserId)

	err = serve.worldRepository.Add(newWorld)
	if err != nil {
		return newWorldId, err
	}

	items, err := serve.itemRepository.GetAll()
	if err != nil {
		return newWorldId, err
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

	return newWorld.GetId(), nil
}
