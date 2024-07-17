package presentation

import (
	application "example.com/on_path_robotics2/application"
)

func RobotsDelegateFactory(gR *application.GetRobots, cR *application.CreateRobot) *RobotsDelegate {
	return &RobotsDelegate{
		getRobotsUseCase:   gR,
		createRobotUseCase: cR,
	}
}

type RobotsDelegate struct {
	getRobotsUseCase   *application.GetRobots
	createRobotUseCase *application.CreateRobot
}

func (d *RobotsDelegate) CreateRobot(request CreateRobotRequest) map[string]string {
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

func (d *RobotsDelegate) GetRobots(request GetRobotsRequest) map[string]string {
	// Aici returnez []Robot
	_ = d.getRobotsUseCase.Invoke(request.UserId)
	return map[string]string{
		"result": "ok",
	}
}
