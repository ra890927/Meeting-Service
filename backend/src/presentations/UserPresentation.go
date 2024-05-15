package presentations

import (
	// "fmt"
	"meeting-center/src/models"
	"meeting-center/src/services"

	"github.com/gin-gonic/gin"
)

type UserPresentation interface {
	RegisterUser(c *gin.Context)
}

type userPresentation struct {
	UserService services.UserService
}

func NewUserPresentation(opt ...services.UserService) UserPresentation {
	if len(opt) == 1 {
		return &userPresentation{
			UserService: opt[0],
		}
	} else {
		return &userPresentation{
			UserService: services.NewUserService(),
		}
	}
}

// RegisterUser registers a new user
// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Router /User [post]
func (up *userPresentation) RegisterUser(c *gin.Context) {
	// Get the user from the request
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Register the user
	createdUser, err := up.UserService.CreateUser(&user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// Return the created user
	c.JSON(200, createdUser)
}
