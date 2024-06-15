package firebase

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	cs "cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
)

type FirebaseStorage struct {
	Client *storage.Client
}

func NewFirebaseStorage(ctx context.Context, app *firebase.App) (*FirebaseStorage, error) {
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseStorage{Client: client}, nil
}

func (fs *FirebaseStorage) Upload(ctx context.Context, bucketName string, fileData []byte, path string) (string, error) {
	// バケットの参照を取得
	bucket, err := fs.Client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("failed to get bucket: %v", err)
	}

	// ファイルのContentTypeを推測
	contentType := http.DetectContentType(fileData)

	// ファイルへの書き込み用のWriterを作成
	wc := bucket.Object(path).NewWriter(ctx)
	wc.ContentType = contentType
	wc.CacheControl = "public, max-age=31536000" // 1年間キャッシュする

	// データをStorageにアップロード
	if _, err := wc.Write(fileData); err != nil {
		return "", fmt.Errorf("failed to write file to Firebase Storage: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// 署名付きURLの生成
	signedURL, err := fs.generateSignedURL(ctx, bucketName, path, 5) // 5年間有効
	if err != nil {
		return "", err
	}

	return signedURL, nil
}

// generateSignedURL 署名付きURLを生成
func (fs *FirebaseStorage) generateSignedURL(ctx context.Context, bucketName, objectName string, expiry time.Duration) (string, error) {
	// 署名付きURLのオプションを設定
	opts := &cs.SignedURLOptions{
		Scheme:  cs.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute), // 有効期限
	}

	// 署名付きURLを生成
	bucket, err := fs.Client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	u, err := bucket.SignedURL(objectName, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
	}
	fmt.Printf("Generated GET signed URL:\n%s\n", u)

	return u, nil
}

func (fs *FirebaseStorage) Download(ctx context.Context, bucketName string, objectName string, destFileName string) error {

	f, err := os.Create(destFileName)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	bh, err := fs.Client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("Bucket(%q): %w", bucketName, err)
	}

	rc, err := bh.Object(objectName).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("Object(%q).NewReader: %w", objectName, err)
	}

	defer rc.Close()

	if _, err := io.Copy(f, rc); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("f.Close: %w", err)
	}

	return nil

}
