package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Logger = middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	LogStatus: true, LogURI: true, LogMethod: true, LogProtocol: true, LogRemoteIP: true,
	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		log.Printf("[Logger] %v - %v %v %v %v\n", v.RemoteIP, v.Protocol, v.Method, v.Status, v.URI)
		return nil
	},
})
