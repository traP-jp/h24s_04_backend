package transfer

import (
	"database/sql"
	"errors"
	"fmt"
	"h24s_04/storage"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type fileUploadResponse struct {
	DL_URL    string `json:"dl_url"`
	Thumb_URL string `json:"thumb_url"`
	Path      string `json:"path"`
}

type ITransferFileService interface {
	UploadFile(c echo.Context) error
	DownloadFile(c echo.Context) error
	Urlupdate()
}

type transferFileService struct {
	uu storage.ITransferFile
	db *sqlx.DB
}

func Service(uu storage.ITransferFile, db *sqlx.DB) ITransferFileService {
	return &transferFileService{uu: uu, db: db}
}

func (h *transferFileService) UploadFile(ctx echo.Context) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid file")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not open file")
	}
	defer src.Close()

	fileData, err := io.ReadAll(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not read file")
	}

	filenametemp, _ := uuid.NewV7()

	fileName := filenametemp.String() + ".pdf"
	dl_url, path, err := h.uu.UploadFile(ctx.Request().Context(), fileData, fileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error uploading file: "+err.Error())
	}

	thumb, err := ctx.FormFile("thumb")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid file")
	}

	src, err = thumb.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not open file")
	}
	defer src.Close()

	thumbData, err := io.ReadAll(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not read file")
	}

	thumbName := ("thumb_" + fileName)
	thumb_url, _, err := h.uu.UploadFile(ctx.Request().Context(), thumbData, thumbName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error uploading file: "+err.Error())
	}

	res := &fileUploadResponse{}
	res.DL_URL = dl_url
	res.Thumb_URL = thumb_url
	res.Path = path

	return ctx.JSON(http.StatusOK, res)
}

func (h *transferFileService) DownloadFile(ctx echo.Context) error {

	slideid := ctx.Param("slideid")
	var dlfile string

	err := h.db.Get(&dlfile, "SELECT `filepath` FROM `Slide` WHERE `id`=?", slideid)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%+v", err))
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%+v", err))
	}

	dlfile = "files/" + dlfile
	dltempid, _ := uuid.NewV7()
	dltempname := dltempid.String() + ".pdf"

	err = h.uu.DownloadFile(ctx.Request().Context(), dlfile, dltempname)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error downloadinging file: "+err.Error())
	}

	return ctx.File(dltempname)
}
