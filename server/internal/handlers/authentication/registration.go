package authentication

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
)

func SetupRegistration(router *gin.Engine) {
	router.POST("/register", registerUser)
}

func registerUser(c *gin.Context) {
	newUserRequest := &user.UserRegisterRequest{}

	if err := c.BindJSON(newUserRequest); err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := user.RegisterUser(newUserRequest)
	if err != nil {
		log.Panic(err)
	}

	c.IndentedJSON(http.StatusAccepted, nil)
}
