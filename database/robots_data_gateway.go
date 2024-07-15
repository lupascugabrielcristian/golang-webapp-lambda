package datbase

import (
	"example.com/on_path_robotics2/application"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type RobotsDataGateway struct {
	Db *DBService
}

func (gateway *RobotsDataGateway) CreateRobot(r application.Robot) bool {

	item, err := attributevalue.MarshalMap(r) // map[string]types.AttributeValue

	if err != nil {
		return false
	}

	gateway.Db.PutRobot(item)

	return true
}
