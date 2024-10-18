package models

import (
	"gorm.io/gorm"
)

type Visitor struct {
	gorm.Model
	IP    string `json:"IP"`
	Count uint   `json:"Count"`
}
