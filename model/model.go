package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln("models:", err)
	}
}

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data.db")
	defer func() {
		closeErr := db.Close()
		CheckErr(closeErr)
	}()
	CheckErr(err)
	return db
}
