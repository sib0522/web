package lib

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ssm"
	"golang.org/x/xerrors"
)

type AWS struct {
}

func NewAWSService() *AWS {
	return &AWS{}
}

// ファイルをs3にアップロード
func (r *AWS) UploadMultiple(filePathList []string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	// アップロードするファイルリストを作る
	fileList := make([]*os.File, 0, len(filePathList))
	for _, path := range filePathList {
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			fileList = append(fileList, file)
			return nil
		}()
		if err != nil {
			return xerrors.Errorf("failed to open file")
		}
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	s3Client := s3.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})

	for _, file := range fileList {
		// アップロードする
		_, err := s3Client.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("AWS_RESOURCE_BUCKET")),
			Key:    aws.String("uploads/" + file.Name()),
			Body:   file,
		})

		// アップロードしたファイルを閉じる（失敗しても閉じる）
		file.Close()
		if err != nil {
			return xerrors.Errorf("err : %w", err)
		}
	}

	return nil
}

// ファイルをs3からダウンロード
func (r *AWS) DownloadMultiple(prefix string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	s3Client := s3.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_RESOURCE_BUCKET")),
		Prefix: aws.String(prefix),
	}

	objectList, err := s3Client.ListObjectsV2(input)
	if err != nil {
		return xerrors.Errorf("err : %w", err)
	}

	fileList := make([]*os.File, 0, len(objectList.Contents))

	for _, obj := range objectList.Contents {
		key := aws.StringValue(obj.Key)
		fmt.Println("download start : ", key)

		output, err := s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("AWS_RESOURCE_BUCKET")),
			Key:    aws.String(key),
		})
		if err != nil {
			xerrors.Errorf("err : %w", err)
			continue
		}

		fileName := strings.Replace(key, "/", "-", -1)
		fileName = fmt.Sprintf("%v/%v", "web/public/images", fileName)

		err = func() error {
			file, err := os.Create(fileName)
			if err != nil {
				return err
			}

			_, err = file.ReadFrom(output.Body)
			if err != nil {
				return err
			}
			fileList = append(fileList, file)

			if _, err = io.Copy(file, output.Body); err != nil {
				return err
			}

			file.Close()
			fmt.Println("download complete : ", key)
			return nil
		}()

		if err != nil {
			return xerrors.Errorf("err : %w", err)
		}
	}

	fmt.Println("download is complete")
	return nil
}

func (r *AWS) DonwloadConfig() ([]byte, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, xerrors.Errorf("err : %w", err)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ssmClient := ssm.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})

	output, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/config/config.yaml"),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return nil, xerrors.Errorf("err : %w", err)
	}

	param := []byte(*output.Parameter.Value)
	return param, nil
}
