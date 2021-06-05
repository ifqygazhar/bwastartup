package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap paramemter di handler
// handler ke service
// service yang menentukan repository mana  yang di call
// repository baru dia ngakses ke db
// repository : GetAll, GetByUserId

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// api/v1/campaign

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.ApiResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of campaign", http.StatusOK, "succes", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// api/v1/campaign/2
	// handler : mapping id yang di url ke sturct input => service , call formatter
	// service : input nya struct input  => mensngkap id di url, manggil repo get campaign id
	//repository : get campaign id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Error to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Error to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Detail campaign", http.StatusOK, "succes", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)

}

// tangkap parameter dari user ke input struct
// ambil current user dari jwt/handler
// pangil service paramter nya input struct (dan juga buat slug)
// panggil repository untuk simpan data campaign baru
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Fail to create campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)

	if err != nil {
		response := helper.ApiResponse("Fail to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.ApiResponse("succes to create campaign", http.StatusOK, "succes", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}

//user masukan input
// handler menangkap input
//mapping dari input ke input struct (ada 2)
//service butuh parameter input dari user dan uri
//repository update data campaign

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputId campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.ApiResponse("Error to update detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Fail to create campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {
		response := helper.ApiResponse("Error to update detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("succes to update campaign", http.StatusOK, "succes", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)

}

// handler tangkap input dan rubah ke struct input
// save image
// service (kondisi msnggil 2 point di repo, panggil repo point satu)
// repository : create_image/save ke dalam campaign images, ubah is_primary true ke false
func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Fail to upload campaign image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser
	userId := currentUser.Id

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("faile to uploade campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//path := "images/" + file.Filename
	path := fmt.Sprint("images/", userId, "-", file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("faile uploade image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.SaveCampaignimage(input, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("faile uploaded campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Campaign image has uploaded", http.StatusOK, "sukses", data)
	c.JSON(http.StatusBadRequest, response)

}
