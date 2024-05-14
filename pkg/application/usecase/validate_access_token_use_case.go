package usecase

import (
	"os"

	iam_service "github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/service"
	"github.com/google/uuid"
)

type ValidateAccessTokenUseCase struct {
	authService iam_service.AuthService
}

func NewValidateAccessTokenUseCase(authService iam_service.AuthService) ValidateAccessTokenUseCase {
	return ValidateAccessTokenUseCase{authService}
}

func ProvideValidateAccessTokenUseCase() ValidateAccessTokenUseCase {
	authService := iam_service.NewAuthService(os.Getenv("AUTH_SECRET"))

	return NewValidateAccessTokenUseCase(
		authService,
	)
}

func (useCase *ValidateAccessTokenUseCase) Execute(accessToken string) (userIdDto uuid.UUID, err error) {
	userId, err := useCase.authService.ValidateAccessToken(accessToken)
	if err != nil {
		return userIdDto, err
	}

	return userId.Uuid(), nil
}
