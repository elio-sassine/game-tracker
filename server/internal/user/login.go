package user

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	origin := c.Request.Header.Get("Origin")
	secure := c.Request.TLS != nil || strings.HasPrefix(strings.ToLower(origin), "https://")

	verifMaxAge := int((60 * time.Minute).Seconds())
	refreshMaxAge := int((60 * time.Hour).Seconds())

	ss := http.SameSiteLaxMode
	if secure {
		ss = http.SameSiteNoneMode
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "JWT_TOKEN",
		Value:    verifTkn,
		Path:     "/",
		MaxAge:   verifMaxAge,
		HttpOnly: true,
		Secure:   secure,
		SameSite: ss,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "JWT_REFRESH",
		Value:    refreshTkn,
		Path:     "/",
		MaxAge:   refreshMaxAge,
		HttpOnly: true,
		Secure:   secure,
		SameSite: ss,
	})

}

func generateToken(email string, id snowflake.ID, expiryTime time.Duration, isRefreshToken bool) string {
	jwtType := "validation"
	if isRefreshToken {
		jwtType = "refresh"
	}

	claims := Claims{
		Email: email,
		Type:  jwtType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryTime)),
			Subject:   id.String(),
			Issuer:    "server",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signingKey))

	if err != nil {
		log.Panic(err)
		return ""
	}

	return token
}

func checkPassword(password string, encryptedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	return err == nil
}
