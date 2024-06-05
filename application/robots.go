package application

import dto "example.com/on_path_robotics2/application/dto"

// SOURCES
type GetRobotsSource interface {
	GetRobots() Robot
}

// USE CASE
type GetRobots struct {
	Source GetRobotsSource
}

func (g GetRobots) Invoke(dto dto.GetRobotsDTO) Robot {
	return g.Source.GetRobots()
}
