package main

import (
	"cdk-app-template/config"
	"cdk-app-template/infrastructure"
)

func main() {
	config.Init()
	infrastructure.BuildStack()
}
