package database

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/satori/go.uuid"
)

//DBQuery is the struct of a query in dynamodb
type DBQuery struct {
	ID        string
	Question  string
	CreatedAt int64
	UpdatedAt int64
}

//CreateQuery adds a query to the database
func (db *DB) CreateQuery(question string) (*DBQuery, error) {
	currentTimestamp := time.Now().UTC().Unix()

	query := &DBQuery{
		ID:        uuid.NewV4().String(),
		Question:  question,
		CreatedAt: currentTimestamp,
		UpdatedAt: currentTimestamp,
	}

	params := &dynamodb.PutItemInput{
		ConditionExpression: aws.String("attribute_not_exists(id)"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(query.ID),
			},
			"question": {
				S: aws.String(query.Question),
			},
			"created_at": {
				N: aws.String(strconv.FormatInt(query.CreatedAt, 10)),
			},
			"updated_at": {
				N: aws.String(strconv.FormatInt(query.UpdatedAt, 10)),
			},
		},
		TableName: aws.String("queries"),
	}

	_, err := db.SVC.PutItem(params)
	return query, err
}

//ShowQuery gets a query by id from the database
func (db *DB) ShowQuery(id string) (*DBQuery, error) {
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("queries"),
	}

	resp, err := db.SVC.GetItem(params)
	if err != nil {
		return nil, err
	}

	createdAt, err := strconv.Atoi(*resp.Item["created_at"].N)
	if err != nil {
		return nil, err
	}

	updatedAt, err := strconv.Atoi(*resp.Item["updated_at"].N)
	if err != nil {
		return nil, err
	}

	query := &DBQuery{
		ID:        *resp.Item["id"].S,
		Question:  *resp.Item["question"].S,
		CreatedAt: int64(createdAt),
		UpdatedAt: int64(updatedAt),
	}

	return query, nil
}

//UpdateQuery updates a single key's value in the database
func (db *DB) UpdateQuery(id, key, value string) (int64, error) {
	currentTimestamp := time.Now().UTC().Unix()

	params := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			key: {
				Action: aws.String("PUT"),
				Value: &dynamodb.AttributeValue{
					S: aws.String(value),
				},
			},
			"updated_at": {
				Action: aws.String("PUT"),
				Value: &dynamodb.AttributeValue{
					N: aws.String(strconv.FormatInt(currentTimestamp, 10)),
				},
			},
		},
		TableName: aws.String("queries"),
	}

	_, err := db.SVC.UpdateItem(params)
	return currentTimestamp, err
}
