package router

import (
	"github.com/JakubC-projects/language-learning-ai/router/pages"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	loadLogin(r)
	loadTemplates(r)
	pages.LoadPages(r)
}
