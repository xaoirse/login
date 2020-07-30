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
