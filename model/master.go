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

func init() {
	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)
	db.AutoMigrate(&Master{})

}
func NewMaster() {

}
func GetMasterByUsername(username string) {

}
