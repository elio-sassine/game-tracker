package authentication

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	projectErrors "github.com/phoenix-of-dawn/game-tracker/server/internal/common/project-errors"
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
	if err == projectErrors.ErrUserNotUnique {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "user with this email already exists"})
		return
	}

	if err != nil {
		log.Panic(err)
	}

	c.IndentedJSON(http.StatusAccepted, nil)
}
