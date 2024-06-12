package framework

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DBService struct {
	client *dynamodb.Client
	Robots []map[string]string
}

func GetDbService() *DBService {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	dbService := DBService{}
	dbService.client = dynamodb.NewFromConfig(cfg)

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

type RobotsDataGateway struct {
	Db *DBService
}

func (r RobotsDataGateway) GetRobots(id string) map[string]string {
	robotsMap := r.Db.GetRobotsForUser(id)
	return robotsMap
}
