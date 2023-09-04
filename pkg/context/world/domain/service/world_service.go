package service

import (
	"errors"
	"math/rand"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

var (
	ErrWorldsCountReachLimit = errors.New("worlds count has reached the limit")
)

type WorldService interface {
	CreateWorld(userId globalcommonmodel.UserId, name string) (globalcommonmodel.WorldId, error)
	DeleteWorld(worldId globalcommonmodel.WorldId) error
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

func (worldServe *worldServe) CreateWorld(userId globalcommonmodel.UserId, name string) (worldId globalcommonmodel.WorldId, err error) {
	worldAccount, err := worldServe.worldAccountRepo.GetWorldAccountOfUser(userId)
	if err != nil {
		return worldId, err
	}
	if !worldAccount.CanAddNewWorld() {
		return worldId, ErrWorldsCountReachLimit
	}

	worldBound, err := worldcommonmodel.NewBound(
		worldcommonmodel.NewPosition(-50, -50),
		worldcommonmodel.NewPosition(50, 50),
	)
	if err != nil {
		return worldId, err
	}

	newWorld := worldmodel.NewWorld(userId, name, worldBound)
	worldId = newWorld.GetId()

	if err = worldServe.worldRepo.Add(newWorld); err != nil {
		return worldId, err
	}

	itemsForStaticUnitType, err := worldServe.itemRepo.GetItemsOfCompatibleUnitType(worldcommonmodel.NewStaticUnitType())
	if err != nil {
		return worldId, err
	}

	if err = commonutil.RangeMatrix(100, 100, func(x int, z int) error {
		randomInt := rand.Intn(len(itemsForStaticUnitType) * 5)
		position := worldcommonmodel.NewPosition(x-50, z-50)
		if randomInt < len(itemsForStaticUnitType) {
			newUnit := unitmodel.NewUnit(
				worldId,
				position,
				itemsForStaticUnitType[randomInt].GetId(),
				worldcommonmodel.NewDownDirection(),
				itemsForStaticUnitType[randomInt].GetCompatibleUnitType(),
				nil,
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

func (worldServe *worldServe) DeleteWorld(worldId globalcommonmodel.WorldId) error {
	unitsInWorld, err := worldServe.unitRepo.GetUnitsOfWorld(worldId)
	if err != nil {
		return err
	}

	for _, unit := range unitsInWorld {
		if err = worldServe.unitRepo.Delete(unit); err != nil {
			return err
		}
	}

	world, err := worldServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}
	world.Delete()
	return worldServe.worldRepo.Delete(world)
}
