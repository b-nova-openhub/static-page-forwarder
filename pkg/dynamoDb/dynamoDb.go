package dynamoDb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"static-page-exposer/pkg/util"
)

type Item struct {
	Key   string
	Value []string
}

func initSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("eu-central-1"),
		Endpoint: aws.String("http://localhost:8000")})
	util.HandleError(err)
	return sess, err
}

func ListAllTables() {
	sess, err := initSession()
	dbSvc := dynamodb.New(sess)

	result, err := dbSvc.ListTables(&dynamodb.ListTablesInput{})
	util.HandleError(err)

	for _, table := range result.TableNames {
		log.Println(*table)
	}
}

func GetItemByKey(tableName string, key string) []string {
	sess, err := initSession()
	dbSvc := dynamodb.New(sess)
	item := getItem(tableName, key, err, dbSvc)
	return convertItemToArray(item)
}

func convertItemToArray(item map[string]*dynamodb.AttributeValue) []string {
	obj := Item{}
	dynamodbattribute.UnmarshalMap(item, &obj)
	return obj.Value
}

func getItem(tableName string, key string, err error, dbSvc *dynamodb.DynamoDB) map[string]*dynamodb.AttributeValue {
	result, err := dbSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(key),
			},
		},
	})

	util.HandleError(err)
	return result.Item
}
