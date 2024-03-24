package controllers

import (
	"net/http"
	"restful-api/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type HomeController struct {
	cfg *config.Config
	db  *mongo.Database
}

func NewHomeController(cfg *config.Config, db *mongo.Database) *HomeController {
	return &HomeController{
		cfg: cfg,
		db:  db,
	}
}

type IndexData struct {
	Title    string
	Username string
}

func (c *HomeController) Index(ctx *gin.Context) {

	data := IndexData{
		Title:    "Go Restful API",
		Username: "Guest",
	}

	ctx.HTML(http.StatusOK, "index.html", data)

}
