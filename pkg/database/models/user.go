package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lmartinezsch/lifepolls-back/pkg/lib/common"
)

// User data model
type User struct {
	gorm.Model
	Username     string
	Name         string
	Lastname     string
	Email        string
	PasswordHash string
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":       u.ID,
		"username": u.Username,
		"name":     u.Name,
		"lastname": u.Lastname,
		"email":    u.Email,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.Name = m["name"].(string)
	u.Lastname = m["lastname"].(string)
	u.Email = m["email"].(string)
}
