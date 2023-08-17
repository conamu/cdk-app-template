package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func GetPingLambda(stack constructs.Construct, id string) awslambda.Function {
	function := awslambda.NewFunction(stack, jsii.String(id), &awslambda.FunctionProps{
		FunctionName: s("Ping"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("internal/app/ping/bootstrap.zip"), nil),
		Handler:      jsii.String("bootstrap.zip"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	return function
}

func GetFunctionVersion(function awslambda.IFunction, stack constructs.Construct, env string) awslambda.Alias {
	lambdaVersion := awslambda.NewVersion(stack, s(env), &awslambda.VersionProps{
		Lambda: function,
	})

	lambdaAlias := awslambda.NewAlias(stack, s(env), &awslambda.AliasProps{
		AliasName: s("staging"),
		Version:   lambdaVersion,
	})

	return lambdaAlias
}
