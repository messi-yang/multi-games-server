package service

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

var (
	ErrWorldsCountReachLimit = errors.New("worlds count has reached the limit")
	ErrDeleteNotWorking      = errors.New("world delete is not working now")
)

type WorldService interface {
	CreateWorld(userId globalcommonmodel.UserId, name string) (worldmodel.World, error)
	DeleteWorld(worldId globalcommonmodel.WorldId) error
}

type worldServe struct {
	worldAccountRepo worldaccountmodel.WorldAccountRepo
	worldRepo        worldmodel.WorldRepo
}

func NewWorldService(
	worldAccountRepo worldaccountmodel.WorldAccountRepo,
	worldRepo worldmodel.WorldRepo,
) WorldService {
	return &worldServe{
		worldAccountRepo: worldAccountRepo,
		worldRepo:        worldRepo,
	}
}

func (worldServe *worldServe) CreateWorld(userId globalcommonmodel.UserId, name string) (newWorld worldmodel.World, err error) {
	worldAccount, err := worldServe.worldAccountRepo.GetWorldAccountOfUser(userId)
	if err != nil {
		return newWorld, err
	}
	if !worldAccount.CanAddNewWorld() {
		return newWorld, ErrWorldsCountReachLimit
	}

	newWorld = worldmodel.NewWorld(userId, name)

	if err = worldServe.worldRepo.Add(newWorld); err != nil {
		return newWorld, err
	}

	return newWorld, nil
}

func (worldServe *worldServe) DeleteWorld(worldId globalcommonmodel.WorldId) error {
	// TODO - We need to figure out if it's a good idea to totally delete a world from database, maybe we can just flag it as archived
	return ErrDeleteNotWorking
}
