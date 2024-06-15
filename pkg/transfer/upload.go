package transfer

import (
	"h24s_04/storage"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// IUploadImageService .
type ITransferFileService interface {
	UploadFile(c echo.Context) error
}

// uploadImageService .
type transferFileService struct {
	uu storage.ITransferFile
}

// NewUploadImageService はUploadImageServiceインスタンスを生成
func Service(uu storage.ITransferFile) ITransferFileService {
	return &transferFileService{uu}
}

// HandleUploadImage .
func (h *transferFileService) UploadFile(ctx echo.Context) error {
	// ファイルをリクエストから取得
	// このサンプルでは、「image」というパラメータで送られてくることを想定
	file, err := ctx.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid file")
	}

	// ファイルの開閉処理
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not open file")
	}
	defer src.Close()

	// データを読み込み
	fileData, err := io.ReadAll(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not read file")
	}

	// usecaseのロジックを呼び出す
	fileName := file.Filename
	url, err := h.uu.UploadFile(ctx.Request().Context(), fileData, fileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error uploading image: "+err.Error())
	}

	// リクエスト元が参照できるURLの文字列が欲しいので、Stringで返却する
	return ctx.String(http.StatusOK, url)
}
