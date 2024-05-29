package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthPresentation interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	WhoAmI(c *gin.Context)
}

type authPresentation struct {
	as services.AuthService
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Status string `json:"status"`
	Data   struct {
		User struct {
			ID        uint      `json:"id"`
			Username  string    `json:"username"`
			Email     string    `json:"email"`
			Role      string    `json:"role"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"user"`
		Token   string `json:"token"`
		Message string `json:"message"`
	} `json:"data"`
}

type LogoutResponse struct {
	Status string `json:"status"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data"`
}

type WhoAmIResponse struct {
	Status string `json:"status"`
	Data   struct {
		User struct {
			ID       uint   `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		} `json:"user"`
	} `json:"data"`
	Message string `json:"message"`
}

func NewAuthPresentation(authServiceArgs ...services.AuthService) AuthPresentation {
	if len(authServiceArgs) == 1 {
		return AuthPresentation(&authPresentation{as: authServiceArgs[0]})
	} else if len(authServiceArgs) == 0 {
		return AuthPresentation(&authPresentation{as: services.NewAuthService()})
	} else {
		panic("Too many arguments")
	}
}

// @Summary Login a user
// @Description Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body LoginParams true "User object"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func (ap authPresentation) Login(c *gin.Context) {
	// Get the user from the request
	var loginParams LoginParams
	var loginResponse LoginResponse
	var user models.User
	err := c.BindJSON(&loginParams)
	if err != nil {
		loginResponse.Status = "fail"
		loginResponse.Data.Message = "Invalid request"
		c.JSON(400, loginResponse)
		return
	}

	// construct user object
	user.Email = loginParams.Email
	user.Password = loginParams.Password

	// Login the user
	loggedUser, token, err := ap.as.Login(&user)
	if err != nil {
		loginResponse.Status = "fail"
		loginResponse.Data.Message = err.Error()
		c.JSON(500, loginResponse)
		return
	}

	c.SetCookie("token", *token, 3600*24, "/", "", false, true)

	// Return the logged user
	loginResponse.Status = "success"
	loginResponse.Data.User.ID = loggedUser.ID
	loginResponse.Data.User.Username = loggedUser.Username
	loginResponse.Data.User.Email = loggedUser.Email
	loginResponse.Data.User.Role = loggedUser.Role
	loginResponse.Data.User.CreatedAt = loggedUser.CreatedAt
	loginResponse.Data.User.UpdatedAt = loggedUser.UpdatedAt
	loginResponse.Data.Token = *token
	loginResponse.Data.Message = "User logged in"
	c.JSON(200, loginResponse)
}

// @Summary Logout a user
// @Description Logout a user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} LogoutResponse
// @Router /auth/logout [post]
func (ap authPresentation) Logout(c *gin.Context) {
	// Get the user from the request
	var logoutResponse LogoutResponse

	// Get the token from cookie
	tokenValue := ""
	token, err := c.Request.Cookie("token")
	if err != nil {
		// if cant get from cookie, get from header for CORS problem
		tokenValue = c.GetHeader("Authorization")
		if tokenValue == "" {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
	} else {
		tokenValue = token.Value
	}

	// Logout the user
	err = ap.as.Logout(&tokenValue)
	if err != nil {
		logoutResponse.Status = "fail"
		logoutResponse.Data.Message = err.Error()
		c.JSON(500, logoutResponse)
		return
	}

	// Logout the user
	c.SetCookie("token", "", -1, "/", "", false, true)

	// Return the logged out message
	logoutResponse.Status = "success"
	logoutResponse.Data.Message = "User logged out"
	c.JSON(200, logoutResponse)
}

// @Summary Get the user who is logged in
// @Description Get the user who is logged in
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} WhoAmIResponse
// @Router /auth/whoami [get]
func (ap authPresentation) WhoAmI(c *gin.Context) {
	var whoAmIResponse WhoAmIResponse

	// get user if from middleware's context
	user, success := c.Get("validate_user")
	if !success {
		whoAmIResponse.Status = "fail"
		whoAmIResponse.Message = "User not found"
		c.JSON(400, whoAmIResponse)
		return
	}
	modelUser := user.(models.User)

	whoAmIResponse.Status = "success"
	whoAmIResponse.Message = "User found"
	whoAmIResponse.Data.User.ID = modelUser.ID
	whoAmIResponse.Data.User.Username = modelUser.Username
	whoAmIResponse.Data.User.Email = modelUser.Email
	whoAmIResponse.Data.User.Role = modelUser.Role

	c.JSON(200, whoAmIResponse)
}
