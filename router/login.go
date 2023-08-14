package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequestData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

var adminLogin = LoginRequestData{
	Username: "admin",
	Password: "admin",
}

func loadLogin(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		var loginData LoginRequestData
		err := c.Bind(&loginData)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if loginData != adminLogin {
			c.HTML(http.StatusOK, "components/error.gohtml", "invalid username or password")
			return
		}
		session := sessions.Default(c)
		session.Set("user", loginData.Username)
		session.Save()
		c.Header("HX-Redirect", "/")
		c.Status(http.StatusNoContent)
	})

	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("user")
		session.Save()
		c.Header("HX-Redirect", "/")
		c.Header("HX-Push", "/")
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})
}
