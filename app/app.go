package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	cv "lol-team-maker/constvariables"
	"lol-team-maker/handler"
	mw "lol-team-maker/middleware"
	"lol-team-maker/models"
)

func main() {
	e := echo.New()

	banner(e)
	models.InitDB()

	e.Use(middleware.Recover())

	e.Use(mw.Logger)
	e.Use(mw.Limitter)

	e.Use(mw.VisitorMiddleware)

	e.GET("/alive", handler.Alive)

	defaultGroup := e.Group("/")
	defaultGroup.Use(mw.IsCookieAuthedMiddleware)
	defaultGroup.GET("/v", handler.GetVisitors)
	defaultGroup.GET("/t", handler.DiscordTest)

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
