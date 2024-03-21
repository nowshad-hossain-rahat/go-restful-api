package controllers

import (
	"net/http"
	"restful-api/config"
	"restful-api/models"
	"restful-api/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	cfg *config.Config
	db  *mongo.Database
}

type LoginData struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type RegisterData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AuthTokenClaims struct {
	Id    primitive.ObjectID `json:"_id"`
	Email string             `json:"email"`
	jwt.StandardClaims
}

// methods

func NewUserController(cfg *config.Config, db *mongo.Database) *UserController {
	return &UserController{
		cfg: cfg,
		db:  db,
	}
}

func (c *UserController) GetAll(ctx *gin.Context) {

	var users []models.User

	cursor, err := c.db.Collection("users").Find(ctx, bson.M{}, options.Find().SetSort(bson.M{
		"createdAt": -1,
	}))

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := cursor.All(ctx, &users); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	var publicUsers []models.PublicUser = make([]models.PublicUser, len(users))

	for index, user := range users {
		newUser := models.PublicUser{
			Id:         user.Id.Hex(),
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			IsVerified: user.IsVerified,
			Role:       user.Role,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		}

		publicUsers[index] = newUser
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  publicUsers,
	})

}

func (c *UserController) Login(ctx *gin.Context) {

	var loginData LoginData

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var user models.User

	err := c.db.Collection("users").FindOne(ctx, bson.M{
		"email": loginData.Email,
	}, nil).Decode(&user)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Invalid credentials",
		})
		return
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))

	if passMatchErr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Invalid credentials",
		})
		return
	}

	claims := AuthTokenClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3600 * 24 * 30).Unix(),
		},
	}

	token := utils.SignToken(claims, c.cfg.App.JWTSecret)

	if token == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
		"user": models.PublicUser{
			Id:         user.Id.Hex(),
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			IsVerified: user.IsVerified,
			Role:       user.Role,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		},
	})
}

func (c *UserController) Register(ctx *gin.Context) {

	var registerData RegisterData

	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var user models.User

	err := c.db.Collection("users").FindOne(ctx, bson.M{
		"email": registerData.Email,
	}, nil).Decode(&user)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(registerData.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Failed to hash the password",
		})
		return
	}

	user = models.User{
		Id:         primitive.NewObjectID(),
		FirstName:  registerData.FirstName,
		LastName:   registerData.LastName,
		Email:      registerData.Email,
		Password:   string(hashedPassword),
		IsVerified: false,
		Role:       "user",
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	result, err := c.db.Collection("users").InsertOne(ctx, user)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Failed to create user",
		})
		return
	}

	token := utils.SignToken(AuthTokenClaims{
		Id:    result.InsertedID.(primitive.ObjectID),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3600 * 24 * 30).Unix(),
		},
	}, c.cfg.App.JWTSecret)

	if token == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Failed to generate the auth token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user": models.PublicUser{
			Id:         result.InsertedID.(primitive.ObjectID).Hex(),
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			Role:       user.Role,
			IsVerified: user.IsVerified,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		},
		"token": token,
	})

}
