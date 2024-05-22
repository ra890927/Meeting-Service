package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AdminRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the context
		userID, _ := c.Get("validate_user_id")
		userIDInt := uint(userID.(uint64))

		// check if the user is an admin
		// if not, return 403
		if userIDInt != 1 {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
