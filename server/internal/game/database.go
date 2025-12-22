package game

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection
var gameCollection *mongo.Collection

func Setup(client *mongo.Client) {
	userCollection = client.Database("test").Collection("users")
	gameCollection = client.Database("test").Collection("games")
}

func TrackGame(userId string, gameId int) error {
	game := gameCollection.FindOne(context.Background(), map[string]int{"id": gameId})

	if game.Err() == mongo.ErrNoDocuments {
		// Fetch game from IGDB and insert into DB
		igdbGame := igdb.GetGame(gameId)
		_, err := gameCollection.InsertOne(context.Background(), igdbGame)
		if err != nil {
			return err
		}
	}

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

func GetTrackedGames(userId string) []Game {
	currUser := userCollection.FindOne(context.Background(), map[string]string{"_id": userId})
	if currUser.Err() == mongo.ErrNoDocuments {
		return []Game{}
	}

	currUserDecoded := struct {
		Games []int `bson:"games"`
	}{}

	err := currUser.Decode(&currUserDecoded)
	if err != nil {
		return []Game{}
	}

	// Use helper to return cached games and fetch any missing ones from IGDB
	return GetOrFetchGames(currUserDecoded.Games)
}

// GetOrFetchGames returns games for the provided game IDs. It first loads cached games
// from the database, then fetches any missing games from IGDB, inserts them into the
// collection, and finally returns all games ordered to match the input IDs slice.
func GetOrFetchGames(gameIDs []int) []Game {
	if len(gameIDs) == 0 {
		return []Game{}
	}

	// Find cached games whose `id` field is in gameIDs
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: gameIDs}}}}
	cursor, err := gameCollection.Find(context.Background(), filter)
	if err != nil {
		return []Game{}
	}

	var cached []Game
	if err := cursor.All(context.Background(), &cached); err != nil {
		return []Game{}
	}

	// Build map of cached games by id
	cachedMap := make(map[int]Game, len(cached))
	for _, g := range cached {
		cachedMap[g.Id] = g
	}

	// Identify missing IDs
	var missing []int
	for _, id := range gameIDs {
		if _, ok := cachedMap[id]; !ok {
			missing = append(missing, id)
		}
	}

	// Fetch missing games from IGDB and insert them
	if len(missing) > 0 {
		igdbGames := igdb.GetGamesByIds(missing)

		for _, ig := range igdbGames {
			if ig.Id == 0 {
				continue
			}
			g := Game{
				Id:               ig.Id,
				Name:             ig.Name,
				Cover:            Cover{Id: ig.Cover.Id, Url: ig.Cover.Url},
				AggregatedRating: ig.AggregatedRating,
			}
			_, _ = gameCollection.InsertOne(context.Background(), g)
			cachedMap[g.Id] = g
		}
	}

	// Build ordered result matching input order
	out := make([]Game, 0, len(gameIDs))
	for _, id := range gameIDs {
		if g, ok := cachedMap[id]; ok {
			out = append(out, g)
		}
	}

	return out
}
