package auth

import "github.com/gin-gonic/gin"

func SetAuthEndpoints(r *gin.Engine) {
	r.GET("/callback", callbackHandler)
	r.GET("/login", loginHandler)
}
