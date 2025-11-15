package mocks

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/stretchr/testify/mock"
)

type MockCampaignRepository struct {
	mock.Mock
}

func (m *MockCampaignRepository) SaveCampaign(ctx context.Context, campaign domains.Campaign) (domains.Campaign, error) {
	args := m.Called(ctx, campaign)
	return args.Get(0).(domains.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) DeleteCampaign(ctx context.Context, campaignId string) error {
	args := m.Called(ctx, campaignId)
	return args.Error(0)
}

func (m *MockCampaignRepository) GetCampaignById(ctx context.Context, campaignId string) (domains.Campaign, error) {
	args := m.Called(ctx, campaignId)
	return args.Get(0).(domains.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) GetCampaignByQuery(ctx context.Context, userId int64, query dto.GetCampaignByQueryRequest) ([]domains.Campaign, error) {
	args := m.Called(ctx, userId, query)
	return args.Get(0).([]domains.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) GetAvailableCampaign(ctx context.Context) ([]domains.Campaign, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domains.Campaign), args.Error(1)
}
