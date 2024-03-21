package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"restful-api/config"
	"restful-api/controllers"
)

func UserRoutes(router *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {

	userController := controllers.NewUserController(cfg, db)

	router.POST("/user/login", userController.Login)
	router.POST("/user/register", userController.Register)
	router.GET("/user/all", userController.GetAll)
}
