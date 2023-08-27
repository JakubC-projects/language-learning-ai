package auth

import (
	"errors"

	"github.com/JakubC-projects/language-learning-ai/src/models"
	"github.com/gin-gonic/gin"
)

var userKey = "auth"

func GetUser(c *gin.Context) models.AuthContext {
	authCtxVal, ok := c.Get(userKey)
	if !ok {
		panic("user not set")
	}
	authCtx, ok := authCtxVal.(models.AuthContext)
	if !ok {
		panic("invalid user type")
	}
	return authCtx
}

func SetUser(c *gin.Context, a models.AuthContext) {
	c.Set(userKey, a)
}

var ErrNotAuthenticated = errors.New("not authenticated")
