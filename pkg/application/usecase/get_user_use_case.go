package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type GetUserUseCase struct {
	userRepo usermodel.UserRepo
}

func NewGetUserUseCase(userRepo usermodel.UserRepo) GetUserUseCase {
	return GetUserUseCase{userRepo}
}

func ProvideGetUserUseCase(uow pguow.Uow) GetUserUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	userRepo := iam_pgrepo.NewUserRepo(uow, domainEventDispatcher)

	return NewGetUserUseCase(userRepo)
}

func (useCase *GetUserUseCase) Execute(userIdDto uuid.UUID) (userDto dto.UserDto, err error) {
	userId := globalcommonmodel.NewUserId(userIdDto)

	user, err := useCase.userRepo.Get(userId)
	if err != nil {
		return userDto, err
	}

	return dto.NewUserDto(user), nil
}
