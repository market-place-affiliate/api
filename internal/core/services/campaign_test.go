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
	"github.com/stretchr/testify/mock"
)

func TestCreateCampaign(t *testing.T) {
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewCampaignService(mockCampaignRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	now := time.Now()
	campaignId := uuid.Must(uuid.NewV4())

	request := dto.CreateCampaignRequest{
		Name:        "Test Campaign",
		UtmCampaign: "test_utm",
		StartAt:     now,
		EndAt:       now.Add(24 * time.Hour),
	}

	expectedCampaign := domains.Campaign{
		Id:          campaignId,
		Name:        request.Name,
		UtmCampaign: request.UtmCampaign,
		StartAt:     request.StartAt,
		EndAt:       request.EndAt,
		UserId:      userId,
	}

	mockCampaignRepo.On("SaveCampaign", ctx, mock.MatchedBy(func(c domains.Campaign) bool {
		return c.Name == request.Name && c.UserId == userId
	})).Return(expectedCampaign, nil)

	result, err := service.CreateCampaign(ctx, userId, request)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, expectedCampaign, result.Data)
	mockCampaignRepo.AssertExpectations(t)
}

func TestGetCampaignByQuery(t *testing.T) {
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewCampaignService(mockCampaignRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	query := dto.GetCampaignByQueryRequest{
		Name: "Test",
	}

	expectedCampaigns := []domains.Campaign{
		{
			Id:          uuid.Must(uuid.NewV4()),
			Name:        "Test Campaign",
			UtmCampaign: "test",
			UserId:      userId,
		},
	}

	mockCampaignRepo.On("GetCampaignByQuery", ctx, userId, query).Return(expectedCampaigns, nil)

	result, err := service.GetCampaignByQuery(ctx, userId, query)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, expectedCampaigns, result.Data)
	mockCampaignRepo.AssertExpectations(t)
}

func TestDeleteCampaignById_Success(t *testing.T) {
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewCampaignService(mockCampaignRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	campaignId := uuid.Must(uuid.NewV4())
	linkId := uuid.Must(uuid.NewV4())

	campaign := domains.Campaign{
		Id:     campaignId,
		UserId: userId,
		Name:   "Test Campaign",
	}

	links := []domains.Link{
		{
			Id:         linkId,
			CampaignId: campaignId,
		},
	}

	mockCampaignRepo.On("GetCampaignById", ctx, campaignId.String()).Return(campaign, nil)
	mockLinkRepo.On("GetLinksByCampaignId", ctx, campaignId.String()).Return(links, nil)
	mockClickRepo.On("DeleteClicksByLinkId", ctx, linkId.String()).Return(nil)
	mockLinkRepo.On("DeleteLink", ctx, linkId.String()).Return(nil)
	mockCampaignRepo.On("DeleteCampaign", ctx, campaignId.String()).Return(nil)

	result, err := service.DeleteCampaignById(ctx, userId, campaignId.String())

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	mockCampaignRepo.AssertExpectations(t)
	mockLinkRepo.AssertExpectations(t)
	mockClickRepo.AssertExpectations(t)
}

func TestDeleteCampaignById_Forbidden(t *testing.T) {
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewCampaignService(mockCampaignRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	campaignId := uuid.Must(uuid.NewV4())

	campaign := domains.Campaign{
		Id:     campaignId,
		UserId: int64(2), // Different user
		Name:   "Test Campaign",
	}

	mockCampaignRepo.On("GetCampaignById", ctx, campaignId.String()).Return(campaign, nil)

	result, err := service.DeleteCampaignById(ctx, userId, campaignId.String())

	assert.NoError(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, 3005, result.Code)
	assert.Equal(t, "You do not have permission to delete this campaign", result.Message)
	mockCampaignRepo.AssertExpectations(t)
}

func TestGetPublicCampaigns(t *testing.T) {
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewCampaignService(mockCampaignRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	query := dto.GetCampaignByQueryRequest{}

	expectedCampaigns := []domains.Campaign{
		{
			Id:   uuid.Must(uuid.NewV4()),
			Name: "Public Campaign",
		},
	}

	mockCampaignRepo.On("GetAvailableCampaign", ctx).Return(expectedCampaigns, nil)

	result, err := service.GetPublicCampaigns(ctx, query)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, expectedCampaigns, result.Data)
	mockCampaignRepo.AssertExpectations(t)
}
