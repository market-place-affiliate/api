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

// CreateLink godoc
// @Summary Create affiliate link
// @Description Create a new affiliate link for a product and campaign
// @Tags link
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateLinkRequest true "Link request"
// @Success 200 {object} dto.LinkResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {object} dto.EmptyResponse "Forbidden"
// @Router /link [post]
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

// GetLinkById godoc
// @Summary Get link by ID
// @Description Get a specific link by its ID
// @Tags link
// @Produce json
// @Param link_id path string true "Link ID"
// @Success 200 {object} dto.LinkResponse
// @Failure 500 {object} dto.EmptyResponse
// @Router /link/{link_id} [get]
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

// GetLinkByShortCode godoc
// @Summary Get link by short code
// @Description Get a link by its short code
// @Tags link
// @Produce json
// @Param short_code path string true "Short code"
// @Success 200 {object} dto.LinkResponse
// @Failure 500 {object} dto.EmptyResponse
// @Router /link/short-code/{short_code} [get]
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

// RedirectLink godoc
// @Summary Redirect to affiliate link
// @Description Track click and redirect to the marketplace affiliate link
// @Tags link
// @Param short_code path string true "Short code"
// @Success 302 {string} string "Redirect to affiliate URL"
// @Router /link/redirect/{short_code} [get]
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

// GetLinksByCampaign godoc
// @Summary Get links by campaign
// @Description Get all links associated with a campaign
// @Tags link
// @Produce json
// @Param campaignId path string true "Campaign ID"
// @Success 200 {object} dto.LinksResponse
// @Failure 500 {object} dto.EmptyResponse
// @Router /link/campaign/{campaignId} [get]
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

// DeleteLink godoc
// @Summary Delete link
// @Description Delete a link and all associated clicks
// @Tags link
// @Produce json
// @Security BearerAuth
// @Param link_id path string true "Link ID"
// @Success 200 {object} dto.EmptyResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {object} dto.EmptyResponse "Forbidden"
// @Router /link/{link_id} [delete]
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
