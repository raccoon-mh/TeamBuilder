package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	var err error
	Db, err = gorm.Open(sqlite.Open("meta/meta.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	Db.AutoMigrate(&Visitor{})
}
