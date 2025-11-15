package mocks

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/stretchr/testify/mock"
)

type MockMarketplaceRepository struct {
	mock.Mock
}

func (m *MockMarketplaceRepository) Save(ctx context.Context, marketplace domains.MarketplaceCredential) (domains.MarketplaceCredential, error) {
	args := m.Called(ctx, marketplace)
	return args.Get(0).(domains.MarketplaceCredential), args.Error(1)
}

func (m *MockMarketplaceRepository) GetByUserIdAndPlatform(ctx context.Context, userId int64, platform string) (domains.MarketplaceCredential, error) {
	args := m.Called(ctx, userId, platform)
	return args.Get(0).(domains.MarketplaceCredential), args.Error(1)
}

func (m *MockMarketplaceRepository) DeleteByUserIdAndPlatform(ctx context.Context, userId int64, platform string) error {
	args := m.Called(ctx, userId, platform)
	return args.Error(0)
}
