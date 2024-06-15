package main

import (
	"h24_04/pkg/_ping"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ping", ping.Ping)

	e.Start(":8080")
}
