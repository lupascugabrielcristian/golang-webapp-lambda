package framework

import (
	"fmt"

	application "example.com/on_path_robotics2/application"
	dto "example.com/on_path_robotics2/application/dto"
	"github.com/aws/aws-lambda-go/events"
)

func HandleGetRobotsRequest(request events.APIGatewayProxyRequest) {
	fmt.Println("Handling GetRobots request")

	dto := new(dto.GetRobotsDTO)
	dto.Id = "request id"
}

type LambdaGateway struct {
	GetRobotsUseCase application.GetRobots
}

func (l LambdaGateway) HandleRequest(request events.APIGatewayProxyRequest) {
	grDTO := dto.GetRobotsDTO{Id: "id"}
	robots := l.GetRobotsUseCase.Invoke(grDTO)
	fmt.Println(robots)
}
