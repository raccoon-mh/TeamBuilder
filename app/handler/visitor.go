package handler

import (
	ctr "lol-team-maker/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetVisitors(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"res": ctr.GetVisitors()})
}
