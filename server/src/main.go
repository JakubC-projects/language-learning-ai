package main

import (
	"github.com/JakubC-projects/language-learning-ai/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadRoutes(r)

	r.Run()
}
