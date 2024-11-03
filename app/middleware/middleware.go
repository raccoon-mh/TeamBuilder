package middleware

import (
	"log"
	"net/http"

	cv "team-builder/constvariables"
	ctr "team-builder/controller"

	"github.com/labstack/echo/v4"
)

func IsCookieAuthedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookietoken, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println(http.ErrNoCookie)
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Unauthorized",
				})
			}
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		c.Set("token", cookietoken.Value)
		return next(c)
	}
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := ctr.GetUserSessionByToken(c.Get("token").(string))
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		if sess.DiscordUser.Email != cv.ADMINEMAIL {
			log.Printf("Unauthorized User : %s", sess.DiscordUser.Email)
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		}
		return next(c)
	}
}
