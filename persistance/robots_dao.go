package persistance

import (
	application "example.com/on_path_robotics2/application"
)

// GET ROBOTS
type GetRobotsDAO struct {
	datasource RobotsDataSource
}

func GetRobotsDAOFactory(ds RobotsDataSource) *GetRobotsDAO {
	return &GetRobotsDAO{
		datasource: ds,
	}
}

func (dao *GetRobotsDAO) GetRobots(userId *string) []application.Robot {
	return dao.datasource.GetRobots(userId)
}

// CREATE ROBOT
type CreateRobotDAO struct {
	datasource RobotsDataSource
}

func CreateRobotDAOFactory(ds RobotsDataSource) *CreateRobotDAO {
	return &CreateRobotDAO{
		datasource: ds,
	}
}

func (dao *CreateRobotDAO) CreateRobot(data application.CreateRobotData) error {
	robot := application.Robot{
		RobotId: "",
		Name:    data.Name,
	}

	dao.datasource.CreateRobot(robot)
	return nil
}
