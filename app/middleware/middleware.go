package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var LoggerConfig = middleware.RequestLoggerConfig{
	LogStatus:   true,
	LogURI:      true,
	LogMethod:   true,
	LogProtocol: true,
	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		log.Printf("[Logger] %v %v %v %v\n", v.Protocol, v.Method, v.Status, v.URI)
		return nil
	},
}
