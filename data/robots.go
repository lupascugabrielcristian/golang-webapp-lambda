package data

import (
	application "example.com/on_path_robotics2/application"
)

// SOURCES
type RobotsDataSourceRemote interface {
	GetRobots(id string) map[string]string
}

// IMPLEMENTATIONS
type GetRobotsRemote struct {
	RemoteDataSource RobotsDataSourceRemote
}

func (g GetRobotsRemote) GetRobots() application.Robot {
	objData := g.RemoteDataSource.GetRobots("userId")
	return RobotFromMap(objData)
}
