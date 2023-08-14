package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	fmt.Println("checking session", user)
	if user == nil {
		c.Header("HX-Redirect", "/login")
		c.Header("HX-Push", "/login")
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.Abort()
	}
}
