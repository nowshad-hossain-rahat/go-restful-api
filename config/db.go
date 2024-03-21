package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg *Config) *mongo.Database {

	if (cfg.MongoDB.URI == "") || (cfg.MongoDB.Name == "") {
		fmt.Println("[!] MongoDB URI or Name is empty")
		os.Exit(1)
	}

	// Connect to MongoDB

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))

	if err != nil {
		panic(err)
	}

	db := client.Database(cfg.MongoDB.Name)

	if db != nil {
		fmt.Println("[+] Connected to MongoDB")
	}

	return db

}
