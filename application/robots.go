package application

import dto "example.com/on_path_robotics2/application/dto"

// SOURCES
type GetRobotsSource interface {
	GetRobots(id string) map[string]string
}

// USE CASE
type GetRobots struct {
	Source GetRobotsSource
}

func (g GetRobots) Invoke(dto dto.GetRobotsDTO) map[string]string {
	return g.Source.GetRobots(dto.Id)
}
