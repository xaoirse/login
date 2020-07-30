package model

import (
	"github.com/jinzhu/gorm"
)

// Admin is model for admin user
type Admin struct {
	gorm.Model
	Username string
	Password string
	Name     string
	Family   string
	Number   string
}
