package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateToken(NewUser.Id)
	if err != nil {
		response := helper.ApiResponse("register faile", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := user.FormastUser(NewUser, token)
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
	token, err := h.authService.GenerateToken(logedUser.Id)
	if err != nil {
		response := helper.ApiResponse("register faile", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := user.FormastUser(logedUser, token)

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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	//tangkap input dari user
	//simpan gambar di folder "/images"
	//di service manggil repo
	//jwt(hardcore, seakan akan user yg login ID = 1)
	//repo ambil data user = 1
	//repo update data user simpan lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Avatar faile uploaded", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id

	//path := "images/" + file.Filename
	path := fmt.Sprint("images/", userId, "-", file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Avatar faile uploaded", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Avatar faile uploaded", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Avatar has uploaded", http.StatusOK, "sukses", data)
	c.JSON(http.StatusBadRequest, response)

}
