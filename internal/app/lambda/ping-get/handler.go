package main

import (
	"cdk-app-template/internal/pkg/constants"
	"cdk-app-template/internal/pkg/logger"
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	log := logger.Create()
	ctx = context.WithValue(ctx, constants.CTX_LOGGER, log)

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
