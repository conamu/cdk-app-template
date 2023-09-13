deploy: build
	- ./scripts/deploy.sh

build:
	- arch=arm64 ./scripts/build.sh

bootstrap:
	- cdk bootstrap

destroy:
	- cdk destroy