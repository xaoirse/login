package model

import (
	"log"

	"github.com/jinzhu/gorm"
	// init sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln("Error in opening db:", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			log.Fatalln("Error when closing db:", err)
		}
	}()
	db.AutoMigrate(
		&Admin{},
		&Action{},
		&Internship{},
		&InternshipModel{},
		&Log{},
		&Master{},
		&Student{},
	)
}
