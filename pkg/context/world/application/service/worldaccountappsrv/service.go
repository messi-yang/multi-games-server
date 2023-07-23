package worldaccountappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/samber/lo"
)

type Service interface {
	GetWorldAccountOfUser(GetWorldAccountOfUserQuery) (dto.WorldAccountDto, error)
	CreateWorldAccount(CreateWorldAccountCommand) error
	QueryWorldAccounts(QueryWorldAccountsQuery) ([]dto.WorldAccountDto, error)
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

func (serve *serve) GetWorldAccountOfUser(query GetWorldAccountOfUserQuery) (worldAccountDto dto.WorldAccountDto, err error) {
	userId := globalcommonmodel.NewUserId(query.UserId)
	worldAccount, err := serve.worldAccountRepo.GetWorldAccountOfUser(userId)
	if err != nil {
		return worldAccountDto, err
	}

	return dto.NewWorldAccountDto(worldAccount), nil
}

func (serve *serve) CreateWorldAccount(command CreateWorldAccountCommand) (err error) {
	userId := globalcommonmodel.NewUserId(command.UserId)
	worldAccount, err := serve.worldAccountRepo.GetWorldAccountByUserId(userId)
	if err != nil {
		return err
	}
	if worldAccount != nil {
		return fmt.Errorf("already has a worldAccount with userId of %s", userId.Uuid().String())
	}
	newWorldAccount := worldaccountmodel.NewWorldAccount(userId)
	return serve.worldAccountRepo.Add(newWorldAccount)
}

func (serve *serve) QueryWorldAccounts(query QueryWorldAccountsQuery) (itemDtos []dto.WorldAccountDto, err error) {
	worldAccounts, err := serve.worldAccountRepo.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(worldAccounts, func(worldAccount worldaccountmodel.WorldAccount, _ int) dto.WorldAccountDto {
		return dto.NewWorldAccountDto(worldAccount)
	}), nil
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
