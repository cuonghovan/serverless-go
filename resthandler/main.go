package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the entry point called for each Lambda event
func Handler(ctx context.Context, req *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.Resource {
	case "/hello":
		switch req.HTTPMethod {
		case http.MethodGet:
			return getHelloMessage(ctx, req)
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
