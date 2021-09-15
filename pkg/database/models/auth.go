package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/lifepolls-back/pkg/lib/common"
)

// Auth data model
type Auth struct {
	gorm.Model
	Username     string
	DisplayName  string
	PasswordHash string
}

// Serialize serializes Auth data
func (u *Auth) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"display_name": u.DisplayName,
	}
}

func (u *Auth) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.DisplayName = m["display_name"].(string)
}
