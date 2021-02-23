package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Registere account failed", 422, "Failed", errorMessage)
		c.JSON(422, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Registere account failed", 400, "Failed", nil)
		c.JSON(400, response)
		return
	}

	formatter := user.FormatUser(newUser, "iniadalahtoken")

	response := helper.APIResponse("Account has been registered", 200, "Success", formatter)

	c.JSON(200, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", 422, "Failed", errorMessage)
		c.JSON(422, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", 400, "Failed", errorMessage)
		c.JSON(400, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "iniadalahtoken")

	response := helper.APIResponse("Successfully loggedin", 200, "Success", formatter)

	c.JSON(200, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", 422, "Failed", errorMessage)
		c.JSON(422, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Email checking failed", 422, "Failed", errorMessage)
		c.JSON(422, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, 200, "Success", data)
	c.JSON(200, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", 400, "Failed", data)

		c.JSON(400, response)
		return
	}

	userID := 1
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", 400, "Failed", data)

		c.JSON(400, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", 400, "Failed", data)

		c.JSON(400, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfully uploaded", 200, "Success", data)

	c.JSON(200, response)
}
