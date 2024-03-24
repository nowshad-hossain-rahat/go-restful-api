package routes

import (
	"restful-api/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectRoutes(router *gin.Engine, apiRouter *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {

	HomeRoutes(router, cfg, db)
	UserRoutes(apiRouter, cfg, db)

}
