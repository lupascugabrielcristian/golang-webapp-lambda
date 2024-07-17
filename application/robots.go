package application

// DATA
type CreateRobotData struct {
	Name string
}

// SOURCES
type GetRobotsSource interface {
	GetRobots(userId *string) []Robot
}

type CreateRobotSource interface {
	CreateRobot(data CreateRobotData) error
}

// USE CASE
type GetRobots struct {
	Source GetRobotsSource
}

func (g *GetRobots) Invoke(userId *string) []Robot {
	return g.Source.GetRobots(userId)
}

type CreateRobot struct {
	Source CreateRobotSource
}

func (g CreateRobot) Invoke(data CreateRobotData) bool {
	g.Source.CreateRobot(data)
	return true
}
