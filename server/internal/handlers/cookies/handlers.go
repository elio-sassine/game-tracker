package cookies

import (
	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
)

func SetupCookies(router *gin.Engine) {
	router.POST("/refresh-token", refreshTokenHandler)
}

func refreshTokenHandler(c *gin.Context) {
	user.RefreshToken(c)
}
