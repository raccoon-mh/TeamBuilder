package models

import (
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model
	oauth2.Token
	DiscordUser
}
