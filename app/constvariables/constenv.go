package constvariables

import (
	"log"

	"github.com/spf13/viper"
)

var (
	DATABASE_PATH string
	PORT          string
)

func init() {
	viper.SetConfigFile("meta/env.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	setEnvs()
}

func setEnvs() {
	DATABASE_PATH = getEnvWithDefault("databasepath")
	PORT = getEnvWithDefault("port")
}

func getEnvWithDefault(path string) string {
	value := viper.GetString("environment." + path + ".value")
	if value == "" {
		value = viper.GetString("environment." + path + ".default")
	}
	return value
}
