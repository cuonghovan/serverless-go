package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jirnexu/sora-go/util"
)

func getHelloMessage(ctx context.Context, req *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Headers:    util.WriteHeaders(),
		StatusCode: http.StatusOK,
		Body:       string("hello world!"),
	}, nil
}
