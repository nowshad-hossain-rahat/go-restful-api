package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"restful-api/config"
	"restful-api/mids"
	"restful-api/routes"
)

func main() {
	var cfg *config.Config = config.Load()

	fmt.Println(*cfg)

	router := gin.New()

	router.Use(mids.CORSMiddleware())

	apiRouter := router.Group("/api")

	var db *mongo.Database = config.Connect(cfg)

	routes.ConnectRoutes(apiRouter, cfg, db)

	router.Run(":" + cfg.App.Port)
}
