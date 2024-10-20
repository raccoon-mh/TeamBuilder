package constvariables

import (
	"log"

	"github.com/spf13/viper"
)

var (
	HOST                string
	PORT                string
	DATABASE_PATH       string
	DiscordClientID     string
	DiscordClientSecret string
	DiscordRedirect     string
	DiscordCallback     string
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
	HOST = getEnvWithDefault("domain")

	PORT = getEnvWithDefault("port")

	DATABASE_PATH = getEnvWithDefault("databasepath")

	DiscordClientID = getEnvWithDefault("discordclientid")
	DiscordClientSecret = getEnvWithDefault("discordclientsecret")
	DiscordRedirect = getEnvWithDefault("discordredirect")
	DiscordCallback = getEnvWithDefault("discordcallback")
}

func getEnvWithDefault(path string) string {
	value := viper.GetString("environment." + path + ".value")
	if value == "" {
		value = viper.GetString("environment." + path + ".default")
	}
	return value
}
