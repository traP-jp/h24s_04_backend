package storage

import (
	"context"
	"errors"
	"fmt"
	"h24s_04/firebase"
	"os"
)

// ITransferFile
type ITransferFile interface {
	UploadFile(ctx context.Context, fileData []byte, fileName string) (string, error)
	// ここにdownloadも書く
}

type TransferFileService struct {
	// storage.goで作った構造体を受け取る
	storageClient *firebase.FirebaseStorage
}

// NewUploadFileService は UploadFileService の新しいインスタンスを生成
func NewTransferFileService(ctx context.Context) (ITransferFile, error) {
	// firebaseAppの初期化
	firebaseApp, err := firebase.InitFirebaseApp(ctx)
	if err != nil {
		panic(err)
	}

	// Storageクライアントの初期化
	storageClient, err := firebase.NewFirebaseStorage(ctx, firebaseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize storage client: %v", err)
	}

	return &TransferFileService{
		storageClient: storageClient,
	}, nil
}

// UploadFile は画像データを受け取り、それを外部サービスへアップロードし、アップロードされた画像のURLを返却する
func (t *TransferFileService) UploadFile(ctx context.Context, fileData []byte, fileName string) (string, error) {
	if len(fileData) == 0 {
		return "", errors.New("file data is empty")
	}

	path := fmt.Sprintf("files/%s", fileName) // 画像の保存先パスを構築、ディレクトリ中の/images/の中に保存される
	bucketName := os.Getenv("FIREBASE_STORAGE_BUCKET")

	// sdk/firebase/storage.goの保存処理を呼び出す
	url, err := t.storageClient.Upload(ctx, bucketName, fileData, path)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to Firebase Storage: %v", err)
	}

	return url, nil
}
