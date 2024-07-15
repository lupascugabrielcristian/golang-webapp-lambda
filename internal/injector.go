package internal

import (
	"example.com/on_path_robotics2/application"
	data "example.com/on_path_robotics2/data"
	persistance "example.com/on_path_robotics2/persistance"
	database "example.com/on_path_robotics2/database"
	presentation "example.com/on_path_robotics2/presentation"
)

func GetDBService() *database.DBService {
	db := database.GetDbService()
	return db
}

func GetRobotsDataGateway() database.RobotsDataGateway {
	return database.RobotsDataGateway{Db: GetDBService()}
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

func getGetRobotsUseCase() *persistance.CreateRobotDAO {
	return &persistance.CreateRobotDAO{}
}

func getRobotsDelegate() *presentation.RobotsDelegate {
	return presentation.RobotsDelegateFactory(g: nil, c: getGetRobotsUseCase())
}

func GetLambdaGateway() *presentation.LambdaGateway {
	return presentation.LambdaGatewayFactory(getRobotsDelegate())
}
