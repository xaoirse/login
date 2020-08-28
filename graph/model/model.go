package model

import (
	"log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	// init sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetDb creates a sqlite db
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
	)
	return db
}

func (action *Action) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
