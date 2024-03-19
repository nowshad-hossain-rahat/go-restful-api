package models

type User struct {
	Id         string `bson:"_id" json:"_id"`
	FirstName  string `bson:"firstName" json:"firstName"`
	LastName   string `bson:"lastName" json:"lastName"`
	Email      string `bson:"email" json:"email"`
	Password   string `bson:"password" json:"password"`
	Role       string `bson:"role" json:"role"`
	IsVerified bool   `bson:"isVerified" json:"isVerified"`
	CreatedAt  int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt  int64  `bson:"updatedAt" json:"updatedAt"`
}
