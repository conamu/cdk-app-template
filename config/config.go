package config

import (
	"github.com/spf13/viper"
	"os"
)

func Init(env string) {

	if env != "production" && env != "staging" && env != "local" {
		env = "staging"
	}

	path := env + ".config.yml"

	file, err := os.Open("config/" + path)
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("yaml")

	err = viper.ReadConfig(file)
	if err != nil {
		panic(err)
	}
}
