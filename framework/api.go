package framework

import (
	"encoding/json"
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
	var requestBody GetRobotsRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		fmt.Printf("Error parsing request body: %v", err)
	}

	getRobotsDTO := dto.GetRobotsDTO{Id: *requestBody.FirstName, Name: *requestBody.LastName}
	robots := l.GetRobotsUseCase.Invoke(getRobotsDTO)
	fmt.Println(robots)
}

type GetRobotsRequest struct {
	Source    *string `json:"Source"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}
