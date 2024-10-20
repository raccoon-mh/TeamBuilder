package discord

import (
	"fmt"
	"io/ioutil"
	cv "lol-team-maker/constvariables"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var discordOAuthEndpoint = oauth2.Endpoint{
	AuthURL:  cv.DiscordAuthorizeEndpoint,
	TokenURL: cv.DiscordTokenEndpoint,
}

var oauthConfig = &oauth2.Config{
	ClientID:     cv.DiscordClientID,
	ClientSecret: cv.DiscordClientSecret,
	RedirectURL:  cv.DiscordCallback,
	Scopes:       []string{"identify", "email"},
	Endpoint:     discordOAuthEndpoint,
}

func DiscordLogin(c echo.Context) error {
	url := oauthConfig.AuthCodeURL("state-token")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func DiscordCallback(c echo.Context) error {
	code := c.QueryParam("code")
	tokenReq, err := oauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	client := oauthConfig.Client(c.Request().Context(), tokenReq)
	resp, err := client.Get(cv.DiscordIdentifyEndpoint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return c.JSON(http.StatusOK, tokenReq)
}
