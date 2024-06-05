package internal

import (
	"fmt"

	"example.com/on_path_robotics2/application"
	dto "example.com/on_path_robotics2/application/dto"
	data "example.com/on_path_robotics2/data"
	framework "example.com/on_path_robotics2/framework"
)

func GetUseCase(dto dto.GetRobotsDTO) {
	fmt.Println("Use case obtained from injector file")
}

func GetComponent() {
	fmt.Println("Component obtained from injector file")
}

func GetGetRobotsRemote() data.GetRobotsRemote {
	return data.GetRobotsRemote{}
}

func GetGetRobots() application.GetRobots {
	return application.GetRobots{Source: GetGetRobotsRemote()}
}

func GetLambdaGateway() framework.LambdaGateway {
	return framework.LambdaGateway{GetRobotsUseCase: GetGetRobots()}
}
