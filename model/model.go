package model

import (
	"log"

	"github.com/jinzhu/gorm"
	// init sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetDb initiate a DB and AutoMigrate all models
func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln("Error in opening db:", err)
	}
	db.AutoMigrate(
		&User{},
		&Action{},
		&Internship{},
		&InternshipModel{},
		&Log{},
		&Session{},
	)
	return db
}
