package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var person Person
	fmt.Println("urmeaza person")
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Println("am luat person")

	msg := fmt.Sprintf("Hello from %v %v by lambda fc", *person.FirstName, *person.LastName)
	responseBody := ResponseBody{
		Message: &msg,
	}

	jbytes, _ := json.Marshal(responseBody)
	jstr := string(jbytes)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       jstr,
	}
	response.Headers["Access-Control-Allow-Origin"] = "*"
	return response, nil
}

type ResponseBody struct {
	Message *string `json:"message"`
}

type Person struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}
