package middleware

import (
	ctr "team-builder/controller"
	"team-builder/models"

	"github.com/labstack/echo/v4"
)

func VisitorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		visitor := models.Visitor{
			IP:    ip,
			Count: 0,
		}
		ctr.UpsertVisitor(visitor)
		return next(c)
	}
}
