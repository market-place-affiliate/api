package mocks

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/stretchr/testify/mock"
)

type MockLinkRepository struct {
	mock.Mock
}

func (m *MockLinkRepository) SaveLink(ctx context.Context, link domains.Link) (domains.Link, error) {
	args := m.Called(ctx, link)
	return args.Get(0).(domains.Link), args.Error(1)
}

func (m *MockLinkRepository) DeleteLink(ctx context.Context, linkId string) error {
	args := m.Called(ctx, linkId)
	return args.Error(0)
}

func (m *MockLinkRepository) GetLinksByProductId(ctx context.Context, productId string) ([]domains.Link, error) {
	args := m.Called(ctx, productId)
	return args.Get(0).([]domains.Link), args.Error(1)
}

func (m *MockLinkRepository) GetLinkById(ctx context.Context, linkId string) (domains.Link, error) {
	args := m.Called(ctx, linkId)
	return args.Get(0).(domains.Link), args.Error(1)
}

func (m *MockLinkRepository) GetLinkByShortCode(ctx context.Context, shortCode string) (domains.Link, error) {
	args := m.Called(ctx, shortCode)
	return args.Get(0).(domains.Link), args.Error(1)
}

func (m *MockLinkRepository) GetLinksByCampaignId(ctx context.Context, campaignId string) ([]domains.Link, error) {
	args := m.Called(ctx, campaignId)
	return args.Get(0).([]domains.Link), args.Error(1)
}

func (m *MockLinkRepository) DeleteLinkByProductId(ctx context.Context, productId string) error {
	args := m.Called(ctx, productId)
	return args.Error(0)
}

func (m *MockLinkRepository) DeleteLinkByCampaignId(ctx context.Context, campaignId string) error {
	args := m.Called(ctx, campaignId)
	return args.Error(0)
}
