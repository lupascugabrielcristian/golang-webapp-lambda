package database

import (
	"context"
	"fmt"
	"log"
	"os"

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
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatal(err)
	}

	provider := cfg.Credentials
	cred, _ := provider.Retrieve(ctx)
	fmt.Fprintln(os.Stdout, "Source: "+cred.Source)
	fmt.Fprintln(os.Stdout, "AccessKeyID: "+cred.AccessKeyID)

	dbService := DBService{}
	dbService.client = dynamodb.NewFromConfig(cfg)
	dbService.ctx = ctx

	return &dbService
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

func (db *DBService) GetRobots(id *string) ([]map[string]types.AttributeValue, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Robots"),
	}

	output, err := db.client.Scan(db.ctx, input)

	return output.Items, err
}
