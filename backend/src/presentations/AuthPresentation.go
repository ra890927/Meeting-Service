package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"

	"github.com/gin-gonic/gin"
)

type AuthPresentation interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authPresentation struct {
}

func NewAuthPresentation(authServiceArgs ...services.AuthService) AuthPresentation {
	if len(authServiceArgs) == 1 {
		return AuthPresentation(&authPresentation{})
	} else if len(authServiceArgs) == 0 {
		return AuthPresentation(&authPresentation{})
	} else {
		panic("Too many arguments")
	}
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary Login a user
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body LoginParams true "User object"
// @Success 200 {object} models.User
// @Router /Auth/login [post]
func (ap authPresentation) Login(c *gin.Context) {
	// Get the user from the request
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Login the user
	loggedUser, token, err := services.NewAuthService().Login(&user)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// Return the logged user
	c.JSON(200, gin.H{"user": loggedUser, "token": token})
}

// @Summary Logout a user
// @Description Logout a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Param token header string true "Token"
// @Success 200
// @Router /Auth/logout [post]
func (ap authPresentation) Logout(c *gin.Context) {
	// Get the user from the request
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Get the token from the request
	token := c.GetHeader("token")

	// Logout the user
	err = services.NewAuthService().Logout(&user, &token)

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// Return the logged user
	c.JSON(200, gin.H{"message": "User logged out"})
}
