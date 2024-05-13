package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/service/googleoauthinfrasrv"
)

type GenerateGoogleOauthUrlUseCase struct {
	googleOauthInfraService googleoauthinfrasrv.Service
}

func NewGenerateGoogleOauthUrlUseCase(googleOauthInfraService googleoauthinfrasrv.Service) GenerateGoogleOauthUrlUseCase {
	return GenerateGoogleOauthUrlUseCase{googleOauthInfraService}
}

func ProvideGenerateGoogleOauthUrlUseCase() GenerateGoogleOauthUrlUseCase {
	googleOauthInfraService := googleoauthinfrasrv.NewService()

	return NewGenerateGoogleOauthUrlUseCase(
		googleOauthInfraService,
	)
}

func (useCase *GenerateGoogleOauthUrlUseCase) Execute(clientRedirectPath string) (googleOauthUrl string) {
	return useCase.googleOauthInfraService.GenerateOauthUrl(clientRedirectPath)
}
