package model

import (
	"github.com/jinzhu/gorm"
)

type master struct {
	gorm.Model
	Number   string
	Username string
	Password string
	Name     string
	Family   string
}

func NewMaster() {

}
func GetMasterByUsername(username string) {

}
