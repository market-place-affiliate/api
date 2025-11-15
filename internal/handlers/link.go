package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
)

type LinkHandler struct {
	linkService ports.LinkService
}

func NewLinkHandler(linkService ports.LinkService) *LinkHandler {
	return &LinkHandler{linkService: linkService}
}

func (h *LinkHandler) CreateLink(g *gin.Context) {
	ctx := g.Request.Context()
	body := dto.CreateLinkRequest{}
	userId := g.GetInt64("userId")
	if err := g.ShouldBindJSON(&body); err != nil {
		g.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.linkService.CreateLink(ctx, userId, body)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

func (h *LinkHandler) GetLinkById(g *gin.Context) {
	ctx := g.Request.Context()
	linkId := g.Param("link_id")
	res, err := h.linkService.GetLinkById(ctx, linkId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

func (h *LinkHandler) GetLinkByShortCode(g *gin.Context) {
	ctx := g.Request.Context()
	code := g.Param("short_code")
	res, err := h.linkService.GetLinkByShortCode(ctx, code)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

func (h *LinkHandler) RedirectLink(g *gin.Context) {
	ctx := g.Request.Context()
	code := g.Param("short_code")
	res, err := h.linkService.ClickByShortCode(ctx, code)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.Redirect(http.StatusFound, res.Data.TargetURL)
}

func (h *LinkHandler) GetLinksByCampaign(g *gin.Context) {
	ctx := g.Request.Context()
	campaignId := g.Param("campaignId")
	res, err := h.linkService.GetLinkByCampaign(ctx, campaignId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

func (h *LinkHandler) DeleteLink(g *gin.Context) {
	ctx := g.Request.Context()
	userId := g.GetInt64("userId")
	linkId := g.Param("link_id")
	if linkId == "" {
		g.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := h.linkService.DeleteLinkById(ctx, userId, linkId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}
