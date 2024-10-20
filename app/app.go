package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	cv "lol-team-maker/constvariables"
	"lol-team-maker/handler"
	mw "lol-team-maker/middleware"
	"lol-team-maker/models"
)

func main() {
	e := echo.New()

	banner(e)
	models.InitDB()

	e.Use(mw.Logger)
	e.Use(mw.Limitter)

	e.Use(mw.VisitorMiddleware)

	e.GET("/alive", handler.Alive)
	e.GET("/v", handler.GetVisitors)

	e.GET("/login", handler.DiscordLogin)
	e.GET("/callback", handler.DiscordCallback)
	e.GET("/userinfo", handler.DiscordUserInfo)

	e.Logger.Fatal(e.Start(":" + cv.PORT))
}

func banner(e *echo.Echo) {
	e.HideBanner = true
	fmt.Print(cv.StartBanner)
}
