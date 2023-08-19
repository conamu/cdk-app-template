package main

import (
	"cdk-app-template/internal/pkg/logger"
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	log, _ := logger.Create()
	log.Info("Ping endpoint hit!")

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body: `{
								"message":"Hello World!",
								"bodyCount": 16
							}`,
		IsBase64Encoded: false,
	}, nil
}
