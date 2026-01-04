package game

import (
	"context"
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection

func Setup(client *mongo.Client) {
	userCollection = client.Database("test").Collection("users")
}

func TrackGame(userId string, gameId string) error {
	// Add game to user's tracked games
	snowflakeId, _ := snowflake.ParseString(userId)
	_, err := userCollection.UpdateByID(
		context.Background(),
		snowflakeId,
		bson.D{{Key: "$addToSet", Value: bson.D{{Key: "games", Value: gameId}}}},
	)

	return err
}

func UntrackGame(userId string, gameId string) error {
	snowflakeId, _ := snowflake.ParseString(userId)
	_, err := userCollection.UpdateByID(
		context.Background(),
		snowflakeId,
		bson.D{{Key: "$pull", Value: bson.D{{Key: "games", Value: gameId}}}},
	)

	return err
}

func GetTrackedGames(userId string) []igdb.Game {
	snowflakeId, err := snowflake.ParseString(userId)
	if err != nil {
		log.Print("Error parsing user ID: ", err)
		return []igdb.Game{}
	}

	currUser := userCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: snowflakeId}})
	if currUser.Err() == mongo.ErrNoDocuments {
		log.Print("No user found with ID: ", userId)
		return []igdb.Game{}
	}

	currUserDecoded := struct {
		Games []string `bson:"games"`
	}{}

	if err := currUser.Decode(&currUserDecoded); err != nil {
		log.Print("Error decoding user data: ", err)
		return []igdb.Game{}
	}

	if len(currUserDecoded.Games) == 0 {
		return []igdb.Game{}
	}

	log.Print("Fetching games for IDs: ", currUserDecoded.Games)
	games := igdb.GetGamesByIds(currUserDecoded.Games)
	log.Print("Fetched games: ", games)
	return games
}
