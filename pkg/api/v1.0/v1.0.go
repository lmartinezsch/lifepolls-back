package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/lmartinezsch/lifepolls-back/pkg/api/v1.0/auth"
	"github.com/lmartinezsch/lifepolls-back/pkg/api/v1.0/post"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
		post.ApplyRoutes(v1)
	}
}
