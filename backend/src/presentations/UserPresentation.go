package presentations

import (
	"meeting-center/src/models"
	"meeting-center/src/services"

	"github.com/gin-gonic/gin"
)

type UserPresentation interface {
	RegisterUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type userPresentation struct {
	userService services.UserService
}

func NewUserPresentation(userServiceArgs ...services.UserService) UserPresentation {
	if len(userServiceArgs) == 1 {
		return UserPresentation(&userPresentation{userService: userServiceArgs[0]})
	} else if len(userServiceArgs) == 0 {
		return UserPresentation(&userPresentation{userService: services.NewUserService()})
	} else {
		panic("Too many arguments")
	}
}

type RegisterUserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterUpdateUserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		User struct {
			ID       uint   `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		} `json:"user"`
	} `json:"data"`
}

type UpdateUserBody struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type GetAllUsersResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Users []struct {
			ID       uint   `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		} `json:"users"`
	} `json:"data"`
}

// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body RegisterUserBody true "User object"
// @Success 200 {object} RegisterUpdateUserResponse
// @Router /user [post]
func (up userPresentation) RegisterUser(c *gin.Context) {
	var registerUserBody RegisterUserBody
	var registerUpdateUserResponse RegisterUpdateUserResponse

	// Get the user from the request
	err := c.BindJSON(&registerUserBody)
	if err != nil {
		registerUpdateUserResponse.Status = "error"
		registerUpdateUserResponse.Message = err.Error()
		c.JSON(400, registerUpdateUserResponse)
		return
	}

	user := models.User{
		Username: registerUserBody.Username,
		Password: registerUserBody.Password,
		Email:    registerUserBody.Email,
	}

	// Register the user
	createdUser, err := up.userService.CreateUser(&user)

	if err != nil {
		registerUpdateUserResponse.Status = "error"
		registerUpdateUserResponse.Message = err.Error()
		c.JSON(500, registerUpdateUserResponse)
		return
	}

	// Return the created user
	registerUpdateUserResponse.Status = "success"
	registerUpdateUserResponse.Message = "User created successfully"
	registerUpdateUserResponse.Data.User.ID = createdUser.ID
	registerUpdateUserResponse.Data.User.Username = createdUser.Username
	registerUpdateUserResponse.Data.User.Email = createdUser.Email
	registerUpdateUserResponse.Data.User.Role = createdUser.Role
	c.JSON(200, registerUpdateUserResponse)
}

// @Summary Get user by ID (要不要一次查詢多個用戶)

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} GetAllUsersResponse
// @Router /user/getAllUsers [get]
func (up userPresentation) GetAllUsers(c *gin.Context) {
	var getAllUsersResponse GetAllUsersResponse

	// Get all users
	users, err := up.userService.GetAllUsers()
	if err != nil {
		getAllUsersResponse.Status = "error"
		getAllUsersResponse.Message = err.Error()
		c.JSON(500, getAllUsersResponse)
		return
	}

	getAllUsersResponse.Status = "success"
	getAllUsersResponse.Message = "Users retrieved successfully"
	getAllUsersResponse.Data.Users = make([]struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}, len(users))

	for i, user := range users {
		getAllUsersResponse.Data.Users[i].ID = user.ID
		getAllUsersResponse.Data.Users[i].Username = user.Username
		getAllUsersResponse.Data.Users[i].Email = user.Email
		getAllUsersResponse.Data.Users[i].Role = user.Role
	}

	c.JSON(200, getAllUsersResponse)
}

// @Summary Update user
// @Description Update user
// @Tags admin
// @Accept json
// @Produce json
// @Param user body UpdateUserBody true "User object"
// @Success 200 {object} RegisterUpdateUserResponse
// @Router /admin/user [put]
func (up userPresentation) UpdateUser(c *gin.Context) {
	operator := c.MustGet("validate_user").(models.User)
	var updateUserBody UpdateUserBody
	var registerUpdateUserResponse RegisterUpdateUserResponse

	// Get the user from the request
	err := c.BindJSON(&updateUserBody)
	if err != nil {
		registerUpdateUserResponse.Status = "error"
		registerUpdateUserResponse.Message = err.Error()
		c.JSON(400, registerUpdateUserResponse)
		return
	}

	user := models.User{
		ID:       updateUserBody.ID,
		Username: updateUserBody.Username,
		Password: updateUserBody.Password,
		Email:    updateUserBody.Email,
		Role:     updateUserBody.Role,
	}

	// Update the user
	updatedUser, err := up.userService.UpdateUser(&operator, &user)
	if err != nil {
		registerUpdateUserResponse.Status = "error"
		registerUpdateUserResponse.Message = err.Error()
		c.JSON(500, registerUpdateUserResponse)
		return
	}

	// Return the updated user
	registerUpdateUserResponse.Status = "success"
	registerUpdateUserResponse.Message = "User updated successfully"
	registerUpdateUserResponse.Data.User.ID = updatedUser.ID
	registerUpdateUserResponse.Data.User.Username = updatedUser.Username
	registerUpdateUserResponse.Data.User.Email = updatedUser.Email
	registerUpdateUserResponse.Data.User.Role = updatedUser.Role
	c.JSON(200, registerUpdateUserResponse)
}
