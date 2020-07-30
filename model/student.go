package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Number      string
	Username    string
	Password    string
	Name        string
	Family      string
	Internships []Internship `gorm:"many2many:inernship_students"`
}

func init() {
	db := GetDb()
	db.AutoMigrate(&Student{})
}

func NewStudent(student *Student) {

}
func GetStudentByUsername(username string) {

}
func GetStudentByNumber(number string) {

}
