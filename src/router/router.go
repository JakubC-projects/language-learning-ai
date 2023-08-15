package router

import (
	"github.com/JakubC-projects/language-learning-ai/src/router/auth"
	"github.com/JakubC-projects/language-learning-ai/src/router/html"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	loadTemplates(r)
	loadMessages(r)
	auth.SetAuthEndpoints(r)
	html.LoadRoutes(r)
}
