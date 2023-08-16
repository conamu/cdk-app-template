package config

import "github.com/spf13/viper"

func Init() {
	viper.SetDefault("aws-account", "038796470268")
	viper.SetDefault("aws-region", "eu-west-1")
}
