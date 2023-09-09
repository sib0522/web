package logic

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func ImageUpload(c echo.Context) error {
	// セッション作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))

	targetPath := "./sample.txt"
	file, err := os.Open(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bucketName := "bak-file-bucket"
	objectKey := "bak-key"

	// アップローダー生成
	uploader := s3manager.NewUploader(sess)
	if _, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	}); err != nil {
		log.Fatal(err)
	}
	log.Printf("upload success (%v)", file.Name())

}
