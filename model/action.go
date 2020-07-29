package model

import (
	"github.com/jinzhu/gorm"
)

type Action struct {
	gorm.Model
	Name             string
	InternshipModels []InternshipModel `gorm:"many2many:internship_model_actions"`
}

func init() {
	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)
	db.AutoMigrate(&Action{})
}
