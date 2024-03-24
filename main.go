package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"restful-api/config"
	"restful-api/mids"
	"restful-api/routes"
	"restful-api/utils"
)

func main() {
	var cfg *config.Config = config.Load()

	router := gin.New()

	router.Use(mids.CORSMiddleware())

	router.Static("/static", "./public/assets/")

	router.LoadHTMLGlob("public/*.html")

	apiRouter := router.Group("/api")

	var db *mongo.Database = config.Connect(cfg)

	routes.ConnectRoutes(router, apiRouter, cfg, db)

	go sendEmail(cfg)

	router.Run(":" + cfg.App.Port)
}

func sendEmail(cfg *config.Config) {

	type TestData struct {
		Name  string
		Title string
		Body  string
		Year  int32
	}

	err := utils.SendTemplatedEmail(
		cfg,
		[]string{
			"nowshad.hossain.rahat@gmail.com",
		},
		"Test email from Go Restful API",
		"Test email body",
		"general.template.html",
		TestData{
			Name:  "Nowshad Hossain",
			Title: "Test email from Go Restful API",
			Body:  "Test email body",
			Year:  int32(time.Now().Year()),
		},
	)

	if err != nil {
		fmt.Println(err.Error())
	}
}
