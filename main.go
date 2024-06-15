package main

import (
	"context"
	ping "h24s_04/pkg/_ping"
	"h24s_04/pkg/genre"
	"h24s_04/pkg/setup"
	"h24s_04/pkg/transfer"
	"h24s_04/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	_db := setup.DBsetup()

	gs := genre.Service(_db)
	//ss := slide.Service(_db)
	uu, err := storage.NewTransferFileService(context.Background())
	if err != nil {
		// エラーハンドリング: uploadは外部サービスを前提にしているので、接続できない場合はpanic
		panic("failed to initialize UploadImageUsecase: " + err.Error())
	}
	tr := transfer.Service(uu)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ping", ping.Ping)

	e.POST("/genres", gs.PostGenres)
	e.GET("/genres/:genreid", gs.GetGenresGenreid)
	e.GET("/genres", gs.GetGenres)

	e.PATCH("/genres/:genreid", gs.PatchGenresGenreid)
	e.DELETE("/genres/:genreid", gs.DeleteGenresGenreid)

	e.POST("/upload", tr.UploadFile)

	e.Start(":8080")
}
