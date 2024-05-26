package middlewares

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"meeting-center/src/models"
)

func AuthRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from cookie
		token, err := c.Request.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Create a new Redis client
		redisClient := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

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

		// Close the Redis client
		redisClient.Close()

		c.Next()
	}
}