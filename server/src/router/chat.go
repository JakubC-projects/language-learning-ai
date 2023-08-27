package router

import (
	"fmt"
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/JakubC-projects/language-learning-ai/src/db"
	"github.com/gin-gonic/gin"
)

func GetChats(c *gin.Context) {
	authCtx := auth.GetUser(c)
	chats, err := db.GetChats(c.Request.Context(), authCtx.UserId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cannot get chats: %w", err))
		return
	}
	c.JSON(http.StatusOK, chats)
}
