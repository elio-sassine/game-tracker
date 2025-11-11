package user

import (
	"context"
	"log"
	"time"

	"github.com/bwmarrin/snowflake"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
var refreshTokenColl *mongo.Collection

type RefreshTokenDoc struct {
	ID        snowflake.ID `bson:"_id"`
	Token     string       `bson:"token"`
	ExpireAt  time.Time    `bson:"expireAt"` // must be time.Time for TTL index
	CreatedAt time.Time    `bson:"createdAt"`
}

func Setup(client *mongo.Client) {
	userCollection = client.Database("test").Collection("users")
	refreshTokenColl = client.Database("test").Collection("refresh_tokens")
}

func InsertRefreshToken(id snowflake.ID, token *string) (*string, error) {
	hashedTokenString, err := bcrypt.GenerateFromPassword([]byte(*token), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}

	_, err = refreshTokenColl.InsertOne(
		context.Background(),
		&RefreshTokenDoc{
			ID:        id,
			Token:     string(hashedTokenString),
			ExpireAt:  time.Now().Add(3 * 24 * time.Hour), // set expiration time
			CreatedAt: time.Now(),
		},
	)

	return token, err
}

func GetRefreshToken(id snowflake.ID) (*string, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var result RefreshTokenDoc
	refreshToken, err := getRefreshTokenWithFilter(filter, &result)

	return refreshToken, err
}

func InsertUser(user *User) (*User, error) {
	_, err := userCollection.InsertOne(context.Background(), user)
	return user, err
}

func GetUserByEmail(email string) (*User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var result User

	user, err := getUserWithFilter(filter, &result)

	return user, err
}

func UserExists(email string) bool {
	user, _ := GetUserByEmail(email)

	return user != nil
}

func GetUserByID(id int64) (*User, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var result User
	user, err := getUserWithFilter(filter, &result)

	return user, err
}

func getUserWithFilter(filter bson.D, result *User) (*User, error) {
	err := userCollection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			println("Error no documents found")
			return nil, err
		}

		log.Panic(err)
	}

	println("documents found")
	println(result.Id)
	return result, nil
}

func getRefreshTokenWithFilter(filter bson.D, result *RefreshTokenDoc) (*string, error) {
	err := refreshTokenColl.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			println("Error no documents found")
			return nil, err
		}

		log.Panic(err)
	}
	println("documents found")
	println(result.Token)

	return &result.Token, nil
}
