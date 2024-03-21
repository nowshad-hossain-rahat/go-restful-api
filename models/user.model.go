package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	FirstName  string `bson:"firstName" json:"firstName"`
	LastName   string `bson:"lastName" json:"lastName"`
	Email      string `bson:"email" json:"email"`
	Password   string `bson:"password" json:"password"`
	Role       string `bson:"role" json:"role"`
	IsVerified bool   `bson:"isVerified" json:"isVerified"`
	CreatedAt  int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt  int64  `bson:"updatedAt" json:"updatedAt"`
}

type PublicUser struct {
	Id         string `json:"_id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	IsVerified bool   `json:"isVerified"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}
