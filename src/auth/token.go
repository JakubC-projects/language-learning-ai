package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/JakubC-projects/language-learning-ai/src/models"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func GetCallbackToken(c *gin.Context) (*oauth2.Token, error) {
	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		return nil, errors.New("invalid state parameter")
	}

	token, err := Auth.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		return nil, errors.New("failed to exchange an authorization code for a token")
	}
	return token, nil
}

func AddStateToSession(c *gin.Context) (string, error) {
	state, err := generateRandomState()
	if err != nil {
		return state, fmt.Errorf("cannot generate state: %w", err)
	}

	// Save the state inside the session.
	session := sessions.Default(c)
	session.Set("state", state)
	if err := session.Save(); err != nil {
		return state, fmt.Errorf("cannot save state: %w", err)
	}

	return state, nil
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	fmt.Println(rawIDToken)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: Auth.ClientID,
	}

	return Auth.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

type IdTokenClaims struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Subject string `json:"subject"`
}

func GetUserFromIdToken(t *oidc.IDToken) models.User {
	var claims IdTokenClaims
	err := t.Claims(&claims)
	fmt.Println(err)
	fmt.Println(claims)

	return models.User{
		Id:             claims.Subject,
		Name:           claims.Name,
		ProfilePicture: claims.Picture,
	}
}
