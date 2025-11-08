package user

import (
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var signingKey = os.Getenv("JWT_KEY")

func CheckUser(email string, password string) bool {
	user, err := GetUserByEmail(email)
	if err != nil {
		return false
	}

	return checkPassword(password, user.Password)
}

func SetTokens(c *gin.Context, email string, id snowflake.ID) {
	verifTkn := generateToken(email, id, 60*time.Minute, false)
	refreshTkn := generateToken(email, id, 3*24*time.Hour, true)

	http.SetCookie(c.Writer, generateCookie(c, verifTkn, false))
	http.SetCookie(c.Writer, generateCookie(c, refreshTkn, true))
}

func checkPassword(password string, encryptedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	return err == nil
}
