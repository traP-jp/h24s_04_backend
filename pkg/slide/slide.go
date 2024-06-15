package slide

import (
	"github.com/labstack/echo/v4"
)



func Getslides(ctx echo.Context) error {
	// var slides []model.Slide みたいに定義してくれれば
	// クエリパラメータを3つ(ジャンル,ソート順,ページ) 命名はそちらに任せます
	// responseで検索対象のslides

}

func Postslides(ctx echo.Context) error {
	// var slide *model.Slide とctx.Bind(slide)してくれれば
	// responseでpostしたslideを返す

}

func Getslidesslideid(ctx echo.Context) error {

}

func Patchslidesslideid(ctx echo.Context) error {

}

func Deleteslidesslideid(ctx echo.Context) error {

}
