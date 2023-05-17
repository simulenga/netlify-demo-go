package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Get information from the data payload.
	var body Payload
	json.Unmarshal([]byte(request.Body), &body)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("%s", body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
