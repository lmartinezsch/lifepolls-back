package middlewares

import "github.com/gin-gonic/gin"

// Authorized blocks unauthorized requestrs
func Authorized(c *gin.Context) {
	_, exists := c.Get("auth")
	if !exists {
		c.AbortWithStatus(401)
		return
	}
}
