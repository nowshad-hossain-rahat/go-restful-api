package routes

import (
	"restful-api/config"
	"github.com/gin-gonic/gin"
)

func ConnectRoutes(router *gin.Engine, config *config.Config) {

	UserRoutes(router, config)

}
