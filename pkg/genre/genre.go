package genre

import (
	"h24s_04/pkg/model"

	"github.com/labstack/echo/v4"
)

func Getgenres(ctx echo.Context) error {
	// var genres []model.Genre みたいに定義してくれれば
	// ジャンル一覧を返す jsonに入れればできそう

}

func Postgenres(ctx echo.Context) error {
	// var genre *model.Genre みたいに定義してerr:=ctx.Bind(genre)みたいに
	// 登録したジャンルを返す

}

func Getgenresgenreid(ctx echo.Context) error {

}

func Patchgenresgenreid(ctx echo.Context) error {

}

func Deletegenresgenreid(ctx echo.Context) error {

}
