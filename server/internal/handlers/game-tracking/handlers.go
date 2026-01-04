package gameTracking

import (
	"log"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/game"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers/middleware"
)

func SetupGameTrackingHandlers(router *gin.Engine) {
	trackingGroup := router.Group("/gameTracking")

	trackingGroup.Use(middleware.AuthRequired())

	trackingGroup.POST("/track", trackGameHandler)
	trackingGroup.POST("/untrack", untrackGameHandler)

	router.GET("/getTrackedGames", getTrackedGamesHandler)
}

func trackGameHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	request := &game.TrackGamesRequest{}
	if !exists {
		c.AbortWithStatus(401)
		return
	}

	switch v := userID.(type) {
	case snowflake.ID:
		userID = v.String()
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "User ID invalid, sign in again"})
		log.Panic("User ID invalid")
		return
	}

	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		log.Panic("Failed to bind JSON: ", err)
		return
	}

	_, err := strconv.ParseInt(request.Game, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Failed to parse game ID: " + err.Error()})
		log.Panic("Failed to parse game ID: ", err)
		return
	}

	game.TrackGame(userID.(string), request.Game)
}

func untrackGameHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	request := &game.TrackGamesRequest{}

	switch v := userID.(type) {
	case snowflake.ID:
		userID = v.String()
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "User ID invalid, sign in again"})
		log.Panic("User ID invalid")
		return
	}

	if !exists {
		c.AbortWithStatus(401)
		return
	}

	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Failed to bind JSON: " + err.Error()})
		log.Panic("Failed to bind JSON: ", err)
		return
	}

	_, err := strconv.ParseInt(request.Game, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Failed to parse game ID: " + err.Error()})
		log.Panic("Failed to parse game ID: ", err)
		return
	}

	game.UntrackGame(userID.(string), request.Game)
}

func getTrackedGamesHandler(c *gin.Context) {
	userID := c.Query("userID")

	if userID == "" {
		c.AbortWithStatus(401)
		return
	}

	trackedGames := game.GetTrackedGames(userID)
	c.IndentedJSON(200, trackedGames)
}
