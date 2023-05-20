package worlddomainsrv

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/commonutil"
)

type Service interface {
	CreateWorld(userId sharedkernelmodel.UserId, name string) (commonmodel.WorldId, error)
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
	return &serve{worldRepo: worldRepo, unitRepo: unitRepo, itemRepo: itemRepo, domainEventDispatcher: domainEventDispatcher}
}

func (serve *serve) CreateWorld(userId sharedkernelmodel.UserId, name string) (worldId commonmodel.WorldId, err error) {
	newWorld := worldmodel.NewWorld(userId, name)
	worldId = newWorld.GetId()

	if err = serve.worldRepo.Add(newWorld); err != nil {
		return worldId, err
	}
	if err = serve.domainEventDispatcher.Dispatch(&newWorld); err != nil {
		return worldId, err
	}

	items, err := serve.itemRepo.GetAll()
	if err != nil {
		return worldId, err
	}

	if err = commonutil.RangeMatrix(100, 100, func(x int, z int) error {
		randomInt := rand.Intn(40)
		position := commonmodel.NewPosition(x-50, z-50)
		if randomInt < 3 {
			newUnit := unitmodel.NewUnit(
				commonmodel.NewUnitId(worldId, position), worldId, position, items[randomInt].GetId(), commonmodel.NewDownDirection(),
			)
			if err = serve.unitRepo.Add(newUnit); err != nil {
				return err
			}
			if err = serve.domainEventDispatcher.Dispatch(&newUnit); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return worldId, err
	}

	return newWorld.GetId(), nil
}
