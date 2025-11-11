package authentication

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupLogin(router *gin.Engine) {
	router.POST("/login", loginUser)
}

func loginUser(c *gin.Context) {
	userLoginRequest := &user.UserLoginRequest{}

	if err := c.BindJSON(userLoginRequest); err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	email := userLoginRequest.Email
	password := userLoginRequest.Password

	if !user.CheckUser(email, password) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	userFetch, err := user.GetUserByEmail(email)
	if err == mongo.ErrNoDocuments {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		log.Panic(err)
	}

	println(userFetch.Id)

	id := userFetch.Id

	user.SetTokens(c, email, id)

	println("User logged in: ", id)
	c.IndentedJSON(http.StatusOK, id.String())
}
