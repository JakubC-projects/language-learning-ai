package main

import (
	"github.com/JakubC-projects/language-learning-ai/api/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
	}))
	api.LoadEndpoints(r)

	r.Run(":8081")
}
