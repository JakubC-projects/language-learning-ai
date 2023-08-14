package pages

import (
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/router/middleware"
	"github.com/gin-gonic/gin"
)

func LoadPages(r *gin.Engine) {
	pageRouter := r.Group("/", middleware.HTMXRedirect)
	pageRouter.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/login.gohtml", nil)
	})

	authenticatedRouter := r.Group("/", middleware.HTMXRedirect, middleware.IsAuthenticated)
	authenticatedRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/home.gohtml", Login{})
	})

}
