package config

import (
	"fmt"
	"os"
)

func Connect(config Config) {

	if (config.MongoDB.URI == "") || (config.MongoDB.Name == "") {
		fmt.Println("p!] MongoDB URI or Name is empty")
		os.Exit(1)
	}

	

}
