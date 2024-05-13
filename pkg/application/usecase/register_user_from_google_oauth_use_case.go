package usecase

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	iam_service "github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/service"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/service/googleoauthinfrasrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	world_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

type RegisterUserFromGoogleOauthUseCase struct {
	userRepo                usermodel.UserRepo
	worldAccountRepo        worldaccountmodel.WorldAccountRepo
	userService             iam_service.UserService
	authService             iam_service.AuthService
	googleOauthInfraService googleoauthinfrasrv.Service
}

func NewRegisterUserFromGoogleOauthUseCase(userRepo usermodel.UserRepo, worldAccountRepo worldaccountmodel.WorldAccountRepo,
	userService iam_service.UserService, authService iam_service.AuthService, googleOauthInfraService googleoauthinfrasrv.Service,
) RegisterUserFromGoogleOauthUseCase {
	return RegisterUserFromGoogleOauthUseCase{userRepo, worldAccountRepo, userService, authService, googleOauthInfraService}
}

func ProvideRegisterUserFromGoogleOauthUseCase(uow pguow.Uow) RegisterUserFromGoogleOauthUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	userRepo := iam_pgrepo.NewUserRepo(uow, domainEventDispatcher)
	worldAccountRepo := world_pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	userService := iam_service.NewUserService(userRepo)
	authService := iam_service.NewAuthService(os.Getenv("AUTH_SECRET"))
	googleOauthInfraService := googleoauthinfrasrv.NewService()

	return NewRegisterUserFromGoogleOauthUseCase(
		userRepo,
		worldAccountRepo,
		userService,
		authService,
		googleOauthInfraService,
	)
}

func (useCase *RegisterUserFromGoogleOauthUseCase) Execute(code string, oauthStateString string) (redirectPath string, err error) {
	googleUserEmailAddress, err := useCase.googleOauthInfraService.GetUserEmailAddress(code)
	if err != nil {
		return
	}

	googleOauthState, err := useCase.googleOauthInfraService.UnmarshalOauthStateString(oauthStateString)
	if err != nil {
		return
	}

	newUserId, err := useCase.userService.CreateUser(googleUserEmailAddress)
	if err != nil {
		return redirectPath, err
	}

	// TODO - handle this side effects by using integration events
	newWorldAccount := worldaccountmodel.NewWorldAccount(newUserId)
	err = useCase.worldAccountRepo.Add(newWorldAccount)
	if err != nil {
		return redirectPath, err
	}

	accessToken, err := useCase.authService.GenerateAccessToken(newUserId)
	if err != nil {
		return redirectPath, err
	}

	clientUrl := os.Getenv("CLIENT_URL")
	return fmt.Sprintf(
		"%s/auth/sign-in-success/?access_token=%v&client_redirect_path=%v",
		clientUrl,
		accessToken,
		googleOauthState.ClientRedirectPath,
	), nil
}
