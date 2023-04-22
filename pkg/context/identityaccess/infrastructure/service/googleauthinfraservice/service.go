package googleauthinfraservice

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	google_api_oauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type Service interface {
	GenerateAuthUrl(command GenerateAuthUrlCommand) (url string)
	GetUserEmailAddress(query GetUserEmailAddressQuery) (emailAddress string, err error)
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

func (serve *serve) GenerateAuthUrl(command GenerateAuthUrlCommand) string {
	return serve.config.AuthCodeURL("state")
}

func (serve *serve) GetUserEmailAddress(query GetUserEmailAddressQuery) (emailAddress string, err error) {
	token, err := serve.config.Exchange(context.TODO(), query.Code)
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
	return userInfo.Email, nil
}
