package models

import (
	cv "lol-team-maker/constvariables"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	var err error
	Db, err = gorm.Open(sqlite.Open(cv.DATABASE_PATH), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	Db.AutoMigrate(&Visitor{})
}
