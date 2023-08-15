package auth

import (
	"context"
	"fmt"

	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/samber/lo"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

var Auth = GetAuthenticator()

func GetAuthenticator() *Authenticator {
	provider := lo.Must(oidc.NewProvider(
		context.Background(),
		"https://"+config.C.Oauth.Domain+"/",
	))

	endpoints := oauth2.Endpoint{
		AuthURL:  fmt.Sprintf("https://%s/oauth/authorize", config.C.Oauth.Domain),
		TokenURL: fmt.Sprintf("https://%s/oauth/token", config.C.Oauth.Domain),
	}

	conf := oauth2.Config{
		ClientID:     config.C.Oauth.ClientID,
		ClientSecret: config.C.Oauth.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s/callback", config.C.Server.Host),
		Endpoint:     endpoints,
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}
}
