package main

import (
	"github.com/gin-gonic/gin"

	"restful-api/config"
	"restful-api/routes"
)

func main() {
	var config config.Config = config.Load()

	router := gin.New()

	routes.ConnectRoutes(router, &config)

	router.Run(":" + config.App.Port)
}
