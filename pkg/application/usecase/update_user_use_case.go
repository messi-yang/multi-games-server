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

type UpdateUserUseCase struct {
	userRepo usermodel.UserRepo
}

func NewUpdateUserUseCase(userRepo usermodel.UserRepo) UpdateUserUseCase {
	return UpdateUserUseCase{userRepo}
}

func ProvideUpdateUserUseCase(uow pguow.Uow) UpdateUserUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	userRepo := iam_pgrepo.NewUserRepo(uow, domainEventDispatcher)

	return NewUpdateUserUseCase(userRepo)
}

func (useCase *UpdateUserUseCase) Execute(userIdDto uuid.UUID, usernameDto string, friendlyNameDto string) (
	updatedUserDto dto.UserDto, err error,
) {
	userId := globalcommonmodel.NewUserId(userIdDto)
	username, err := globalcommonmodel.NewUsername(usernameDto)
	if err != nil {
		return updatedUserDto, err
	}

	friendlyName, err := usermodel.NewFriendlyName(friendlyNameDto)
	if err != nil {
		return updatedUserDto, err
	}

	user, err := useCase.userRepo.Get(userId)
	if err != nil {
		return updatedUserDto, err
	}

	user.UpdateUsername(username)
	user.UpdateFriendlyName(friendlyName)

	if err = useCase.userRepo.Update(user); err != nil {
		return updatedUserDto, err
	}

	updatedUser, err := useCase.userRepo.Get(userId)
	if err != nil {
		return updatedUserDto, err
	}

	return dto.NewUserDto(updatedUser), nil
}
