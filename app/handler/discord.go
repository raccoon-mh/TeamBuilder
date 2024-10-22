package handler

import (
	"encoding/json"
	"fmt"
	"log"
	ctr "lol-team-maker/controller"
	"lol-team-maker/models"
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
	userInfo, err := ctr.DiscordGetUserInfo(c, tokenReq.AccessToken)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	var userSession models.UserSession
	userSession.Token = *tokenReq
	userSession.DiscordUser = userInfo
	ctr.UpsertUserSession(userSession)

	userInfoJSON, err := json.Marshal(userInfo)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	log.Println("tokenReq", tokenReq)

	html := fmt.Sprintf(`
		<script>
			window.opener.postMessage({ token:'%s', user: '%s' }, '*');
			window.close();
		</script>`,
		tokenReq.AccessToken,
		string(userInfoJSON),
	)

	return c.HTML(http.StatusOK, html)
}

func DiscordUserInfo(c echo.Context) error {
	userInfo, err := ctr.DiscordGetUserInfo(c, "")
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, userInfo)
}

func GetUserSessions(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"res": ctr.GetUserSessions()})
}

func DiscordTest(c echo.Context) error {
	cookie, err := c.Cookie("authToken")
	if err != nil {
		if err == http.ErrNoCookie {
			return c.JSON(http.StatusOK, map[string]interface{}{"test": http.ErrNoCookie})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"test": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"test": cookie.Value})
}
