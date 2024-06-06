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

func (g GetRobotsRemote) GetRobots(id string) application.Robot {
	objData := g.RemoteDataSource.GetRobots(id)
	return RobotFromMap(objData)
}
