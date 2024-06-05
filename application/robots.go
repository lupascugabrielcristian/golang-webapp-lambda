package application

import dto "example.com/on_path_robotics2/application/dto"

// SOURCES
type GetRobotsSource interface {
	GetRobots() []string
}

// USE CASE
type GetRobots struct {
	Source GetRobotsSource
}

func (g GetRobots) Invoke(dto dto.GetRobotsDTO) []string {
	return g.Source.GetRobots()
}
