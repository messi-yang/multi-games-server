package worldaccountappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetWorldAccountOfUser(GetWorldAccountOfUserQuery) (dto.WorldAccountDto, error)
	CreateWorldAccount(CreateWorldAccountCommand) (worldAccountIdDto uuid.UUID, err error)
	QueryWorldAccounts(QueryWorldAccountsQuery) ([]dto.WorldAccountDto, error)
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
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worldAccount, err := serve.worldAccountRepo.GetWorldAccountOfUser(userId)
	if err != nil {
		return worldAccountDto, err
	}

	return dto.NewWorldAccountDto(worldAccount), nil
}

func (serve *serve) CreateWorldAccount(command CreateWorldAccountCommand) (worldAccountIdDto uuid.UUID, err error) {
	userId := sharedkernelmodel.NewUserId(command.UserId)
	_, worldAccountFound, err := serve.worldAccountRepo.FindWorldAccountByUserId(userId)
	if err != nil {
		return worldAccountIdDto, err
	}
	if worldAccountFound {
		return worldAccountIdDto, fmt.Errorf("already has a worldAccount with userId of %s", userId.Uuid().String())
	}
	newWorldAccount := worldaccountmodel.NewWorldAccount(userId, 0, 1)
	if err = serve.worldAccountRepo.Add(newWorldAccount); err != nil {
		return worldAccountIdDto, err
	}
	return newWorldAccount.GetId().Uuid(), nil
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
