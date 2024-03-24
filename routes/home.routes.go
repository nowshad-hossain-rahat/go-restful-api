package routes

import (
	"restful-api/config"
	"restful-api/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func HomeRoutes(router *gin.Engine, cfg *config.Config, db *mongo.Database) {

	homeController := controllers.NewHomeController(cfg, db)

	router.GET("/", homeController.Index)

}
