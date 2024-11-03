package handler

import (
	"net/http"
	ctr "team-builder/controller"

	"github.com/labstack/echo/v4"
)

func GetVisitors(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"res": ctr.GetVisitors()})
}
