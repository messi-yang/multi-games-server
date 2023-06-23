package service

import (
	"math/rand"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

type WorldService interface {
	CreateWorld(userId sharedkernelmodel.UserId, name string) (sharedkernelmodel.WorldId, error)
}

type worldServe struct {
	worldAccountRepo worldaccountmodel.WorldAccountRepo
	worldRepo        worldmodel.WorldRepo
	unitRepo         unitmodel.UnitRepo
	itemRepo         itemmodel.ItemRepo
}

func NewWorldService(
	worldAccountRepo worldaccountmodel.WorldAccountRepo,
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
) WorldService {
	return &worldServe{
		worldAccountRepo: worldAccountRepo,
		worldRepo:        worldRepo,
		unitRepo:         unitRepo,
		itemRepo:         itemRepo,
	}
}

func (worldServe *worldServe) CreateWorld(userId sharedkernelmodel.UserId, name string) (worldId sharedkernelmodel.WorldId, err error) {
	worldAccount, err := worldServe.worldAccountRepo.GetWorldAccountOfUser(userId)
	if err != nil {
		return worldId, err
	}
	if err = worldAccount.AddWorldsCount(); err != nil {
		return worldId, err
	}
	if err = worldServe.worldAccountRepo.Update(worldAccount); err != nil {
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
