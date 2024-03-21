package routes

import (
	"restful-api/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectRoutes(router *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {

	UserRoutes(router, cfg, db)

}
