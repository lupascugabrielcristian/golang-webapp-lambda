package main

import (
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandlerFunction_for_prod_robots(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Path = "https://qkb8myc9t1.execute-api.eu-central-1.amazonaws.com/prod/robots"
	request.Body = `{"source": "B", "firstName": "abc1", "lastName": "NU stiu"}`
	_, err := handler(request)

	if err != nil {
		t.Fatal("Should not get an error")
	}
}

func TestHandlerFunction_should_not_pass_for_invalid_path(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Path = "https://qkb8myc9t1.execute-api.eu-central-1.amazonaws.com/unknown/robots"
	request.Body = `{"source": "B", "firstName": "abc1", "lastName": "NU stiu"}`
	_, err := handler(request)

	if err == nil {
		t.Fatal("Should not get an error")
	}
}

func TestHandlerFunction_for_prod_get_robots(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Path = "https://qkb8myc9t1.execute-api.eu-central-1.amazonaws.com/prod/get_robots"
	request.Body = `{"source": "B", "firstName": "abc1", "lastName": "NU stiu"}`
	_, err := handler(request)

	if err != nil {
		t.Fatal("Should not get an error")
	}
}

func TestHandlerFunction_for_prod_create_robot(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Path = "https://qkb8myc9t1.execute-api.eu-central-1.amazonaws.com/prod/create_robot"
	request.Body = `{"source": "B", "name": "abc1"}`
	resp, err := handler(request)

	if err != nil {
		t.Fatal("Should not get an error")
	}

	if resp.Headers["Access-Control-Allow-Origin"] != "https://master.d3cwzm2wqq04zv.amplifyapp.com" {
		t.Fatal("Incorrect Allow-Origin header")
	}

	if !strings.Contains(resp.Body, "robot with name abc1") {
		t.Fatalf("Incorrect response: %s", resp.Body)
	}
}
