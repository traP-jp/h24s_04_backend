package slide

import (
	"database/sql"
	"errors"
	"fmt"
	"h24s_04/pkg/model"
	"net/http"

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
	var slides []model.Slide
	orderby := ctx.QueryParam("orderby")
	if orderby != "genre_id" && orderby != "title" && orderby != "posted_at" && orderby != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "`orderby` must equals `genre_id`, `title`, `posted_at` or ``.")
	}
	sortorder := ctx.QueryParam("sortorder")
	if sortorder != "ASC" && sortorder != "DESC" && sortorder != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "`sortorder` must equals `ASC`, `DESC` or ``.")
	}
	err := s.db.Select(&slides, fmt.Sprintf("SELECT * FROM `Slide` ORDER BY %s %s", orderby, sortorder))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, slides)
}

// func (s *SlideService) PostSlides(ctx echo.Context) error {
// 	// var slide *model.Slide とctx.Bind(slide)してくれれば
// 	// responseでpostしたslideを返す

// }

func (s *SlideService) GetSlidesSlideid(ctx echo.Context) error {
	slideid := ctx.Param("slideid")
	var slide model.Slide
	err := s.db.Get(&slide, "SELECT * FROM `Slide` WHERE `id` = ?", slideid)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%+v", err))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, slide);
}

// func (s *SlideService) PatchSlidesSlideid(ctx echo.Context) error {

// }

// func (s *SlideService) PatchSlidesSlideid(ctx echo.Context) error {

// }

func (s *SlideService) DeleteSlidesSlideid(ctx echo.Context) error {
	slideid := ctx.Param("slideid")
	var res model.Slide
	err := s.db.Get(&res, "SELECT * FROM `Slide` WHERE `id` = ?", slideid)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%+v", err))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	_, err = s.db.Exec("DELETE FROM `Slide` WHERE `id` = ?", slideid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, res)
}
