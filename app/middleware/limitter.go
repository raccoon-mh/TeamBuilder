package middleware

import (
	"time"

	"github.com/labstack/echo/v4/middleware"
)

var Limitter = middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(
	middleware.RateLimiterMemoryStoreConfig{
		Rate:      1,
		Burst:     10,
		ExpiresIn: 5 * time.Minute,
	},
))
