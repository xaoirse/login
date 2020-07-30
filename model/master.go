package model

import (
	"github.com/jinzhu/gorm"
)

type Master struct {
	gorm.Model
	Number      string
	Username    string
	Password    string
	Name        string
	Family      string
	Internships []Internship `gorm:"many2many:internship_masters"`
}
