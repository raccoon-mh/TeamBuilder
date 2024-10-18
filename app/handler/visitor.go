package handler

import (
	"lol-team-maker/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetVisitors(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"res": models.GetVisitors()})
}
