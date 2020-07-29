package model

import (
	"github.com/jinzhu/gorm"
)

type InternshipModel struct {
	gorm.Model
	Name    string
	Actions []Action `gorm:"many2many:internship_model_actions"`
}
type Internship struct {
	gorm.Model
	Name              string
	Masters           []Master `gorm:"many2many:internship_masters"`
	InternshipModel   InternshipModel
	InternshipModelID int
}

func init() {
	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		checkErr(closeErr)
	}()
	checkErr(err)
	db.AutoMigrate(&InternshipModel{})
	db.AutoMigrate(&Internship{})

}
