package post

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/lifepolls-back/pkg/database/models"
	"github.com/lmartinezsch/lifepolls-back/pkg/lib/common"
)

// Post type alias
type Post = models.Post

// User type alias
type User = models.User

// Auth type alias
type Auth = models.Auth

// JSON type alias
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Text   string `json:"text" binding:"required"`
		UserID int    `json:"user_id" binding:"required"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	//auth := c.MustGet("auth").(Auth)

	var user User
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", requestBody.UserID).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	post := Post{Text: requestBody.Text, User: user}
	db.NewRecord(post)
	db.Create(&post)
	c.JSON(200, post.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var posts []Post

	if cursor == "" {
		if err := db.Preload("User").Limit(10).Order("id desc").Find(&posts).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("User").Limit(10).Order("id desc").Where(condition, cursor).Find(&posts).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(posts)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = posts[i].Serialize()
	}

	c.JSON(200, serialized)
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var post Post

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, post.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	var post Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if post.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	db.Delete(&post)
	c.Status(204)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	type RequestBody struct {
		Text string `json:"text" binding:"required"`
	}

	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var post Post
	if err := db.Preload("User").Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if post.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	post.Text = requestBody.Text
	db.Save(&post)
	c.JSON(200, post.Serialize())
}
