package config

import (
	"os"

	"github.com/spf13/viper"
)

type App struct {
	Name      string
	Port      string
	Root      string
	JWTSecret string
}

type MongoDB struct {
	URI  string
	Name string
}

type Config struct {
	App     App
	MongoDB MongoDB
}

func Load() *Config {

	cwd, err := os.Getwd()

	if err != nil {
		os.Exit(1)
	}

	viper.SetConfigFile(cwd + "/config.json")

	err = viper.ReadInConfig()

	if err != nil {
		os.Exit(1)
	}

	var config Config = Config{
		App: App{
			Name:      viper.GetString("App.Name"),
			Port:      viper.GetString("App.Port"),
			JWTSecret: viper.GetString("App.JWTSecret"),
			Root:      cwd,
		},
		MongoDB: MongoDB{
			URI:  viper.GetString("MongoDB.URI"),
			Name: viper.GetString("MongoDB.Name"),
		},
	}

	if config.App.Name == "" {
		config.App.Name = "Restful API"
	}

	if config.App.Port == "" {
		config.App.Port = "3030"
	}

	if config.App.JWTSecret == "" {
		panic("[!] JWT Secret is empty")
	}

	if config.MongoDB.URI == "" {
		config.MongoDB.URI = "mongodb://localhost:27017"
	}

	if config.MongoDB.Name == "" {
		config.MongoDB.Name = "go_restful_api"
	}

	return &config

}
