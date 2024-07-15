package presentation

import (
	"fmt"

	application "example.com/on_path_robotics2/application"
	dto "example.com/on_path_robotics2/application/dto"
)

type RobotsDelegate struct {
	getRobotsUseCase   application.GetRobots
	createRobotUseCase application.CreateRobot
}

func (d *RobotsDelegate) CreateRoobot(request CreateRobotRequest) map[string]string {
	data := application.CreateRobotData{
		Name: *request.Name,
	}
	success := d.createRobotUseCase.Invoke(data)

	if success {
		return map[string]string{
			"result": "ok",
		}
	} else {
		return map[string]string{
			"result": "nok",
		}
	}
}

func (d *RobotsDelegate) GetRobots(dto dto.GetRobotsDTO) map[string]string {
	return map[string]string{
		"RobotId": fmt.Sprintf("to implement %s", dto.Id),
		"Name":    fmt.Sprintf("from db %s", dto.Name),
	}
}
