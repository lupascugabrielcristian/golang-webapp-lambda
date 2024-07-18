package presentation

import (
	"testing"

	application "example.com/on_path_robotics2/application"
)

type GetRobotsSourceMock struct {
}

func (gr *GetRobotsSourceMock) GetRobots(userId *string) []application.Robot {
	return []application.Robot{
		{
			RobotId: "id1",
			Name:    "Robot 1",
		},
		{
			RobotId: "id2",
			Name:    "Robot 2",
		},
	}
}

type CreateRobotSourceMock struct {
}

func (cr *CreateRobotSourceMock) CreateRobot(data application.CreateRobotData) error {
	return nil
}

func TestGetRobots(t *testing.T) {
	getRobotsUseCase := application.GetRobots{Source: &GetRobotsSourceMock{}}
	createRobotsUseCase := application.CreateRobot{Source: &CreateRobotSourceMock{}}

	userId := "some user"
	request := GetRobotsRequest{
		UserId: &userId,
	}

	robotsDelegate := RobotsDelegateFactory(&getRobotsUseCase, &createRobotsUseCase)
	robotsJson := robotsDelegate.GetRobots(request)

	if string(robotsJson) != "[{\"RobotId\":\"id1\",\"Name\":\"Robot 1\"},{\"RobotId\":\"id2\",\"Name\":\"Robot 2\"}]" {
		t.Fail()
	}
}
