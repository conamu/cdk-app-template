package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/spf13/viper"
	"os"
	"regexp"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type stackProps struct {
	awscdk.StackProps
}

var stage string

func BuildStack() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	stage = os.Getenv("ENV")

	// This is necessary to be able to use git branch names in cloudformation stacks
	stage = removeNumbersAndSpecialChars(stage)

	requireApiKey := true

	if stage != "production" && stage != "staging" {
		requireApiKey = false
	}

	stack := newStack(app, "CdkAppStack-"+stage, &stackProps{
		awscdk.StackProps{
			StackName: s("Cdk-App-Template-" + stage),
			Env:       env(),
		},
	})

	lambdaApiMeta := getLambdas(stack, stage)

	// Grant permissions to api gateway to invoke functions
	for _, meta := range lambdaApiMeta {
		meta.apiFunctionVersion.GrantInvoke(awsiam.NewServicePrincipal(s("apigateway.amazonaws.com"), &awsiam.ServicePrincipalOpts{}))
	}
	ApiGatewayRoot := buildApiGateway(stack,
		"transactions-api-"+stage,
		"Transaction API "+stage,
		"Api for transactions and orders")

	buildApiResources(stack, ApiGatewayRoot, lambdaApiMeta, requireApiKey, stage)

	// Set api gateway id for easier testing
	awscdk.Tags_Of(ApiGatewayRoot).Add(s("_custom_id_"), s("gofq6f9983"), &awscdk.TagProps{})

	awscdk.NewCfnOutput(stack, s("api-url"), &awscdk.CfnOutputProps{
		Value: ApiGatewayRoot.Url(),
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	acc := viper.GetString("aws-account")
	reg := viper.GetString("aws-region")
	return &awscdk.Environment{
		Account: s(acc),
		Region:  s(reg),
	}
}

func s(s string) *string {
	return jsii.String(s)
}

func newStack(scope constructs.Construct, id string, props *stackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	return stack
}

func removeNumbersAndSpecialChars(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z]+")
	return reg.ReplaceAllString(input, "")
}
