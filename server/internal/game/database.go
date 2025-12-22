package game

import "go.mongodb.org/mongo-driver/v2/mongo"

var userCollection *mongo.Collection
var gameCollection *mongo.Collection

func Setup(client *mongo.Client) {
	userCollection = client.Database("test").Collection("users")
	gameCollection = client.Database("test").Collection("games")
}

func TrackGame(userId string, gameId int) error {
	return nil
}

func UntrackGame(userId string, gameId int) error {
	return nil
}
