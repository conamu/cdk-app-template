arch := arm64

local-deploy:
	- rm -rd cdk.out
	- cd internal/app/ping && GOOS=linux GOARCH=$(arch) go build -o bootstrap . && zip bootstrap.zip bootstrap
	- cdklocal synth
	- cdklocal bootstrap
	- cdklocal deploy --all

localstack-up:
	- docker-compose -f localstack.compose up -d

localstack-down:
	- docker-compose -f localstack.compose down

build:
	- ./scripts/build.sh