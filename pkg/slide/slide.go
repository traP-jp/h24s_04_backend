package slide

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type SlideService struct {
	db *sqlx.DB
}

func Service(db *sqlx.DB) *SlideService {
	return &SlideService{db: db}
}

func (s *SlideService) GetSlides(ctx echo.Context) error {
	// var slides []model.Slide みたいに定義してくれれば
	// クエリパラメータを3つ(ジャンル,ソート順,ページ) 命名はそちらに任せます
	// responseで検索対象のslides

}

func (s *SlideService) PostSlides(ctx echo.Context) error {
	// var slide *model.Slide とctx.Bind(slide)してくれれば
	// responseでpostしたslideを返す

}

func (s *SlideService) GetSlidesSlideid(ctx echo.Context) error {

}

func (s *SlideService) PatchSlidesSlideid(ctx echo.Context) error {

}

func (s *SlideService) DeleteSlidesSlideid(ctx echo.Context) error {

}
