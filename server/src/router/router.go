package router

import (
	"github.com/JakubC-projects/language-learning-ai/src/router/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	authenticated := r.Group("/", middleware.IsAuthenticated)
	authenticated.GET("/chats", GetChats)
	r.POST("/chats", createChat)
	authenticated.GET("/messages", GetChats)
}
