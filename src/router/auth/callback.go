package auth

import (
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/gin-gonic/gin"
)

func callbackHandler(c *gin.Context) {

	token, err := auth.GetCallbackToken(c)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}

	idToken, err := auth.VerifyIDToken(c.Request.Context(), token)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}

	user := auth.GetUserFromIdToken(idToken)
	auth.SaveUserSession(c, user)

	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
