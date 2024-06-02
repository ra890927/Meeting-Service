package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"meeting-center/src/clients"
	"meeting-center/src/models"
)

func AuthRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from cookie first, if not found, get it from header
		token, err := c.Request.Cookie("token")
		if err != nil {
			// not found in cookie, get it from header's Authorization
			tokenValue := c.GetHeader("Authorization")
			if tokenValue == "" {
				c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
				return
			} else {
				token = &http.Cookie{Name: "token", Value: tokenValue}
			}
		}

		redisClient := clients.GetRedisInstance()

		// Check if the token is valid
		marshaledValue, err := redisClient.Get(redisClient.Context(), token.Value).Result()
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		// Unmarshal the value
		var user models.User
		err = json.Unmarshal([]byte(marshaledValue), &user)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Set the user ID in the context
		c.Set("validate_user", user)

		c.Next()
	}
}
