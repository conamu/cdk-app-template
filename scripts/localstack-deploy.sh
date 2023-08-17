export ENV=local
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export CDK_DEPLOY_ACCOUNT=000000000000
export CDK_DEPLOY_REGION=eu-west-1
export AWS_DEFAULT_REGION=eu-west-1
cdklocal synth
cdklocal bootstrap -v
cdklocal deploy --require-approval never -v