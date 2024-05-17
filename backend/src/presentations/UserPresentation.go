package presentations

import (
	// "fmt"
	"meeting-center/src/models"
	"meeting-center/src/services"

	"github.com/gin-gonic/gin"
)

type UserPresentation interface {
	RegisterUser(c *gin.Context)
	UpdateUser(c *gin.Context)
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
	// filter out the parameters that are not needed
	filtered_user := models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}
	// Register the user
	createdUser, err := up.UserService.CreateUser(&filtered_user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// Return the created user
	c.JSON(200, createdUser)
}

// @Summary Get user by ID (要不要一次查詢多個用戶)

// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Router /User [put]
func (up *userPresentation) UpdateUser(c *gin.Context) {
	// Get the user from the request
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	// filter out the parameters that are not needed
	filtered_user := models.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}

	// Update the user
	updatedUser, err := up.UserService.UpdateUser(&filtered_user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// Return the updated user
	c.JSON(200, updatedUser)
}
