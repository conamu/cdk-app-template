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

func getLambdas(stack constructs.Construct, stage string) []*apiResourceMetadata {
	log, _ := logger.Create()
	dirs, err := os.ReadDir("internal/app/lambda")
	if err != nil {
		panic(err)
	}

	var apiMeta []*apiResourceMetadata

	for _, dir := range dirs {
		log.Info("Building Lambda: " + dir.Name())

		dataStrings := strings.Split(dir.Name(), "-")
		name := dataStrings[0]
		method := dataStrings[1]

		function, functionName := buildLambda(stack, dir.Name(), stage)
		version := buildLambdaVersion(stack, function, functionName)

		md := &apiResourceMetadata{
			apiPath:            name,
			apiMethod:          strings.ToUpper(method),
			apiFunctionVersion: version,
		}

		apiMeta = append(apiMeta, md)
	}
	return apiMeta
}

func buildLambda(stack constructs.Construct, path, stage string) (awslambda.IFunction, string) {
	name := path + "-" + stage
	function := awslambda.NewFunction(stack, jsii.String(path+"-lambda-"+stage), &awslambda.FunctionProps{
		FunctionName: &name,
		Code:         awslambda.AssetCode_FromAsset(jsii.String("internal/app/lambda/"+path+"/bootstrap.zip"), nil),
		Handler:      jsii.String("bootstrap.zip"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		CurrentVersionOptions: &awslambda.VersionOptions{
			RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
		},
	})

	return function, name
}

func buildLambdaVersion(stack constructs.Construct, function awslambda.IFunction, name string) awslambda.IVersion {

	version := awslambda.NewVersion(stack, s(name+"-version"), &awslambda.VersionProps{
		RemovalPolicy: awscdk.RemovalPolicy_RETAIN,
		Lambda:        function,
	})
	return version
}
