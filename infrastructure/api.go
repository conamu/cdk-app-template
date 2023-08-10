package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func GetApiGateway(stack constructs.Construct, id, name, description string) awsapigateway.RestApi {
	api := awsapigateway.NewRestApi(stack, jsii.String(id), &awsapigateway.RestApiProps{
		Description: jsii.String(description),
		RestApiName: jsii.String(name),
	})
	return api
}
