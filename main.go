package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
		fmt.Printf("Error parsing request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Invalid request"}`,
		}, nil
	}

	fmt.Println("am luat person")

	msg := fmt.Sprintf("Hello from %v %v by lambda fc", *person.FirstName, *person.LastName)
	responseBody := ResponseBody{
		Message: &msg,
	}

	jbytes, _ := json.Marshal(responseBody)
	jstr := string(jbytes)

	hdrs := map[string]string{}

	if *person.Source == "A" {
		hdrs["Access-Control-Allow-Origin"] = "http://localhost:3000"
	} else if strings.Contains(*person.Source, "B") {
		hdrs["Access-Control-Allow-Origin"] = "https://master.d3cwzm2wqq04zv.amplifyapp.com" // From published app on prod stage
	} else if strings.Contains(*person.Source, "C") {
		hdrs["Access-Control-Allow-Origin"] = "https://8db35e87142744a9b114c7ba8978a032.vfs.cloud9.eu-central-1.amazonaws.com" // From Cloud9 env
	} else {
		hdrs["Access-Control-Allow-Origin"] = "http://localhost:3000"
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       jstr,
		Headers:    hdrs,
	}

	// response.Headers["Access-Control-Allow-Origin"] = "*"
	return response, nil
}

type ResponseBody struct {
	Message *string `json:"message"`
}

type Person struct {
	Source    *string `json:"Source"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}
