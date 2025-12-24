package game

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection

func Setup(client *mongo.Client) {
	userCollection = client.Database("test").Collection("users")
}

func TrackGame(userId string, gameId int) error {
	// Add game to user's tracked games
	snowflakeId, _ := snowflake.ParseString(userId)
	_, err := userCollection.UpdateByID(
		context.Background(),
		snowflakeId,
		bson.D{{Key: "$addToSet", Value: bson.D{{Key: "games", Value: gameId}}}},
	)

	return err
}

func UntrackGame(userId string, gameId int) error {
	snowflakeId, _ := snowflake.ParseString(userId)
	_, err := userCollection.UpdateByID(
		context.Background(),
		snowflakeId,
		bson.D{{Key: "$pull", Value: bson.D{{Key: "games", Value: gameId}}}},
	)

	return err
}

func GetTrackedGames(userId string) []igdb.Game {
	currUser := userCollection.FindOne(context.Background(), map[string]string{"_id": userId})
	if currUser.Err() == mongo.ErrNoDocuments {
		return []igdb.Game{}
	}

	currUserDecoded := struct {
		Games []int `bson:"games"`
	}{}

	err := currUser.Decode(&currUserDecoded)
	if err != nil {
		return []igdb.Game{}
	}

	games := igdb.GetGamesByIds(currUserDecoded.Games)
	return games
}
