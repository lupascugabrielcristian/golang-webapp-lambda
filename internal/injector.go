package internal

import (
	"example.com/on_path_robotics2/application"
	data "example.com/on_path_robotics2/data"
	presentation "example.com/on_path_robotics2/presentation"
)

func GetDBService() *presentation.DBService {
	// db := &framework.DBService{}
	// db.Robots = []map[string]string{
	// 	{"id": "abc1", "name": "Robot1"},
	// 	{"id": "abc2", "name": "Robot2"},
	// 	{"id": "abc3", "name": "Robot3"},
	// }
	// return db

	db := presentation.GetDbService()
	return db
}

func GetRobotsDataGateway() presentation.RobotsDataGateway {
	return presentation.RobotsDataGateway{Db: GetDBService()}
}

func GetGetRobotsRemote() data.GetRobotsRemote {
	return data.GetRobotsRemote{RemoteDataSource: GetRobotsDataGateway()}
}

func GetGetRobots() *application.GetRobots {
	return &application.GetRobots{Source: GetGetRobotsRemote()}
}

func GetCreateRobot() *application.CreateRobot {
	return &application.CreateRobot{}
}

func GetLambdaGateway() presentation.LambdaGateway {
	return presentation.LambdaGateway{GetRobotsUseCase: GetGetRobots(), CreateRobotUseCase: GetCreateRobot()}
}
