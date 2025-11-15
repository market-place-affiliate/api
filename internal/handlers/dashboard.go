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

func (h *DashboardHandler) GetDashboardData(g *gin.Context) {
	ctx := g.Request.Context()
	userId := g.GetInt64("userId")
	start := g.Query("start")
	end := g.Query("end")
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
