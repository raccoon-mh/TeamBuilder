package handler

import (
	"log"
	ctr "lol-team-maker/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DiscordLogin(c echo.Context) error {
	url := ctr.DiscordOauthConfig.AuthCodeURL("state-token")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func DiscordCallback(c echo.Context) error {
	code := c.QueryParam("code")
	tokenReq, err := ctr.DiscordOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, tokenReq)
}

func DiscordUserInfo(c echo.Context) error {
	userInfo, err := ctr.DiscordGetUserInfo(c, "")
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, userInfo)
}
