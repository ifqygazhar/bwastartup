package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
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
	return
}
