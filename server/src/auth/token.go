package auth

import (
	"fmt"

	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/JakubC-projects/language-learning-ai/src/models"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwt"
)

func VerifyToken(c *gin.Context) (models.AuthContext, error) {

	token, err := parseToken(c)
	if err != nil {
		return models.AuthContext{}, fmt.Errorf("cannot parse token: %w", err)
	}

	err = jwt.Validate(
		token,
		jwt.WithAudience(config.C.Oauth.Audience),
	)
	if err != nil {
		return models.AuthContext{}, fmt.Errorf("invalid token: %w", err)
	}

	userId := token.Subject()
	return models.AuthContext{
		UserId: userId,
	}, nil
}

func parseToken(c *gin.Context) (jwt.Token, error) {
	token, err := jwt.ParseRequest(
		c.Request,
		jwt.WithKeySet(jwkSet),
		jwt.WithHeaderKey("Authorization"),
	)
	return token, err
}
