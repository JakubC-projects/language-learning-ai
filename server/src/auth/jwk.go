package auth

import (
	"context"
	"fmt"

	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/lestrrat-go/jwx/jwk"
)

var jwkSet jwk.Set

func init() {
	keyURL := fmt.Sprintf("https://%s/.well-known/jwks.json", config.C.Oauth.Domain)
	fetchedKeySet, err := jwk.Fetch(context.Background(), keyURL)
	if err != nil {
		fmt.Printf("Failed to fetch key set: %v", err)
		panic(err)
	}
	jwkSet = fetchedKeySet
}
