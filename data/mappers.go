package data

import (
	application "example.com/on_path_robotics2/application"
)

func RobotFromMap(objData map[string]string) application.Robot {
	return application.Robot{
		Id:   objData["id"],
		Name: objData["name"],
	}
}
