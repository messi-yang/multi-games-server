package worldappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/commonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type Service interface {
	QueryWorlds(worldsTransformer func([]worldmodel.WorldAgg), errorTransformer func(error))
	CreateWorld(userIdDto uuid.UUID, worldTransformer func(worldmodel.WorldAgg), errorTransformer func(error))
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

func (serve *serve) QueryWorlds(worldsTransformer func([]worldmodel.WorldAgg), errorTransformer func(error)) {
	worlds, err := serve.worldRepository.GetAll()
	if err != nil {
		errorTransformer(err)
		return
	}
	worldsTransformer(worlds)
}

func (serve *serve) CreateWorld(userIdDto uuid.UUID, worldTransformer func(worldmodel.WorldAgg), errorTransformer func(error)) {
	userId := usermodel.NewUserIdVo(userIdDto)

	worldId := worldmodel.NewWorldIdVo(uuid.New())
	newWorld := worldmodel.NewWorldAgg(worldId, userId)

	err := serve.worldRepository.Add(newWorld)
	if err != nil {
		errorTransformer(err)
		return
	}

	items, err := serve.itemRepository.GetAll()
	if err != nil {
		errorTransformer(err)
		return
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

	worldTransformer(newWorld)
}
