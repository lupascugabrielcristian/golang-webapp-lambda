package persistance

import application "example.com/on_path_robotics2/application"

type RobotsDataSource interface {
	CreateRobot(r application.Robot) bool
	GetRobots(userId *string) []application.Robot
}
