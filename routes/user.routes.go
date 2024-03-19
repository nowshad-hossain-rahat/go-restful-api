package routes

import (
	"github.com/gin-gonic/gin"

	"restful-api/config"
	"restful-api/controllers"
)

func UserRoutes(router *gin.Engine, config *config.Config) {

	userController := controllers.NewUserController(*config)

	router.GET("/api/user/all", userController.GetAll)
}
