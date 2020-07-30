package model

import (
	"github.com/jinzhu/gorm"
	// init sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func init() {
	db := GetDb()
	db.AutoMigrate(&Admin{})
}

// NewAdmin create new admin if it was new
func NewAdmin(admin *Admin) bool {
	db := GetDb()

	// TODO validate values
	if db.NewRecord(admin) {
		db.Create(admin)
		return true
	}
	return false
}

// GetAdminByUsername get username as string and return an admin
func GetAdminByUsername(username string) (*Admin, bool) {
	db := GetDb()
	var admin Admin
	count := 0
	db.Where("username = ?", username).First(&admin).Count(&count)
	if count == 0 {
		return nil, false
	}
	return &admin, true
}

// TODO
// func EditAdmin
