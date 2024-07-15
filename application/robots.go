package application

import dto "example.com/on_path_robotics2/application/dto"

// DATA
type CreateRobotData struct {
	Name string
}

// SOURCES
type GetRobotsSource interface {
	GetRobots(id string) map[string]string
}

type CreateRobotSource interface {
	CreateRobot(data CreateRobotData)
}

// USE CASE
type GetRobots struct {
	Source GetRobotsSource
}

func (g GetRobots) Invoke(dto dto.GetRobotsDTO) map[string]string {
	return g.Source.GetRobots(dto.Id)
}

type CreateRobot struct {
	Source CreateRobotSource
}

func (g CreateRobot) Invoke(data CreateRobotData) bool {
	g.Source.CreateRobot(data)
	return true
}
