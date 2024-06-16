package storage

import (
	"context"
	"errors"
	"fmt"
	"h24s_04/firebase"
	"log"
	"os"
)

type ITransferFile interface {
	UploadFile(ctx context.Context, fileData []byte, fileName string) (string, string, error)
	UpdateURL(filePath string) (string, error)
}

type TransferFileService struct {
	storageClient *firebase.FirebaseStorage
}

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

func (t *TransferFileService) UploadFile(ctx context.Context, fileData []byte, fileName string) (string, string, error) {
	if len(fileData) == 0 {
		return "", "", errors.New("file data is empty")
	}

	path := fmt.Sprintf("files/%s", fileName)
	bucketName := os.Getenv("FIREBASE_STORAGE_BUCKET")

	// firebase/storage.goの保存処理を呼び出す
	url, err := t.storageClient.Upload(ctx, bucketName, fileData, path)
	if err != nil {
		return "", "", fmt.Errorf("failed to upload image to Firebase Storage: %v", err)
	}

	return url, path, nil
}

func (t *TransferFileService) UpdateURL(filePath string) (string, error) {

	bucketName := os.Getenv("FIREBASE_STORAGE_BUCKET")
	log.Println(bucketName)

	url, err := t.storageClient.GenerateSignedURL(bucketName, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to Firebase Storage: %v", err)
	}

	return url, nil
}
