package user

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RefreshToken(c *gin.Context) {
	cookie, err := c.Request.Cookie("JWT_REFRESH")
	if err != nil {
		log.Println("No refresh token cookie found: ", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := ValidateCookie(cookie, true)
	if err != nil {
		log.Println("Invalid refresh token: ", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id, err := snowflake.ParseString(claims.ID)

	tokenHash, err := GetRefreshToken(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No refresh token found for user ID: ", id)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Panic(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(*tokenHash), []byte(cookie.Value)) != nil {
		log.Println("Refresh token hash does not match")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// If the claims are valid, generate a new access token
	accessToken := generateToken(claims.Email, id, 60*time.Minute, false)
	http.SetCookie(c.Writer, generateCookie(c, accessToken, false))
}

func ValidateCookie(cookie *http.Cookie, isRefreshToken bool) (*Claims, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if token == nil || err != nil {
		println("Err parsing token: " + err.Error())
		return nil, err
	}

	id, err := snowflake.ParseString(claims.ID)
	if err != nil {
		println("Err parsing id: " + err.Error() + " for token: " + claims.ID)
		return nil, err
	}

	if isRefreshToken {
		GetRefreshToken(id)
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

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
