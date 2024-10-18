package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Alive(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "alive"})
}
