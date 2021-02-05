package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// function to handle connecting to dynamodb
// @return dynamodb service
func New() *dynamodb.DynamoDB {
	return dynamodb.New(
		session.New(),
		aws.NewConfig().WithRegion("us-west-2"),
	)
}
