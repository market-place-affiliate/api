package services

import (
	"context"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetDashboardMetrics_Success(t *testing.T) {
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)

	service := NewDashboardService(mockClickRepo, mockProductRepo)

	ctx := context.Background()
	userId := int64(1)
	startDate := time.Now().AddDate(0, 0, -7)
	endDate := time.Now()
	productId := uuid.Must(uuid.NewV4())

	metrics := []dto.MetrictItem{
		{
			Date:         time.Now().Format("2006-01-02"),
			ClickCount:   10,
			CampaignId:   uuid.Must(uuid.NewV4()),
			CampaignName: "Test Campaign",
			Marketplace:  "lazada",
		},
	}

	product := domains.Product{
		Id:       productId,
		Title:    "Test Product",
		ImageUrl: "https://example.com/image.jpg",
		UserId:   userId,
	}

	mockClickRepo.On("CountClicksByDateRange", ctx, userId, startDate, endDate).Return(metrics, nil)
	mockClickRepo.On("CountTopProductClickByDateRange", ctx, userId, startDate, endDate).Return(productId, int64(100), nil)
	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)

	result, err := service.GetDashboardMetrics(ctx, userId, startDate, endDate)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, metrics, result.Data.Metrics)
	assert.Equal(t, product, result.Data.TopProduct.Product)
	assert.Equal(t, int64(100), result.Data.TopProduct.Clicks)
	mockClickRepo.AssertExpectations(t)
	mockProductRepo.AssertExpectations(t)
}
