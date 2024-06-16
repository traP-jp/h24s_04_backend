package slide

import (
	"database/sql"
	"errors"
	"fmt"
	"h24s_04/pkg/model"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
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
	if orderby == "" {
		orderby = "posted_at"
	}
	if orderby != "genre_id" && orderby != "title" && orderby != "posted_at" {
		return echo.NewHTTPError(http.StatusBadRequest, "`orderby` must equals `genre_id`, `title`, `posted_at` or ``.")
	}
	sortorder := ctx.QueryParam("sortorder")
	if sortorder == "" {
		sortorder = "DESC"
	}
 	if sortorder != "ASC" && sortorder != "DESC" && sortorder != "" {
		return echo.NewHTTPError(http.StatusBadRequest, "`sortorder` must equals `ASC`, `DESC` or ``.")
	}
	err := s.db.Select(&slides, fmt.Sprintf("SELECT * FROM `Slide` ORDER BY %s %s", orderby, sortorder))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, slides)
}

func (s *SlideService) PostSlides(ctx echo.Context) error {
	slide := &model.Slide{}
	err := ctx.Bind(slide)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	slide.Id, err = uuid.NewV7()
	slide.Posted_at = time.Now()
	slide.URL_updated_at = time.Now()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	_, err = s.db.Exec("INSERT INTO `Slide` (`id`, `dl_url`, `thumb_url`, `title`, `genre_id`, `posted_at`, `description`, `url_updated_at`, `filepath`) VALUES (?,?,?,?,?,?,?,?,?)", slide.Id, slide.DL_url, slide.Thumb_url, slide.Title, slide.Genre_id, slide.Posted_at, slide.Description, slide.URL_updated_at, slide.Filepath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, slide)
}

func (s *SlideService) GetSlidesSlideid(ctx echo.Context) error {
	slideid := ctx.Param("slideid")
	var slide model.Slide
	err := s.db.Get(&slide, "SELECT * FROM `Slide` WHERE `id` = ?", slideid)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%+v", err))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}
	return ctx.JSON(http.StatusOK, slide)
}

func (s *SlideService) PatchSlidesSlideid(ctx echo.Context) error {
	slideid := ctx.Param("slideid")
	var res model.Slide
	err := s.db.Get(&res, "SELECT * FROM `Slide` WHERE `id` = ?", slideid)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%+v", err))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	req := &model.Slide{}
	err = ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", err))
	}

	// id, posted_at, url_updated_at, filepath は書き換えられない
	_, err = s.db.Exec("UPDATE `Slide` SET `dl_url` = ?, `thumb_url` = ?, `title` = ?, `genre_id` = ?, `description` = ? WHERE `id` = ?", req.DL_url, req.Thumb_url, req.Title, req.Genre_id, req.Description, slideid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	err = s.db.Get(&res, "SELECT * FROM `Slide` WHERE `id` = ?", slideid)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	return ctx.JSON(http.StatusOK, res)
}

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
