package middleware

import (
	"net/http"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/phoenix-of-dawn/game-tracker/server/internal/user"
)

const ctxUserIDKey = "userID"

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("JWT_TOKEN")
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		claims, err := user.ValidateCookie(cookie, false)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		if claims.Subject == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
			return
		}

		if id, err := strconv.ParseInt(claims.Subject, 10, 64); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
		} else {
			c.Set(ctxUserIDKey, snowflake.ID(id))
			c.Next()
		}
	}
}

func AuthOptional() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("JWT_TOKEN")
		if err != nil {
			c.Next()
			return
		}

		claims, err := user.ValidateCookie(cookie, false)
		if err != nil {
			c.Next()
			return
		}

		if claims.Subject == "" {
			c.Next()
			return
		}

		if id, err := strconv.ParseInt(claims.Subject, 10, 64); err != nil {
			c.Next()
		} else {
			c.Set(ctxUserIDKey, snowflake.ID(id))
			c.Next()
		}
	}
}
