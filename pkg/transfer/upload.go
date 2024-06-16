package transfer

import (
	"h24s_04/storage"
	"io"
	"net/http"

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

	fileName := file.Filename
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
	
	err:=h.uu.DownloadFile(ctx.Request().Context(),)
	return ctx.String(http.StatusOK,filename)
}
