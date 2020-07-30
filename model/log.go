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
	db := GetDb()
	db.AutoMigrate(&Log{})

}
