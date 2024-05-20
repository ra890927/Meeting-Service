package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AdminRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the context
		userID, _ := c.Get("validate_user_id")

		// check if the user is an admin
		// if not, return 403
		if userID != 1 {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
