package post

import (
	"github.com/gin-gonic/gin"
	"github.com/lmartinezsch/lifepolls-back/pkg/lib/middlewares"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	post := r.Group("/posts")
	{
		post.POST("/", middlewares.Authorized, create)
		post.GET("/", list)
		post.GET("/:id", read)
		post.DELETE("/:id", middlewares.Authorized, remove)
		post.PATCH("/:id", middlewares.Authorized, update)
	}
}
