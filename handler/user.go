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
		errors := helper.FormatValidationError(err)
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

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan  input (email & password)
	// input di tangkap handler
	// mapping dari input user ke input struct
	// input struct kita pasing ke service
	// di service mencari dengan bantuan repository user dengan email X

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("login faile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	logedUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.ApiResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formater := user.FormastUser(logedUser, "tokentokentoken")

	response := helper.ApiResponse("Sukses login", http.StatusOK, "sukses", formater)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailable(c *gin.Context) {
	// ada input email dari user
	// input email di mapping ke struct input
	// struct input di pasing ke service
	// service akan memanggil repository - email sudah ada atau belom
	// repository akan melakukan query ke database

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Email checking faile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"error": "Server error"}
		response := helper.ApiResponse("Email checking faile", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string
	if isEmailAvailable {
		metaMessage = "email available"
	} else {
		metaMessage = "email has been register"
	}
	response := helper.ApiResponse(metaMessage, http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)

}