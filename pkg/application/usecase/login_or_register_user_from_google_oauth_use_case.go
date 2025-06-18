package usecase

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	iam_service "github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/service"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/service/googleoauthinfrasrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldaccountmodel"
	world_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
)

type LoginOrRegisterUserFromGoogleOauthUseCase struct {
	userRepo                usermodel.UserRepo
	worldAccountRepo        worldaccountmodel.WorldAccountRepo
	userService             iam_service.UserService
	authService             iam_service.AuthService
	googleOauthInfraService googleoauthinfrasrv.Service
}

func NewLoginOrRegisterUserFromGoogleOauthUseCase(userRepo usermodel.UserRepo, worldAccountRepo worldaccountmodel.WorldAccountRepo,
	userService iam_service.UserService, authService iam_service.AuthService, googleOauthInfraService googleoauthinfrasrv.Service,
) LoginOrRegisterUserFromGoogleOauthUseCase {
	return LoginOrRegisterUserFromGoogleOauthUseCase{userRepo, worldAccountRepo, userService, authService, googleOauthInfraService}
}

func ProvideLoginOrRegisterUserFromGoogleOauthUseCase(uow pguow.Uow) LoginOrRegisterUserFromGoogleOauthUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	userRepo := iam_pgrepo.NewUserRepo(uow, domainEventDispatcher)
	worldAccountRepo := world_pgrepo.NewWorldAccountRepo(uow, domainEventDispatcher)
	userService := iam_service.NewUserService(userRepo)
	authService := iam_service.NewAuthService(os.Getenv("AUTH_SECRET"))
	googleOauthInfraService := googleoauthinfrasrv.NewService()

	return NewLoginOrRegisterUserFromGoogleOauthUseCase(
		userRepo,
		worldAccountRepo,
		userService,
		authService,
		googleOauthInfraService,
	)
}

func (useCase *LoginOrRegisterUserFromGoogleOauthUseCase) Execute(code string, oauthStateString string) (redirectPath string, err error) {
	clientUrl := os.Getenv("CLIENT_URL")

	fmt.Println("code", code)
	fmt.Println("oauthStateString", oauthStateString)

	googleUserEmailAddress, err := useCase.googleOauthInfraService.GetUserEmailAddress(code)
	if err != nil {
		fmt.Println(code, err)
		return
	}

	googleOauthState, err := useCase.googleOauthInfraService.UnmarshalOauthStateString(oauthStateString)
	if err != nil {
		return
	}

	existingUser, err := useCase.userRepo.GetUserByEmailAddress(googleUserEmailAddress)
	if err != nil {
		return redirectPath, err
	}
	if existingUser != nil {
		accessToken, err := useCase.authService.GenerateAccessToken(existingUser.GetId())
		if err != nil {
			return redirectPath, err
		}

		return fmt.Sprintf(
			"%s/auth/sign-in-success/?access_token=%v&client_redirect_path=%v",
			clientUrl,
			accessToken,
			googleOauthState.ClientRedirectPath,
		), nil
	}

	newUserId, err := useCase.userService.CreateUser(googleUserEmailAddress)
	if err != nil {
		fmt.Println(err)
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

	return fmt.Sprintf(
		"%s/auth/sign-in-success/?access_token=%v&client_redirect_path=%v",
		clientUrl,
		accessToken,
		googleOauthState.ClientRedirectPath,
	), nil
}
