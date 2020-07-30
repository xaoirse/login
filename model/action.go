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
	db := GetDb()
	db.AutoMigrate(&Action{})
}
