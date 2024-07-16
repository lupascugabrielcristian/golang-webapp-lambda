package internal

import (
	"example.com/on_path_robotics2/application"
	database "example.com/on_path_robotics2/database"
	"example.com/on_path_robotics2/persistance"
	presentation "example.com/on_path_robotics2/presentation"
)

func getDBService() *database.DBService {
	db := database.GetDbService()
	return db
}

func getRobotsDataGatewayFactory() *database.RobotsDataGateway {
	return database.RobotsDataGatewayFactory(getDBService())
}

func getGetRobotsDAO() *persistance.GetRobotsDAO {
	return persistance.GetRobotsDAOFactory(getRobotsDataGatewayFactory())
}

func getGetRobotsUseCase() *application.GetRobots {
	return &application.GetRobots{Source: getGetRobotsDAO()}
}

func getCreateRobotDAO() *persistance.CreateRobotDAO {
	return persistance.CreateRobotDAOFactory(getRobotsDataGatewayFactory())
}

func getCreateRobotsUseCase() *application.CreateRobot {
	return &application.CreateRobot{Source: getCreateRobotDAO()}
}

func getRobotsDelegate() *presentation.RobotsDelegate {
	return presentation.RobotsDelegateFactory(getGetRobotsUseCase(), getCreateRobotsUseCase())
}

func GetLambdaGateway() *presentation.LambdaGateway {
	return presentation.LambdaGatewayFactory(getRobotsDelegate())
}
