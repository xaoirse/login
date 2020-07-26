package model

import (
	"log"

	"github.com/jinzhu/gorm"
	// init sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln("models/admin.go:", err)
	}
}

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
	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)
	db.AutoMigrate(&Admin{})
}

// NewAdmin make an admin by username and password
func NewAdmin(username, password, name, family, Number string) (*Admin, bool) {

	theAdmin, ok := GetAdmiByUsername(username)
	if ok {
		return theAdmin, false
	}

	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)

	newAdmin := Admin{
		Username: username,
		Password: password,
		Name:     name,
		Family:   family,
		Number:   Number,
	}
	db.Create(&newAdmin)
	return &newAdmin, true
}

// GetAdmiByUsername get username as string and return an admin
func GetAdmiByUsername(username string) (*Admin, bool) {
	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)
	admin := Admin{
		Username: username,
	}
	// TODO return admin
	count := 0
	db.Where("username = ?", username).First(&admin).Count(&count)
	if count == 0 {
		return nil, false
	}
	return &admin, true
}
