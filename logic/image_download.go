package logic

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func ImageDownload(c echo.Context) error {
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

	svc := s3.New(sess)
	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatal(err)
	}

	rc := obj.Body
	defer rc.Close()
	buf := make([]byte, 100)
	_, err = rc.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)
}
