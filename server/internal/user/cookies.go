package user

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func generateCookie(c *gin.Context, token string, isRefreshToken bool) *http.Cookie {
	maxAge := int((60 * time.Minute).Seconds())
	name := "JWT_TOKEN"
	if isRefreshToken {
		maxAge = int((3 * 24 * time.Hour).Seconds())
		name = "JWT_REFRESH"
	}

	origin := c.Request.Header.Get("Origin")
	secure := c.Request.TLS != nil || strings.HasPrefix(strings.ToLower(origin), "https://")

	ss := http.SameSiteLaxMode
	if secure {
		ss = http.SameSiteNoneMode
	}

	return &http.Cookie{
		Name:     name,
		Value:    token,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   secure,
		SameSite: ss,
	}
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
