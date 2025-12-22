package user

import (
	"github.com/bwmarrin/snowflake"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       snowflake.ID `json:"_id" bson:"_id"`
	Username string       `json:"username" bson:"username"`
	Password string       `json:"password" bson:"password"`
	Email    string       `json:"email" bson:"email"`
	Games    []string     `json:"games" bson:"games"`
}

type UserRequest struct {
	Id       snowflake.ID `json:"_id" bson:"_id"`
	Username string       `json:"username" bson:"username"`
	Games    []string     `json:"games" bson:"games"`
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	Type  string `json:"typ"`
	jwt.RegisteredClaims
}
