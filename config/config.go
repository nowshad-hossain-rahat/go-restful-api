package config

import (
	"os"

	"github.com/spf13/viper"
)

type App struct {
	Name string
	Port string
	Root string
}

type MongoDB struct {
	URI string
	Name string
}

type Config struct {
	App App
	MongoDB MongoDB
}

func Load() Config {

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
			Name: viper.GetString("App.Name"),
			Port: viper.GetString("App.Port"),
			Root: cwd,
		},
	}

	if config.App.Name == "" {
		config.App.Name = "Restful API"
	}

	if config.App.Port == "" {
		config.App.Port = "3030"
	}

	return config

}
