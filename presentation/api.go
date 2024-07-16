package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type GetRobotsRequest struct {
	UserId    *string `json:"userId"`
	Source    *string `json:"source"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type CreateRobotRequest struct {
	Source *string `json:"source"`
	Name   *string `json:"name"`
}

func LambdaGatewayFactory(r *RobotsDelegate) *LambdaGateway {
	return &LambdaGateway{
		robotsDelegate: r,
	}
}

type LambdaGateway struct {
	robotsDelegate *RobotsDelegate
}

type ResponseBody struct {
	Message *string `json:"message"`
}

func (l LambdaGateway) GetInvalidRequestResponse(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       `{"error": "Invalid request"}`,
	}
}

func (l LambdaGateway) HandleGetRobotsRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody GetRobotsRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		fmt.Printf("Error parsing request body: %v", err)
		errorResp, err := generateErrorReponse()
		addHeaders(&errorResp, *requestBody.Source)
		return errorResp, err
	}

	robotData := l.robotsDelegate.GetRobots(requestBody)

	response, err := generateResponse(robotData)
	addHeaders(&response, *requestBody.Source)
	return response, err
}

func (l LambdaGateway) HandleCreateRobotRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody CreateRobotRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)

	if err != nil {
		fmt.Printf("Error parsing request body: %v", err)
		errorResp, err := generateErrorReponse()
		addHeaders(&errorResp, *requestBody.Source)
		return errorResp, err
	}

	robotData := l.robotsDelegate.CreateRoobot(requestBody)

	response, err := generateResponse(robotData)
	addHeaders(&response, *requestBody.Source)
	return response, err
}

func generateErrorReponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       `{"error": "Invalid request"}`,
	}, nil
}

func generateResponse(data map[string]string) (events.APIGatewayProxyResponse, error) {
	jbytes, _ := json.Marshal(data)
	jstr := string(jbytes)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       jstr,
	}

	return response, nil
}

func addHeaders(response *events.APIGatewayProxyResponse, source string) {
	hdrs := map[string]string{}

	if source == "A" {
		hdrs["Access-Control-Allow-Origin"] = "http://localhost:3000"
	} else if strings.Contains(source, "B") {
		hdrs["Access-Control-Allow-Origin"] = "https://master.d3cwzm2wqq04zv.amplifyapp.com" // From published app on prod stage
	} else if strings.Contains(source, "C") {
		hdrs["Access-Control-Allow-Origin"] = "https://8db35e87142744a9b114c7ba8978a032.vfs.cloud9.eu-central-1.amazonaws.com" // From Cloud9 env
	} else {
		hdrs["Access-Control-Allow-Origin"] = "http://localhost:3000"
	}

	response.Headers = hdrs
}
