package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/spf13/viper"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type stackProps struct {
	awscdk.StackProps
}

func newStack(scope constructs.Construct, id string, props *stackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	return stack
}

func BuildStack() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	stack := newStack(app, "CdkAppStack", &stackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	ApiGatewayRoot := GetApiGateway(stack,
		"transactions-api",
		"Transaction API",
		"Api for transactions and orders")

	// Set api gateway id for easier testing
	awscdk.Tags_Of(ApiGatewayRoot).Add(s("_custom_id_"), s("gofq6f9983"), &awscdk.TagProps{})

	PingLambda := GetPingLambda(stack, "ping-lambda")

	GetDynamoDb(stack, "customer-table")

	PingIntegration := awsapigateway.NewLambdaIntegration(PingLambda, &awsapigateway.LambdaIntegrationOptions{})

	ApiGatewayRoot.Root().
		AddResource(s("ping"), &awsapigateway.ResourceOptions{}).
		AddMethod(s("GET"), PingIntegration, &awsapigateway.MethodOptions{})

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
		Account: s(viper.GetString("aws-account")),
		Region:  s(viper.GetString("aws-region")),
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

func s(s string) *string {
	return jsii.String(s)
}
