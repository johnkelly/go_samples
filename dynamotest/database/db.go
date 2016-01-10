package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//DB is the dynamodb connection
type DB struct {
	SVC *dynamodb.DynamoDB
}

//New creates a new dynamodb connection
func New(publicKey, secretKey string) *DB {
	creds := credentials.NewStaticCredentials(publicKey, secretKey, "")

	return &DB{
		SVC: dynamodb.New(session.New(), aws.NewConfig().WithCredentials(creds).WithRegion("us-east-1")),
	}
}
