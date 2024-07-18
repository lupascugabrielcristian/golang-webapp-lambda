package presentation

import (
	"encoding/json"

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

func (d *RobotsDelegate) CreateRobot(request CreateRobotRequest) []byte {
	data := application.CreateRobotData{
		Name: *request.Name,
	}
	success := d.createRobotUseCase.Invoke(data)

	if success {
		return []byte(`{"result": "ok"}`)
	} else {
		return []byte(`{"result": "nok"}`)
	}
}

func (d *RobotsDelegate) GetRobots(request GetRobotsRequest) []byte {
	robots := d.getRobotsUseCase.Invoke(request.UserId)

	jsonData, err := json.Marshal(robots)
	if err != nil {
		return []byte(`{"result": "nok"}`)
	} else {
		return jsonData
	}
}
