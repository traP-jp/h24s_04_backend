package genre

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type GenreService struct {
	db *sqlx.DB
}

type Genre struct {
	id        string `json:"id,omitempty"  db:"id"`
	genrename string `json:"genrename,omitempty"  db:"genrename"`
}

func Service(db *sqlx.DB) *GenreService {
	return &GenreService{db: db}

}

// func (s *GenreService) GetGenres(ctx echo.Context) error {
// 	// var genres []model.Genre みたいに定義してくれれば
// 	// ジャンル一覧を返す jsonに入れればできそう

// }

func (s *GenreService) PostGenres(ctx echo.Context) error {
	// var genre *model.Genre みたいに定義してerr:=ctx.Bind(genre)みたいに
	// 登録したジャンルを返す
	genre := &Genre{}
	err := ctx.Bind(genre)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	return ctx.JSON(http.StatusOK, genre)

}

// func (s *GenreService) GetGenresGenreid(ctx echo.Context) error {

// }

// func (s *GenreService) PatchGenresGenreid(ctx echo.Context) error {

// }

// func (s *GenreService) DeleteGenresGenreid(ctx echo.Context) error {

// }
