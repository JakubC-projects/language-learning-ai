package auth

import (
	"encoding/gob"
	"errors"

	"github.com/JakubC-projects/language-learning-ai/src/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var userSessionKey = "user"

func init() {
	gob.Register(models.User{})
}

func GetUser(c *gin.Context) (models.User, error) {
	sessionValue := sessions.Default(c).Get(userSessionKey)
	user, ok := sessionValue.(models.User)
	if !ok {
		return user, ErrNotAuthenticated
	}
	return user, nil
}

func SaveUserSession(c *gin.Context, u models.User) error {
	session := sessions.Default(c)
	session.Set(userSessionKey, u)
	return session.Save()
}

func RemoveUserSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Delete(userSessionKey)
	return session.Save()
}

var ErrNotAuthenticated = errors.New("not authenticated")
