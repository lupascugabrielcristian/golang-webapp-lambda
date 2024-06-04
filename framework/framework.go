package framework

import (
	"fmt"

	framework "example.com/on_path_robotics2/framework/dto"
	"github.com/aws/aws-lambda-go/events"
)

func HandleGetRobotsRequest(request events.APIGatewayProxyRequest) {
	fmt.Println("Handling GetRobots request")

	dto := new(framework.GetRobotsDTO)
	dto.Id = "request id"
}
