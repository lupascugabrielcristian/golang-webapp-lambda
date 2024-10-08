package database

import (
	application "example.com/on_path_robotics2/application"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/google/uuid"
)

type RobotsDataGateway struct {
	Db *DBService
}

func RobotsDataGatewayFactory(db *DBService) *RobotsDataGateway {
	return &RobotsDataGateway{Db: db}
}

func (gateway *RobotsDataGateway) CreateRobot(r application.Robot) bool {
	r.RobotId = uuid.New().String()

	item, err := attributevalue.MarshalMap(r) // map[string]types.AttributeValue

	if err != nil {
		return false
	}

	gateway.Db.PutRobot(item)

	return true
}

func (gateway *RobotsDataGateway) GetRobots(userId *string) []application.Robot {
	itemsOutput, error := gateway.Db.GetRobots(userId)

	robots := []application.Robot{}

	if error != nil {
		return robots
	}

	error = attributevalue.UnmarshalListOfMaps(itemsOutput, &robots)

	if error != nil {
		return robots
	}

	return robots
}
