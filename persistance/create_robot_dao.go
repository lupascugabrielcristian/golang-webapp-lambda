package persistance

import (
	application "example.com/on_path_robotics2/application"
)

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
		RobotId: "to not add here",
		Name:    data.Name,
	}

	dao.datasource.CreateRobot(robot)
	return nil
}
