package worldappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
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
	QueryWorlds()
	CreateWorld(userIdDto uuid.UUID)
}

type serve struct {
	worldRepository worldmodel.Repository
	unitRepository  unitmodel.Repository
	itemRepository  itemmodel.Repository
	presenter       Presenter
}

func NewService(worldRepository worldmodel.Repository, unitRepository unitmodel.Repository, itemRepository itemmodel.Repository, presenter Presenter) Service {
	return &serve{
		worldRepository: worldRepository,
		unitRepository:  unitRepository,
		itemRepository:  itemRepository,
		presenter:       presenter,
	}
}

func (serve *serve) QueryWorlds() {
	worlds, err := serve.worldRepository.GetAll()
	if err != nil {
		serve.presenter.OnError(err)
		return
	}
	worldDtos := lo.Map(worlds, func(world worldmodel.WorldAgg, _ int) dto.WorldAggDto {
		return dto.NewWorldAggDto(world)
	})
	responseDto := QueryWorldsResponseDto(worldDtos)
	serve.presenter.OnSuccess(responseDto)
}

func (serve *serve) CreateWorld(userIdDto uuid.UUID) {
	userId := usermodel.NewUserIdVo(userIdDto)

	worldId := worldmodel.NewWorldIdVo(uuid.New())
	newWorld := worldmodel.NewWorldAgg(worldId, userId)

	err := serve.worldRepository.Add(newWorld)
	if err != nil {
		serve.presenter.OnError(err)
		return
	}

	items, err := serve.itemRepository.GetAll()
	if err != nil {
		serve.presenter.OnError(err)
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

	newWorldDto := dto.NewWorldAggDto(newWorld)
	serve.presenter.OnSuccess(CreateWorldResponseDto(newWorldDto))
}
