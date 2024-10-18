package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	cv "lol-team-maker/constvalue"
	"lol-team-maker/handler"
	mw "lol-team-maker/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	fmt.Print(cv.StartBanner)
	e.Use(mw.Logger)
	e.Use(mw.Limitter)

	e.GET("/alive", handler.Alive)

	e.Logger.Fatal(e.Start(":8000"))
}
