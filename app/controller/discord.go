package controller

import (
	"encoding/json"
	"fmt"
	"io"
	cv "lol-team-maker/constvariables"
	"lol-team-maker/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var discordOAuthEndpoint = oauth2.Endpoint{
	AuthURL:  cv.DiscordAuthorizeEndpoint,
	TokenURL: cv.DiscordTokenEndpoint,
}

var DiscordOauthConfig = &oauth2.Config{
	ClientID:     cv.DiscordClientID,
	ClientSecret: cv.DiscordClientSecret,
	RedirectURL:  cv.DiscordCallback,
	Scopes:       []string{"identify", "email"},
	Endpoint:     discordOAuthEndpoint,
}

func DiscordGetUserInfo(c echo.Context, accessToken string) (models.DiscordUser, error) {
	var user models.DiscordUser

	client := DiscordOauthConfig.Client(c.Request().Context(), &oauth2.Token{AccessToken: accessToken})
	resp, err := client.Get(cv.DiscordIdentifyEndpoint)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(body, &user); err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func CreateUserSession(v models.UserSession) {
	models.Db.Create(&v)
}

func UpsertUserSession(v models.UserSession) {
	var UserSession models.UserSession
	res := models.Db.Where("email = ?", v.DiscordUser.Email).First(&UserSession)
	if res.RowsAffected > 0 {
		models.Db.Save(&UserSession)
	} else {
		models.Db.Create(&v)
	}
}

func GetUserSessions() []models.UserSession {
	var UserSessions []models.UserSession
	models.Db.Find(&UserSessions)
	return UserSessions
}

func DeleteUserSessionByEmail(email string) {
	var UserSession models.UserSession
	res := models.Db.Where("email = ?", email).First(&UserSession)
	if res.RowsAffected > 0 {
		models.Db.Delete(&UserSession)
	}
}

func DeleteUserSessionByID(id uint) {
	var UserSession models.UserSession
	res := models.Db.Where("id = ?", id).First(&UserSession)
	if res.RowsAffected > 0 {
		models.Db.Delete(&UserSession)
	}
}
