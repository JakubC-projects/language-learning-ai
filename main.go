package main

import (
	"github.com/JakubC-projects/language-learning-ai/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sessionStore := cookie.NewStore([]byte("secret"))
	sessionStore.Options(sessions.Options{
		Secure:   true,
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("login", sessionStore))

	router.LoadRoutes(r)

	r.Run()
}
