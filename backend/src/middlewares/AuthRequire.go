package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
		userID, err := redisClient.Get(redisClient.Context(), token.Value).Result()
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		// convert the user ID to uint
		userIDUint, err := strconv.ParseUint(userID, 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// Set the user ID in the context
		c.Set("validate_user_id", userIDUint)

		// Close the Redis client
		redisClient.Close()

		c.Next()
	}
}
