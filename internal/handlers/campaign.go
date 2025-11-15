package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
)

type CampaignHandler struct {
	campaignService ports.CampaignService
}

func NewCampaignHandler(campaignService ports.CampaignService) *CampaignHandler {
	return &CampaignHandler{campaignService: campaignService}
}

// CreateCampaign godoc
// @Summary Create a new campaign
// @Description Create a new marketing campaign
// @Tags campaign
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateCampaignRequest true "Campaign request"
// @Success 200 {object} dto.CampaignResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /campaign [post]
func (h *CampaignHandler) CreateCampaign(g *gin.Context) {
	ctx := g.Request.Context()
	body := dto.CreateCampaignRequest{}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.AbortWithStatus(400)
		return
	}
	userId := g.GetInt64("userId")

	res, err := h.campaignService.CreateCampaign(ctx, userId, body)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

// GetCampaigns godoc
// @Summary Get user campaigns
// @Description Get all campaigns for the authenticated user with optional filters
// @Tags campaign
// @Produce json
// @Security BearerAuth
// @Param utm_campaign query string false "UTM campaign filter"
// @Success 200 {object} dto.CampaignsResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /campaign [get]
func (h *CampaignHandler) GetCampaigns(g *gin.Context) {
	ctx := g.Request.Context()
	userId := g.GetInt64("userId")
	body := dto.GetCampaignByQueryRequest{}
	if err := g.ShouldBindQuery(&body); err != nil {
		g.AbortWithStatus(400)
		return
	}
	res, err := h.campaignService.GetCampaignByQuery(ctx, userId, body)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

// DeleteCampaign godoc
// @Summary Delete campaign
// @Description Delete a campaign and all associated links and clicks
// @Tags campaign
// @Produce json
// @Security BearerAuth
// @Param campaign_id path string true "Campaign ID"
// @Success 200 {object} dto.EmptyResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {object} dto.EmptyResponse "Forbidden"
// @Router /campaign/{campaign_id} [delete]
func (h *CampaignHandler) DeleteCampaign(g *gin.Context) {
	ctx := g.Request.Context()
	userId := g.GetInt64("userId")
	campaignId := g.Param("campaign_id")
	if campaignId == "" {
		g.AbortWithStatus(400)
		return
	}
	res, err := h.campaignService.DeleteCampaignById(ctx, userId, campaignId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

// GetPublicCampaigns godoc
// @Summary Get public campaigns
// @Description Get all public campaigns available for everyone
// @Tags campaign
// @Produce json
// @Param utm_campaign query string false "UTM campaign filter"
// @Success 200 {object} dto.CampaignsResponse
// @Failure 400 {string} string "Bad Request"
// @Router /campaign/available [get]
func (h *CampaignHandler) GetPublicCampaigns(g *gin.Context) {
	ctx := g.Request.Context()
	body := dto.GetCampaignByQueryRequest{}
	if err := g.ShouldBindQuery(&body); err != nil {
		g.AbortWithStatus(400)
		return
	}
	res, err := h.campaignService.GetPublicCampaigns(ctx, body)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}
