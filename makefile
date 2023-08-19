deploy-staging: build
	- ./scripts/deploy.sh

local-deploy: build
	- rm -rd cdk.out
	- ./scripts/localstack-deploy.sh

localstack-up:
	- docker-compose -f localstack.compose up -d

localstack-down:
	- docker-compose -f localstack.compose down

build:
	- ./scripts/build.sh

bootstrap:
	- cdk bootstrap

destroy:
	- cdk destroy