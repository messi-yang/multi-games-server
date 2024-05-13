package googleoauthinfrasrv

import (
	"context"
	"fmt"
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	google_api_oauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type GoogleOauthState struct {
	ClientRedirectPath string `json:"clientRedirectPath"`
}

type Service interface {
	GenerateOauthUrl(clientRedirectPath string) (oauthUrl string)
	UnmarshalOauthStateString(oauthStateString string) (oauthState GoogleOauthState, err error)
	GetUserEmailAddress(code string) (emailAddress globalcommonmodel.EmailAddress, err error)
}

type serve struct {
	config *oauth2.Config
}

var serveSingleton *serve

func NewService() Service {
	if serveSingleton != nil {
		return serveSingleton
	}
	return &serve{
		config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  fmt.Sprintf("%s/api/auth/oauth2/google/redirect", os.Getenv("SERVER_URL")),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (serve *serve) GenerateOauthUrl(clientRedirectPath string) (oauthUrl string) {
	stateInBytes := jsonutil.Marshal(GoogleOauthState{
		ClientRedirectPath: clientRedirectPath,
	})
	state := string(stateInBytes)
	return serve.config.AuthCodeURL(state)
}

func (serve *serve) UnmarshalOauthStateString(oauthStateString string) (oauthState GoogleOauthState, err error) {
	return jsonutil.Unmarshal[GoogleOauthState]([]byte(oauthStateString))
}

func (serve *serve) GetUserEmailAddress(code string) (emailAddress globalcommonmodel.EmailAddress, err error) {
	token, err := serve.config.Exchange(context.TODO(), code)
	if err != nil {
		return emailAddress, err
	}
	tokenSource := serve.config.TokenSource(context.TODO(), token)
	oauth2Service, err := google_api_oauth2.NewService(context.TODO(), option.WithTokenSource(tokenSource))
	if err != nil {
		return emailAddress, err
	}
	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return emailAddress, err
	}
	emailAddress, err = globalcommonmodel.NewEmailAddress(userInfo.Email)
	if err != nil {
		return emailAddress, err
	}
	return emailAddress, nil
}
