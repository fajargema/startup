package handler

import (
	"bwastartup/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(500, "error.html", nil)
		return
	}
	c.HTML(200, "user_index.html", gin.H{"users": users})
}

func (h *userHandler) New(c *gin.Context) {
	c.HTML(200, "user_new.html", nil)
}

func (h *userHandler) Create(c *gin.Context) {
	var input user.FormCreateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(200, "user_new.html", input)
		return
	}

	registerInput := user.RegisterUserInput{}
	registerInput.Name = input.Name
	registerInput.Email = input.Email
	registerInput.Occupation = input.Occupation
	registerInput.Password = input.Password

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		c.HTML(500, "error.html", nil)
		return
	}

	c.Redirect(302, "/users")
}

func (h *userHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.userService.GetUserById(id)
	if err != nil {
		c.HTML(500, "error.html", nil)
	}
	c.HTML(200, "user_edit.html", registeredUser)
}

func (h *userHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input user.FormUpdateUserInput
	err := c.ShouldBind(&input)
	if err != nil {

	}

	input.ID = id

	_, err = h.userService.UpdateUser(input)
	if err != nil {
		c.HTML(500, "error.html", nil)
	}

	c.Redirect(302, "/users")
}
