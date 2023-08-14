package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"payment-and-order-provider-template/infrastructure"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PaymentAndOrderProviderTemplateStackProps struct {
	awscdk.StackProps
}

func NewPaymentAndOrderProviderTemplateStack(scope constructs.Construct, id string, props *PaymentAndOrderProviderTemplateStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("PaymentAndOrderProviderTemplateQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stack := NewPaymentAndOrderProviderTemplateStack(app, "PaymentAndOrderProviderTemplateStack", &PaymentAndOrderProviderTemplateStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	ApiGatewayRoot := infrastructure.GetApiGateway(stack,
		"transactions-api",
		"Transaction API",
		"Api for transactions and orders")

	// Set api gateway id for easier testing
	awscdk.Tags_Of(ApiGatewayRoot).Add(str("_custom_id_"), str("gofq6f9983"), &awscdk.TagProps{})

	PingLambda := infrastructure.GetPingLambda(stack, "ping-lambda")

	infrastructure.GetDynamoDb(stack, "customer-table")

	PingIntegration := awsapigateway.NewLambdaIntegration(PingLambda, &awsapigateway.LambdaIntegrationOptions{})

	ApiGatewayRoot.Root().
		AddResource(jsii.String("ping"), &awsapigateway.ResourceOptions{}).
		AddMethod(jsii.String("GET"), PingIntegration, &awsapigateway.MethodOptions{})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String("000000000000"),
		Region:  jsii.String("eu-west-1"),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}

func str(s string) *string {
	return jsii.String(s)
}
