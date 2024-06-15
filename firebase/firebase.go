package firebase

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func InitFirebaseApp(ctx context.Context) (*firebase.App, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")

		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	// 環境変数からアカウント情報を読み込み
	serviceAccount := os.Getenv("FIREBASE_SERVICE_ACCOUNT")
	if serviceAccount == "" {
		return nil, fmt.Errorf("FIREBASE_SERVICE_ACCOUNT environment variable is not set")
	}

	opt := option.WithCredentialsJSON([]byte(serviceAccount))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	return app, nil
}
