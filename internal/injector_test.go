package internal

import (
	"fmt"
	"strings"
	"testing"

	application "example.com/on_path_robotics2/application"

	"github.com/aws/aws-lambda-go/events"
)

func TestLambdaGateway(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Body = `{"source": "B", "firstName": "abc1", "lastName": "NU stiu"}`
	response, err := GetLambdaGateway().HandleRequest(request)

	if err != nil {
		t.Fatal("Should not get an error")
	}

	// Checking that Allow-Origin header is added
	if response.Headers["Access-Control-Allow-Origin"] != "https://master.d3cwzm2wqq04zv.amplifyapp.com" {
		t.Log(response.Headers["Access-Control-Allow-Origin"])
		t.Fatal("Incorrect header")
	}

	// Checking that data is returned
	if !strings.Contains(response.Body, "abc1") {
		t.Fatal("Incorrect response")
	}

	if !strings.Contains(response.Body, "Robot1") {
		t.Fatal("Incorrect response")
	}
}

func TestCreateTable(t *testing.T) {
	dbService := GetDBService()

	// Should create the client
	dbService.CreateTables()
}

func TestAddDocumentToRobotsTable(t *testing.T) {
	dbService := GetDBService()
	robot := application.Robot{
		RobotId: "someid2",
		Name:    "Some Name 2",
	}

	// Should add the new document to the Robots table
	err := dbService.PutRobot(robot)
	fmt.Println(err)
}
