package genre

import (
	"database/sql"
	"errors"
	"fmt"
	"h24s_04/pkg/model"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type GenreService struct {
	db *sqlx.DB
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
	genre := &model.Genre{}
	err := ctx.Bind(genre)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	genre.Id, err = uuid.NewV7()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	_, err = s.db.Exec("INSERT INTO `Genre` (`id`, `genrename`) VALUES (?,?)", genre.Id, genre.Genre_name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, genre)

}

func (s *GenreService) GetGenresGenreid(ctx echo.Context) error {
	genreid := ctx.Param("genreid")
	var genre model.Genre
	err := s.db.Get(&genre, "SELECT * FROM Genre WHERE `id` = ?", genreid)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%+v", err))
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, genre)
}

// func (s *GenreService) PatchGenresGenreid(ctx echo.Context) error {

// }

// func (s *GenreService) DeleteGenresGenreid(ctx echo.Context) error {

// }
