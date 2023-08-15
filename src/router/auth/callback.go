package auth

import (
	"errors"
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/JakubC-projects/language-learning-ai/src/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func callbackHandler(c *gin.Context) {

	_, err := getCallbackToken(c)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}

	err = auth.SaveUserSession(c, models.User{Username: "Test"})

	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}

	// Redirect to logged in page.
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func getCallbackToken(c *gin.Context) (*oauth2.Token, error) {
	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		return nil, errors.New("invalid state parameter")
	}

	token, err := auth.Auth.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		return nil, errors.New("failed to exchange an authorization code for a token")
	}
	return token, nil
}
