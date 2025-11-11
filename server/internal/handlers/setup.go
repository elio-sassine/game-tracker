package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers/authentication"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/handlers/cookies"
)

func Setup(router *gin.Engine) {
	setupGames(router)
	setupUser(router)
	authentication.SetupRegistration(router)
	authentication.SetupLogin(router)
	cookies.SetupCookies(router)
}
