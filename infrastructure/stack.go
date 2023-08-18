package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
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

	ApiGatewayRoot := GetApiGateway(stack,
		"transactions-api-"+stage,
		"Transaction API "+stage,
		"Api for transactions and orders")

	dep := awsapigateway.NewDeployment(stack, s("app-deployment-"+stage), &awsapigateway.DeploymentProps{
		Api: ApiGatewayRoot,
	})

	apiStage := awsapigateway.NewStage(stack, s(stage+"-stage"), &awsapigateway.StageProps{
		StageName:  s(stage),
		Deployment: dep,
	})

	// Set api gateway id for easier testing
	awscdk.Tags_Of(ApiGatewayRoot).Add(s("_custom_id_"), s("gofq6f9983"), &awscdk.TagProps{})

	PingLambda := GetPingLambda(stack, "ping-lambda-"+stage)

	NewPingLambdaVersion := awslambda.NewVersion(stack, s("ping-"+stage+"-version"), &awslambda.VersionProps{
		RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
		Lambda:        PingLambda,
	})

	NewPingLambdaVersion.GrantInvoke(awsiam.NewServicePrincipal(s("apigateway.amazonaws.com"), &awsiam.ServicePrincipalOpts{}))

	PingIntegration := awsapigateway.NewLambdaIntegration(NewPingLambdaVersion, &awsapigateway.LambdaIntegrationOptions{})

	ApiGatewayRoot.Root().
		AddResource(s("ping"), &awsapigateway.ResourceOptions{}).
		AddMethod(s("GET"), PingIntegration, &awsapigateway.MethodOptions{
			ApiKeyRequired: &requireApiKey,
		})

	var usagePlanApiStages []*awsapigateway.UsagePlanPerApiStage
	usagePlanApiStages = append(usagePlanApiStages, &awsapigateway.UsagePlanPerApiStage{
		Api:   ApiGatewayRoot,
		Stage: apiStage,
	})

	awsapigateway.NewUsagePlan(stack, s("default-usage-plan-"+stage), &awsapigateway.UsagePlanProps{
		ApiStages: &usagePlanApiStages,
		Name:      s("default-plan-" + stage),
	})

	ApiGatewayRoot.SetDeploymentStage(apiStage)

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
