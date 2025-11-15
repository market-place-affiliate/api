package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/market-place-affiliate/api/internal/core/ports"
)

type DashboardHandler struct {
	dashboardService ports.DashboardService
}

func NewDashboardHandler(dashboardService ports.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService: dashboardService}
}

// GetDashboardData godoc
// @Summary Get dashboard metrics
// @Description Get dashboard analytics including clicks, products, and performance metrics
// @Tags dashboard
// @Produce json
// @Security BearerAuth
// @Param start_at query string false "Start date (YYYY-MM-DD)" default("7 days ago")
// @Param end_at query string false "End date (YYYY-MM-DD)" default("tomorrow")
// @Success 200 {object} dto.DashboardResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {string} string "Unauthorized"
// @Router /dashboard/metrics [get]
func (h *DashboardHandler) GetDashboardData(g *gin.Context) {
	ctx := g.Request.Context()
	userId := g.GetInt64("userId")
	start := g.Query("start_at")
	end := g.Query("end_at")
	if start == "" {
		start = time.Now().AddDate(0, 0, -7).Format(time.DateOnly)
	}
	if end == "" {
		end = time.Now().AddDate(0, 0, 1).Format(time.DateOnly)
	}
	startTime, err := time.Parse(time.DateOnly, start)
	if err != nil {
		g.JSON(400, gin.H{"error": "Invalid start date format"})
		return
	}
	endTime, err := time.Parse(time.DateOnly, end)
	if err != nil {
		g.JSON(400, gin.H{"error": "Invalid end date format"})
		return
	}
	res, err := h.dashboardService.GetDashboardMetrics(ctx, userId, startTime, endTime)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(200, res)
}
