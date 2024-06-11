package main

import (
	"errors"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	internal "example.com/on_path_robotics2/internal"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	matched, err := regexp.Match(`.*/robots$`, []byte(request.Path))
	if err == nil && matched {
		return internal.GetLambdaGateway().HandleRequest(request)
	}

	matched, err = regexp.Match(`.*/get_robots$`, []byte(request.Path))
	if err == nil && matched {
		return internal.GetLambdaGateway().HandleRequest(request)
	}

	matched, err = regexp.Match(`.*/create_robot$`, []byte(request.Path))
	if err == nil && matched {
		return internal.GetLambdaGateway().HandleCreateRobotRequest(request)
	}

	return internal.GetLambdaGateway().GetInvalidRequestResponse(request), errors.New("path not defined. Request path is " + request.Path)
}
