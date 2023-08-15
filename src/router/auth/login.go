package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our login.
func loginHandler(ctx *gin.Context) {
	state, err := generateRandomState()
	if err != nil {
		errCtx := fmt.Errorf("cannot generate random state: %w", err)
		ctx.String(http.StatusInternalServerError, errCtx.Error())
		return
	}

	// Save the state inside the session.
	session := sessions.Default(ctx)
	session.Set("state", state)

	if err := session.Save(); err != nil {
		errCtx := fmt.Errorf("cannot save session: %w", err)
		ctx.String(http.StatusInternalServerError, errCtx.Error())
		return
	}
	fmt.Println(config.C)
	ctx.Redirect(http.StatusTemporaryRedirect, auth.Auth.AuthCodeURL(state))
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
