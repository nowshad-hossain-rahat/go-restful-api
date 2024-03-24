package config

import (
	"fmt"
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

type Smtp struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Secure bool
}

type Config struct {
	App     App
	MongoDB MongoDB
	Smtp    Smtp
}

func Load() *Config {

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("[!] CWD not found")
		os.Exit(1)
		return nil
	}

	viper.SetConfigFile(cwd + "/config.json")

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println("[!] Config file not found")
		os.Exit(1)
		return nil
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
		Smtp: Smtp{
			Host:   viper.GetString("Smtp.Host"),
			Port:   viper.GetString("Smtp.Port"),
			User:   viper.GetString("Smtp.User"),
			Pass:   viper.GetString("Smtp.Pass"),
			Secure: viper.GetBool("Smtp.Secure"),
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

	if config.Smtp.Host == "" {
		config.Smtp.Host = "smtp.mail.yahoo.com"
	}

	if config.Smtp.Port == "" {
		config.Smtp.Port = "587"
	}

	if config.Smtp.User == "" {
		config.Smtp.User = "nhrprime@yahoo.com"
	}

	if config.Smtp.Pass == "" {
		config.Smtp.Pass = ""
	}

	return &config

}
