package datbase

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DBService struct {
	client *dynamodb.Client
	ctx    context.Context
	Robots []map[string]string
}

func GetDbService() *DBService {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatal(err)
	}

	dbService := DBService{}
	dbService.client = dynamodb.NewFromConfig(cfg)
	dbService.ctx = ctx

	return &dbService
}

func (db *DBService) GetRobotsForUser(userId string) map[string]string {
	// Filter Robots map so that id == userId
	for _, value := range db.Robots {
		if value["id"] == userId {
			return value
		}
	}

	if db.client == nil {
		log.Fatal("Cannot create db client")
	}

	return nil
}

func (db *DBService) CreateTables() {
	if db.client == nil {
		log.Fatal("Cannot create db client")
	}

	resp, err := db.client.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})

	if err != nil {
		log.Fatal("Cannot list tables. " + err.Error())
	}

	db.createRobotsTable()

	log.Println(resp)
}

func (db *DBService) createRobotsTable() {
	param := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("RobotId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("Name"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("RobotId"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("Name"),
				KeyType:       types.KeyTypeRange,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("Robots"),
	}

	_, err := db.client.CreateTable(context.TODO(), param)
	if err != nil {
		log.Fatal("Cannot create table. " + err.Error())
	}
}

func (db *DBService) PutRobot(item map[string]types.AttributeValue) error {

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Robots"),
		Item:      item,
	}

	_, err := db.client.PutItem(db.ctx, input)
	return err
}

func (db *DBService) GetRobots(id string) map[string]string {
	robotsMap := db.GetRobotsForUser(id)
	return robotsMap
}
