package handler

import (
	"bwastartup/helper"
	"bwastartup/user"

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
