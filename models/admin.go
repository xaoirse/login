package models

import (
	"log"

	"github.com/jinzhu/gorm"
	// init sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func checkErr(err error) {
	if err != nil {
		log.Println("models/admin.go:", err)
	}
}

// Admin is model for admin user
type Admin struct {
	gorm.Model
	username string
	password string
}

// TODO NewAdmin(username , password)

// GetAdmiByUsername get username as string and return an admin
func GetAdmiByUsername(dbName *string, username *string) {
	db, err := gorm.Open("sqlite3", "test.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)
	// TODO return admin

}
