package controllers

import (
	"net/http"
	"restful-api/config"
	"restful-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	config config.Config
}

func NewUserController(config config.Config) *UserController {
	return &UserController{
		config: config,
	}
}

func (c *UserController) GetAll(ctx *gin.Context) {

	var users []models.User = []models.User{}

	users = append(users, models.User{
		Id:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "qYn6B@example.com",
		Password:  "123456",
		IsVerified: true,
		Role:      "user",
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  users,
	})

}
