package handler

import "github.com/gin-gonic/gin"

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) Index(c *gin.Context) {
	c.HTML(200, "user_index.html", nil)
}
