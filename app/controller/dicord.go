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

func DiscordGetUserInfo(c echo.Context, accesstoken string) (models.DiscordUser, error) {
	var user models.DiscordUser

	client := DiscordOauthConfig.Client(c.Request().Context(), &oauth2.Token{AccessToken: accesstoken})
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
