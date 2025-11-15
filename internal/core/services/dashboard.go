package services

import (
	"context"
	"net/http"
	"time"

	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
)

type dashboardService struct {
	clickRepo   ports.ClickRepository
	productRepo ports.ProductRepository
}

func NewDashboardService(clickRepo ports.ClickRepository, productRepo ports.ProductRepository) ports.DashboardService {
	return &dashboardService{clickRepo: clickRepo, productRepo: productRepo}
}

func (s *dashboardService) GetDashboardMetrics(ctx context.Context, userId int64, startDate, endDate time.Time) (dto.Response[dto.DashboardMetricsResponse], error) {
	metricts, err := s.clickRepo.CountClicksByDateRange(ctx, userId, startDate, endDate)
	if err != nil {
		return dto.Response[dto.DashboardMetricsResponse]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4001,
			Message:  "Failed to count clicks by date range",
		}, err
	}
	productId, clickCount, err := s.clickRepo.CountTopProductClickByDateRange(ctx, userId, startDate, endDate)
	if err != nil {
		return dto.Response[dto.DashboardMetricsResponse]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4002,
			Message:  "Failed to get top product clicks by date range",
		}, err
	}
	product, err := s.productRepo.GetProductById(ctx, productId.String())
	if err != nil {
		return dto.Response[dto.DashboardMetricsResponse]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4003,
			Message:  "Failed to get product by id",
		}, err
	}
	topProduct := dto.TopProduct{
		Product: product,
		Clicks:  clickCount,
	}
	return dto.Response[dto.DashboardMetricsResponse]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Data: dto.DashboardMetricsResponse{
			Metrics:    metricts,
			TopProduct: topProduct,
		},
	}, nil
}
