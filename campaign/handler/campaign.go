package handler

import (
	"crowdfund-go/campaign"
	"crowdfund-go/campaign/request"
	"crowdfund-go/campaign/response"
	"crowdfund-go/entity"
	"crowdfund-go/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) Campaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.Campaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error getting campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", response.FormatMany(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) Create(c *gin.Context) {
	var req request.Create

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	req.User = currentUser

	newCampaign, err := h.service.Create(req)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", response.Format(newCampaign))
	c.JSON(http.StatusOK, response)
}
