package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var logger = log.New(os.Stdout, "OUT ", log.Llongfile)

type TOut struct {
	Message string
}

func Handle(req *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	msg := TOut{Message: "pong"}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bytes),
	}
	return &res, nil
}

func main() {
	lambda.Start(Handle)
}