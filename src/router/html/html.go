package html

import (
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/JakubC-projects/language-learning-ai/src/chat"
	"github.com/JakubC-projects/language-learning-ai/src/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func LoadRoutes(r *gin.Engine) {
	authenticatedRouter := r.Group("/", middleware.IsAuthenticated, middleware.HTMXRedirect)
	authenticatedRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/home.html", HomePage{Messages: chat.Chat})
	})
	authenticatedRouter.GET("/components/user", func(c *gin.Context) {
		user := lo.Must(auth.GetUser(c))
		c.HTML(http.StatusOK, "components/user.html", user)
	})
}
