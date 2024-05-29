package worldaccountappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

type Service interface {
	HandleWorldCreatedDomainEvent(worldmodel.WorldCreated) error
	HandleWorldDeletedDomainEvent(worldmodel.WorldDeleted) error
}

type serve struct {
	worldAccountRepo worldaccountmodel.WorldAccountRepo
}

func NewService(
	worldAccountRepo worldaccountmodel.WorldAccountRepo,
) Service {
	return &serve{
		worldAccountRepo: worldAccountRepo,
	}
}

func (serve *serve) HandleWorldCreatedDomainEvent(worldCreated worldmodel.WorldCreated) error {
	worldAccount, err := serve.worldAccountRepo.GetWorldAccountOfUser(worldCreated.GetUserId())
	if err != nil {
		return err
	}
	worldAccount.AddWorldsCount()
	return serve.worldAccountRepo.Update(worldAccount)
}

func (serve *serve) HandleWorldDeletedDomainEvent(worldDeleted worldmodel.WorldDeleted) error {
	worldAccount, err := serve.worldAccountRepo.GetWorldAccountOfUser(worldDeleted.GetUserId())
	if err != nil {
		return err
	}
	worldAccount.SubtractWorldsCount()
	return serve.worldAccountRepo.Update(worldAccount)
}
