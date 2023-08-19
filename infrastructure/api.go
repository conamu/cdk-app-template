package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type apiMetadata struct {
	apiPath   string
	apiMethod string
}

func buildApiGateway(stack constructs.Construct, id, name, description string) awsapigateway.RestApi {
	if stage == "" {
		stage = "local"
	}

	api := awsapigateway.NewRestApi(stack, jsii.String(id), &awsapigateway.RestApiProps{
		Description: jsii.String(description),
		RestApiName: jsii.String(name),
		Deploy:      jsii.Bool(false),
	})

	return api
}

func buildApiResources(stack constructs.Construct,
	api awsapigateway.RestApi, meta []*apiMetadata,
	versions []awslambda.IVersion, requireApiKey bool) {

	for idx, ver := range versions {
		integration := awsapigateway.NewLambdaIntegration(ver, nil)

		api.Root().
			AddResource(s(meta[idx].apiPath), nil).
			AddMethod(s(meta[idx].apiMethod), integration, &awsapigateway.MethodOptions{
				ApiKeyRequired: &requireApiKey,
			})
	}

	dep := awsapigateway.NewDeployment(stack, s("app-deployment-"+stage), &awsapigateway.DeploymentProps{
		Api: api,
	})

	apiStage := awsapigateway.NewStage(stack, s(stage+"-stage"), &awsapigateway.StageProps{
		StageName:  s(stage),
		Deployment: dep,
	})

	var usagePlanApiStages []*awsapigateway.UsagePlanPerApiStage
	usagePlanApiStages = append(usagePlanApiStages, &awsapigateway.UsagePlanPerApiStage{
		Api:   api,
		Stage: apiStage,
	})

	awsapigateway.NewUsagePlan(stack, s("default-usage-plan-"+stage), &awsapigateway.UsagePlanProps{
		ApiStages: &usagePlanApiStages,
		Name:      s("default-plan-" + stage),
	})

	api.SetDeploymentStage(apiStage)
}
