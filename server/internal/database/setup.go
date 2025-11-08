package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client
var userColl *mongo.Collection
var refreshTokenColl *mongo.Collection

func Setup() {
	url := os.Getenv("DATABASE_URL")
	username := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	log.Print(url, username, pass)
	Client, _ = mongo.Connect(options.Client().ApplyURI("mongodb://" + username + ":" + pass + "@" + url))
	ctx := context.Background()

	print(Client)
	err := Client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Panic(err)
	}

	userColl = Client.Database("test").Collection("users")

	refreshTokenColl = Client.Database("test").Collection("refresh_tokens")

	user.Setup(Client)

	// Will throw an error if the definitions of the index models change
	createIndexes()
}

func createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := userColl.Indexes().CreateOne(ctx, indexModel)

	if err != nil {
		log.Panic(err)
	}

	ttlIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "expireAt", Value: 1}},
		Options: options.Index().SetName("expire_at_ttl").SetExpireAfterSeconds(0),
	}

	_, err = refreshTokenColl.Indexes().CreateOne(ctx, ttlIndex)

	if err != nil {
		log.Panic(err)
	}

}
