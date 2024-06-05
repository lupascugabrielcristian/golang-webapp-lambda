package internal

import (
	"example.com/on_path_robotics2/application"
	data "example.com/on_path_robotics2/data"
	framework "example.com/on_path_robotics2/framework"
)

func GetDBService() *framework.DBService {
	return &framework.DBService{}
}

func GetRobotsDataGateway() framework.RobotsDataGateway {
	return framework.RobotsDataGateway{Db: GetDBService()}
}

func GetGetRobotsRemote() data.GetRobotsRemote {
	return data.GetRobotsRemote{RemoteDataSource: GetRobotsDataGateway()}
}

func GetGetRobots() application.GetRobots {
	return application.GetRobots{Source: GetGetRobotsRemote()}
}

func GetLambdaGateway() framework.LambdaGateway {
	return framework.LambdaGateway{GetRobotsUseCase: GetGetRobots()}
}
