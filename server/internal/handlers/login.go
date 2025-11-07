package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
)

func setupLogin(router *gin.Engine) {
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

	userFetch, _ := user.GetUserByEmail(email)
	id := userFetch.Id

	user.SetTokens(c, email, id)

	c.IndentedJSON(http.StatusOK, id)
}
