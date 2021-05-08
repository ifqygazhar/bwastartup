package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

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
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("register faile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	NewUser, err := h.userService.RegisterUserInput(input)
	if err != nil {
		response := helper.ApiResponse("register faile", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formater := user.FormastUser(NewUser, "tokentokentoken ")
	response := helper.ApiResponse("account has been register", http.StatusOK, "sukses", formater)
	c.JSON(http.StatusOK, response)
}
