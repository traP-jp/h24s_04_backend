package main

import (
	ping "h24s_04/pkg/_ping"
	"h24s_04/pkg/genre"
	"h24s_04/pkg/setup"
	"h24s_04/pkg/slide"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	_db := setup.DBsetup()

	gs := genre.Service(_db)
	ss := slide.Service(_db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ping", ping.Ping)

	e.POST("/genres", gs.PostGenres)
	e.GET("/genres/:genreid", gs.GetGenresGenreid)
	e.GET("/genres", gs.GetGenres)

	e.PATCH("/genres/:genreid", gs.PatchGenresGenreid)
	e.DELETE("/genres/:genreid", gs.DeleteGenresGenreid)

	e.GET("slides/:slideid", ss.GetSlidesSlideid)

	e.Start(":8080")
}
