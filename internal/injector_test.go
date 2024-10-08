package internal

import (
	"fmt"
	"strings"
	"testing"

	application "example.com/on_path_robotics2/application"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestLambdaGateway(t *testing.T) {
	request := events.APIGatewayProxyRequest{}
	request.Body = `{"source": "B", "firstName": "abc1", "lastName": "NU stiu"}`
	response, err := GetLambdaGateway().HandleGetRobotsRequest(request)

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
	dbService := getDBService()

	// Should create the client
	dbService.CreateTables()
}

func TestAddDocumentToRobotsTable(t *testing.T) {
	dbService := getDBService()
	robot := map[string]types.AttributeValue{
		"RobotId": &types.AttributeValueMemberS{Value: "Some id"},
		"Name":    &types.AttributeValueMemberS{Value: "Cristi 1"},
	}

	// Should add the new document to the Robots table
	err := dbService.PutRobot(robot)
	fmt.Println(err)
}

func TestGetRobotsUseCase(t *testing.T) {
	userId := "some id"
	getRobots := getGetRobotsUseCase()

	robots := getRobots.Invoke(&userId)

	if len(robots) != 3 {
		t.Fatal("Should be 3 items in database")
	}
}

func TestCreateRobotUseCase(t *testing.T) {
	data := application.CreateRobotData{
		Name: "Robot 1",
	}
	createRobot := getCreateRobotsUseCase()

	res := createRobot.Invoke(data)

	if res != true {
		t.Fatal("Not true")
	}
}
