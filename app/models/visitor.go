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
