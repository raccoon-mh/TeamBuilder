package middleware

import (
	ctr "lol-team-maker/controller"
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
		ctr.UpsertVisitor(visitor)
		return next(c)
	}
}