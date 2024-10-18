package middleware

import (
	"log"
	"time"

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

var Limitter = middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(
	middleware.RateLimiterMemoryStoreConfig{
		Rate:      1,
		Burst:     10,
		ExpiresIn: 5 * time.Minute,
	},
))
