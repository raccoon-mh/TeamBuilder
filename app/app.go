package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	cv "team-builder/constvariables"
	"team-builder/handler"
	mw "team-builder/middleware"
	"team-builder/models"
)

func main() {
	e := echo.New()

	banner(e)
	models.InitDB()

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{cv.PAGESHOST},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))

	e.Use(mw.Logger)
	e.Use(mw.Limitter)

	e.Use(mw.VisitorMiddleware)

	e.GET("/alive", handler.Alive)

	defaultGroup := e.Group("/admin")
	defaultGroup.Use(mw.IsCookieAuthedMiddleware)
	defaultGroup.Use(mw.IsAdmin)

	visitorsGroup := defaultGroup.Group("/visitor")
	visitorsGroup.GET("/list", handler.GetVisitors)

	// visitorsGroup := defaultGroup.Group("/visitor")
	// defaultGroup.GET("/t", handler.DiscordTest)

	authGroup := e.Group("/auth")
	authGroup.GET("/login", handler.DiscordLogin)
	authGroup.GET("/callback", handler.DiscordCallback)

	userGroup := e.Group("/user")
	userGroup.Use(mw.IsCookieAuthedMiddleware)
	userGroup.GET("/info", handler.DiscordUserInfo)

	e.Logger.Fatal(e.Start(":" + cv.PORT))
}

func banner(e *echo.Echo) {
	e.HideBanner = true
	fmt.Print(cv.StartBanner)
}
