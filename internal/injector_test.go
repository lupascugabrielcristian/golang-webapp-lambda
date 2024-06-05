package internal

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestLambdaGateway(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	GetLambdaGateway().HandleRequest(request)
}
