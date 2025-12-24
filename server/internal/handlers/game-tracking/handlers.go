package gameTracking

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/game"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers/middleware"
)

func SetupGameTrackingHandlers(router *gin.Engine) {
	trackingGroup := router.Group("/gameTracking")

	trackingGroup.Use(middleware.AuthRequired())
	{
		trackingGroup.POST("/track", trackGameHandler)
		trackingGroup.POST("/untrack", untrackGameHandler)
		trackingGroup.GET("/games", getTrackedGamesHandler)
	}
}

func trackGameHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	request := &game.TrackGamesRequest{}
	if !exists {
		c.AbortWithStatus(401)
		return
	}

	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatus(400)
		return
	}

	gameId, err := strconv.ParseInt(request.Game, 64, 64)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	game.TrackGame(userID.(string), int(gameId))
}

func untrackGameHandler(c *gin.Context) {
}

func getTrackedGamesHandler(c *gin.Context) {
}
