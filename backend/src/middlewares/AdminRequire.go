package middlewares

import (
	"github.com/gin-gonic/gin"
	"meeting-center/src/models"
)

func AdminRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the context
		user := c.MustGet("validate_user").(models.User)

		// check if the user is an admin
		// if not, return 403
		if user.Role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}
