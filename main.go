package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/lilpacy/gin-nginx-next-docker-min/external"
	"os"
)


func main() {
	// クライアントの設定
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	// DynamoDB Localを利用する場合はEndpointのURLを設定する
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	// デフォルトでは東京リージョンを指定
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

	table := db.Table("Posts")
	
	router := external.Setup(table)
	router.Run(":9000")
}
