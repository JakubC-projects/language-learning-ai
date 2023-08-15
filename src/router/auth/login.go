package auth

import (
	"fmt"
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/gin-gonic/gin"
)

// Handler for our login.
func loginHandler(ctx *gin.Context) {
	state, err := auth.AddStateToSession(ctx)
	if err != nil {
		errCtx := fmt.Errorf("cannot save session: %w", err)
		ctx.String(http.StatusInternalServerError, errCtx.Error())
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, auth.Auth.AuthCodeURL(state))
}
