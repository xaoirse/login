package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Number      string
	NaCode      string
	Username    string
	Password    string
	Name        string
	Family      string
	Role        string // admin master student
	Phone       string
	Internships []Internship `gorm:"many2many:inernship_students"`
}
