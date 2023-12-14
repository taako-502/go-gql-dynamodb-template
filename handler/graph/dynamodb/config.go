package ddbmanager

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// 引数はオプショナルにすること
func New(endpoint string) *dynamo.DB {
	var sess *session.Session
	if os.Getenv("ENVIRONMENT") == "local" {
		sess = session.Must(session.NewSession(getAwsConfig(endpoint)))
	} else {
		sess = session.Must(session.NewSession())
	}

	// DynamoDBサービスクライアントを作成
	return dynamo.New(sess)
}

func getAwsConfig(endpoint string) *aws.Config {
	return &aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(
			"dammy",
			"dammy",
			"dammy",
		),
	}
}
