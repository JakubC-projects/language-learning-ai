package middleware

import (
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) {
	_, err := auth.GetUser(c)
	if err == nil {
		c.Header("HX-Redirect", "/login")
		c.Header("HX-Push", "/login")
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.Abort()
	}
}
