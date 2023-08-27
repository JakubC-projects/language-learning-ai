package middleware

import (
	"fmt"
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) {
	user, err := auth.VerifyToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("not logged in: %w", err))
		return
	}
	auth.SetUser(c, user)
}
