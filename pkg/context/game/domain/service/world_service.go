package service

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/commonutil"
)

type WorldService interface {
	CreateWorld(userId sharedkernelmodel.UserId, name string) (sharedkernelmodel.WorldId, error)
}

type worldServe struct {
	gamerRepo gamermodel.GamerRepo
	worldRepo worldmodel.WorldRepo
	unitRepo  unitmodel.UnitRepo
	itemRepo  itemmodel.ItemRepo
}

func NewWorldService(
	gamerRepo gamermodel.GamerRepo,
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
) WorldService {
	return &worldServe{
		gamerRepo: gamerRepo,
		worldRepo: worldRepo,
		unitRepo:  unitRepo,
		itemRepo:  itemRepo,
	}
}

func (worldServe *worldServe) CreateWorld(userId sharedkernelmodel.UserId, name string) (worldId sharedkernelmodel.WorldId, err error) {
	gamer, err := worldServe.gamerRepo.GetGamerByUserId(userId)
	if err != nil {
		return worldId, err
	}
	if err = gamer.AddWorldsCount(); err != nil {
		return worldId, err
	}
	if err = worldServe.gamerRepo.Update(gamer); err != nil {
		return worldId, err
	}

	newWorld := worldmodel.NewWorld(userId, name)
	worldId = newWorld.GetId()

	if err = worldServe.worldRepo.Add(newWorld); err != nil {
		return worldId, err
	}

	items, err := worldServe.itemRepo.GetAll()
	if err != nil {
		return worldId, err
	}

	if err = commonutil.RangeMatrix(100, 100, func(x int, z int) error {
		randomInt := rand.Intn(len(items) * 5)
		position := commonmodel.NewPosition(x-50, z-50)
		if randomInt < len(items) {
			newUnit := unitmodel.NewUnit(
				unitmodel.NewUnitId(worldId, position), worldId, position, items[randomInt].GetId(), commonmodel.NewDownDirection(),
			)
			if err = worldServe.unitRepo.Add(newUnit); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return worldId, err
	}

	return newWorld.GetId(), nil
}
