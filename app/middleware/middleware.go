package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsCookieAuthedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("authToken")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Unauthorized",
				})
			}
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		log.Printf("Auth token: %s", cookie.Value)
		return next(c)
	}
}
