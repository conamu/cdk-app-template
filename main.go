package main

import (
	"cdk-app-template/config"
	"cdk-app-template/infrastructure"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	config.Init(env)
	infrastructure.BuildStack()
}
