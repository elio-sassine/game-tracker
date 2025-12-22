package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/igdb"
)

func setupGames(router *gin.Engine) {
	router.GET("/search", getGames)
	router.GET("/game", getGame)
}

func getGames(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.String(http.StatusBadRequest, "No name provided")
	}

	games := igdb.GetGamesByName(name)

	c.JSON(http.StatusOK, &games)
}

func getGame(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.ParseInt(idString, 10, 32)

	if err != nil {
		c.String(http.StatusBadRequest, "ID is not valid")
	}

	fmt.Println(id)

	game := igdb.GetGame(int(id))

	c.JSON(http.StatusOK, game)
}
