package internal

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestLambdaGateway(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Body = `{"source": "B", "firstName": "abc1", "lastName": "NU stiu"}`
	response, err := GetLambdaGateway().HandleRequest(request)

	if err != nil {
		t.Fatal("Should not get an error")
	}

	if response.Headers["Access-Control-Allow-Origin"] != "https://master.d3cwzm2wqq04zv.amplifyapp.com" {
		t.Log(response.Headers["Access-Control-Allow-Origin"])
		t.Fatal("Incorrect header")
	}
}
