package middleware

import (
	"lol-team-maker/models"
	"time"

	"github.com/labstack/echo/v4"
)

func VisitorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		visitor := models.Visitor{
			IP:        ip,
			Count:     0,
			Timestamp: time.Now(),
		}
		models.UpsertVisitor(visitor)
		return next(c)
	}
}
