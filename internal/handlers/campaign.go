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
