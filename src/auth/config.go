package auth

import (
	"fmt"

	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var Auth = GetAuthenticator()

func GetAuthenticator() *oauth2.Config {

	endpoints := oauth2.Endpoint{
		AuthURL:  fmt.Sprintf("https://%s/oauth/authorize", config.C.Oauth.Domain),
		TokenURL: fmt.Sprintf("https://%s/oauth/token", config.C.Oauth.Domain),
	}

	conf := oauth2.Config{
		ClientID:     config.C.Oauth.ClientID,
		ClientSecret: config.C.Oauth.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s/callback", config.C.Server.Host),
		Endpoint:     endpoints,
		Scopes:       []string{oidc.ScopeOpenID},
	}

	return &conf
}
