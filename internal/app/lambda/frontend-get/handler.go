package main

import (
	"cdk-app-template/internal/pkg/constants"
	"cdk-app-template/internal/pkg/logger"
	"context"
	"github.com/aws/aws-lambda-go/events"
)

const html = `
<html>
	<h1>Hello World</h1>
	<form hx-get="/staging/enhance" hx-target=#content>
		<input id="name" name="name" type="text">
		<button type="submit">Submit</button>
	</form>
	<button hx-get=/staging/enhance?name=constantin hx-target=#content>Test!</button>
	<div id=content></div>
</html>
<script src="https://unpkg.com/htmx.org@1.9.5" integrity="sha384-xcuj3WpfgjlKF+FXhSQFQ0ZNr39ln+hwjN3npfM9VBnUskLolQAcN80McRIVOPuO" crossorigin="anonymous"></script>
`

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	log := logger.Create()
	ctx = context.WithValue(ctx, constants.CTX_LOGGER, log)

	log.Info("Frontend endpoint hit!")

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		MultiValueHeaders: nil,
		Body:              html,
		IsBase64Encoded:   false,
	}, nil
}
