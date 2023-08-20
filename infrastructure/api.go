package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type apiFunctionResource struct {
	apiPath            string
	apiMethod          string
	apiFunctionVersion awslambda.IVersion
}

func buildApiGateway(stack constructs.Construct, name string) awsapigateway.RestApi {
	if stage == "" {
		stage = "local"
	}

	api := awsapigateway.NewRestApi(stack, jsii.String(name), &awsapigateway.RestApiProps{
		RestApiName: jsii.String(name),
		Deploy:      jsii.Bool(false),
	})

	return api
}

func buildApiResources(stack constructs.Construct,
	api awsapigateway.RestApi, meta []*apiFunctionResource, requireApiKey bool, stage string) {

	resources := make(map[string]awsapigateway.Resource)

	// Build resources
	for _, data := range meta {
		if _, ok := resources[data.apiPath]; ok {
			continue
		}
		resources[data.apiPath] = awsapigateway.NewResource(stack, s(data.apiPath+"-"+stage), &awsapigateway.ResourceProps{
			Parent:   api.Root(),
			PathPart: s(data.apiPath),
		})
	}

	// Build methods
	for _, data := range meta {
		integration := awsapigateway.NewLambdaIntegration(data.apiFunctionVersion, nil)

		awsapigateway.NewMethod(stack,
			s(data.apiPath+"-"+data.apiMethod+"-"+stage),
			&awsapigateway.MethodProps{
				HttpMethod:  s(data.apiMethod),
				Resource:    resources[data.apiPath],
				Integration: integration,
				Options: &awsapigateway.MethodOptions{
					ApiKeyRequired: &requireApiKey,
				},
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

	apiKey := awsapigateway.NewApiKey(stack, s("default-key-"+stage), nil)

	awsapigateway.NewUsagePlan(stack, s("default-usage-plan-"+stage), &awsapigateway.UsagePlanProps{
		ApiStages: &usagePlanApiStages,
		Name:      s("default-plan-" + stage),
	}).AddApiKey(apiKey, nil)

	api.SetDeploymentStage(apiStage)
}
