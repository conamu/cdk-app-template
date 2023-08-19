package infrastructure

import (
	"cdk-app-template/internal/pkg/logger"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
	"strings"
)

func getLambdas(stack constructs.Construct, stage string) ([]awslambda.IFunction, []*apiMetadata) {
	log, _ := logger.Create()
	dirs, err := os.ReadDir("internal/app/lambda")
	if err != nil {
		panic(err)
	}

	var functions []awslambda.IFunction
	var apiMeta []*apiMetadata

	for _, dir := range dirs {
		log.Info("Building Lambda: " + dir.Name())

		dataStrings := strings.Split(dir.Name(), "-")
		name := dataStrings[0]
		method := dataStrings[1]

		md := &apiMetadata{
			apiPath:   name,
			apiMethod: strings.ToUpper(method),
		}

		function := buildLambda(stack, dir.Name(), stage)
		functions = append(functions, function)
		apiMeta = append(apiMeta, md)
	}
	return functions, apiMeta
}

func buildLambda(stack constructs.Construct, path, stage string) awslambda.IFunction {
	function := awslambda.NewFunction(stack, jsii.String(path+"-lambda-"+stage), &awslambda.FunctionProps{
		FunctionName: s(path + "-" + stage),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("internal/app/lambda/"+path+"/bootstrap.zip"), nil),
		Handler:      jsii.String("bootstrap.zip"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		CurrentVersionOptions: &awslambda.VersionOptions{
			RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
		},
	})

	return function
}

func buildLambdaVersions(stack constructs.Construct, functions []awslambda.IFunction) []awslambda.IVersion {
	var lambdaVersions []awslambda.IVersion

	for _, function := range functions {
		version := awslambda.NewVersion(stack, s("ping-"+stage+"-version"), &awslambda.VersionProps{
			RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
			Lambda:        function,
		})
		lambdaVersions = append(lambdaVersions, version)
	}
	return lambdaVersions
}
