package transfer

import (
	"h24s_04/storage"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITransferFileService interface {
	UploadFile(c echo.Context) error
}

type transferFileService struct {
	uu storage.ITransferFile
}

func Service(uu storage.ITransferFile) ITransferFileService {
	return &transferFileService{uu: uu}
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
	url, err := h.uu.UploadFile(ctx.Request().Context(), fileData, fileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error uploading image: "+err.Error())
	}

	return ctx.String(http.StatusOK, url)
}
