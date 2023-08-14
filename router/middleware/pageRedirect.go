package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTMXRedirect(c *gin.Context) {
	if !isHTMX(c) {
		c.HTML(http.StatusOK, "index.gohtml", nil)
		c.Abort()
	} else {
		c.Next()
	}
}

func isHTMX(c *gin.Context) bool {
	return c.GetHeader("Hx-Request") == "true"
}
