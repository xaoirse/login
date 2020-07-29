package model

import (
	"github.com/jinzhu/gorm"
)

type Log struct {
	gorm.Model
	Student Student
	Action  Action
	Master  Master
}

func init() {
	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)
	db.AutoMigrate(&Log{})

}
