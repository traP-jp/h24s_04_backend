package main

import (
	"context"
	ping "h24s_04/pkg/_ping"
	"h24s_04/pkg/genre"
	"h24s_04/pkg/setup"

	"h24s_04/pkg/transfer"
	"h24s_04/storage"

	"h24s_04/pkg/slide"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	_db := setup.DBsetup()

	gs := genre.Service(_db)

	ss := slide.Service(_db)
	uu, err := storage.NewTransferFileService(context.Background())
	if err != nil {
		// エラーハンドリング: uploadは外部サービスを前提にしているので、接続できない場合はpanic
		panic("failed to initialize UploadImageUsecase: " + err.Error())
	}
	tr := transfer.Service(uu, _db)

	setup.Cronsetup(tr)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/ping", ping.Ping)

	e.POST("/api/genres", gs.PostGenres)
	e.GET("/api/genres/:genreid", gs.GetGenresGenreid)
	e.GET("/api/genres", gs.GetGenres)

	e.PATCH("/api/genres/:genreid", gs.PatchGenresGenreid)
	e.DELETE("/api/genres/:genreid", gs.DeleteGenresGenreid)

	e.GET("/api/slides/:slideid", ss.GetSlidesSlideid)
	e.DELETE("/api/slides/:slideid", ss.DeleteSlidesSlideid)
	e.GET("/api/slides", ss.GetSlides)
	e.PATCH("/api/slides/:slideid", ss.PatchSlidesSlideid)
	e.POST("/api/slides", ss.PostSlides)

	e.POST("/api/upload", tr.UploadFile)

	e.Start(":8080")
}
