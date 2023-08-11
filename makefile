arch := arm64

local-deploy:
	- rm -rd cdk.out
	- make build
	- ./scripts/localstack-deploy.sh

localstack-up:
	- docker-compose -f localstack.compose up -d

localstack-down:
	- docker-compose -f localstack.compose down

build:
	- ./scripts/build.sh