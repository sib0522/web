package logic

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func Boot(c echo.Context) []byte {
	// 東京リージョン
	const region = "ap-northeast-1"

	// セッション作成
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(region)},
		Profile: "default",
	})
	if err != nil {
		return xerrors.Errorf("failed to create new session")
	}

	svc := ssm.New(sess)

	parameter := "\t/web/config.yml"
	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(parameter),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return xerrors.Errorf("failed to get parameter")
	}

	// パラメータをバイト列で返す
	value := *res.Parameter.Value
	return []byte(value)
}
