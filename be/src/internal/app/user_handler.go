package app

import (
	"net/http"
	"strconv"
	"workout-tracker/m/v0.0.0/src/internal/domain"
	"workout-tracker/m/v0.0.0/src/internal/exceptions"
	"workout-tracker/m/v0.0.0/src/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(us *UserService) *UserHandler {
	return &UserHandler{
		UserService: us,
	}
}

func (uh *UserHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := uh.UserService.UserRepository.GetUsers()
	a := c.Request.Header.Get("Authorization")
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to get users", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
	}
	usersInterface := make([]interface{}, len(users)+1)
	for i, user := range users {
		usersInterface[i] = user
	}
	usersInterface[len(users)] = map[string]string{"total": a}
	h := utils.NewResponse(http.StatusOK, "Users retrieved successfully", usersInterface)
	c.JSON(http.StatusOK, h)
}
func (uh *UserHandler) LoginHandler(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		e := exceptions.NewInvalidRequestError("Invalid request", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	username := requestBody.Username
	password := requestBody.Password

	accessToken, err := uh.UserService.Login(username, password)
	if err != nil {
		if err.Error() == "record not found" {
			e := exceptions.NewInvalidRequestError("Invalid username or password", err)
			errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		e := exceptions.NewAuthenticationError(err.Error())
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}
	resultInterface := make([]interface{},1)
	resultInterface[0] = map[string]string{"access_token": accessToken}
	h := utils.NewResponse(http.StatusOK, "Login successful", resultInterface) 
	c.JSON(http.StatusOK, h)
}

func (uh *UserHandler) RegisterHandler(c *gin.Context){
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		e := exceptions.NewInvalidRequestError("Invalid request", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if requestBody.Username == "" || requestBody.Password == "" || requestBody.Fullname == "" || requestBody.Email == "" {
		e := exceptions.NewInvalidRequestError("Invalid request", nil)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	user := domain.User{
		Username: requestBody.Username,
		Password: requestBody.Password,
		Fullname: requestBody.Fullname,
		Email:    requestBody.Email,
	}

	accessToken, err := uh.UserService.Register(&user)
	if err != nil {
		if err.Error() == "username already exists" {
			e := exceptions.NewInvalidRequestError("Username already exists", err)
			errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		e := exceptions.NewAuthenticationError(err.Error())
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusUnauthorized, errResponse)
		return
	}
	resultInterface := make([]interface{},1)
	resultInterface[0] = map[string]string{"access_token": accessToken}
	h := utils.NewResponse(http.StatusOK, "Registration successful", resultInterface) 
	c.JSON(http.StatusCreated, h)
}

func (uh *UserHandler) GetUserByIDHandler(c *gin.Context){
	id, err := utils.ExtractTokenID(c)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Invalid ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(e.ClientError.Code, errResponse)
		return
	}
	profile, err := uh.UserService.UserRepository.GetUserByID(int(id))
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to get user", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(e.ClientError.Code, errResponse)
		return
	}
	profileInterface := []interface{}{profile}
	h := utils.NewResponse(http.StatusOK, "User retrieved successfully", profileInterface)
	c.JSON(http.StatusOK, h)
}

func (uh *UserHandler) UpdateUserHandler(c *gin.Context){
	id := c.Params.ByName("id")
	profileID, err := strconv.Atoi(id)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Invalid ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(e.ClientError.Code, errResponse)
		return
	}

	var requestBody struct {
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		e := exceptions.NewInvalidRequestError("Invalid request", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	username := requestBody.Username
	fullname := requestBody.Fullname
	email := requestBody.Email

	profile, err := uh.UserService.UserRepository.GetUserByID(profileID)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to get user", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(e.ClientError.Code, errResponse)
		return
	}

	profile.Username = username
	profile.Fullname = fullname
	profile.Email = email

	updatedProfile, err := uh.UserService.UserRepository.UpdateUser(profile)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to update user", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	profileInterface := []interface{}{updatedProfile}
	h := utils.NewResponse(http.StatusOK, "User updated successfully", profileInterface)
	c.JSON(http.StatusOK, h)
}