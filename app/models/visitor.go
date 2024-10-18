package models

import (
	"time"
)

type Visitor struct {
	ID        uint      `gorm:"primaryKey"`
	IP        string    `json:"ip"`
	Count     uint      `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}

func CreateVisitor(v Visitor) {
	Db.Create(&v)
}

func GetVisitors() []Visitor {
	var visitors []Visitor
	Db.Find(&visitors)
	return visitors
}

func UpsertVisitor(v Visitor) {
	var visitor Visitor
	res := Db.Where("ip = ?", v.IP).First(&visitor)
	if res.RowsAffected > 0 {
		visitor.Timestamp = time.Now()
		visitor.Count++
		Db.Save(&visitor)
	} else {
		Db.Create(&v)
	}
}
