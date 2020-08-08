package model

import (
	"github.com/jinzhu/gorm"
)

type Log struct {
	gorm.Model
	Student User
	Action  Action
	Master  User
}
