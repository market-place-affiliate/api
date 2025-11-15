package mocks

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/stretchr/testify/mock"
)

type MockOfferRepository struct {
	mock.Mock
}

func (m *MockOfferRepository) SaveOffer(ctx context.Context, offer domains.Offer) error {
	args := m.Called(ctx, offer)
	return args.Error(0)
}

func (m *MockOfferRepository) DeleteOffer(ctx context.Context, offerId string) error {
	args := m.Called(ctx, offerId)
	return args.Error(0)
}

func (m *MockOfferRepository) GetOffersByProductId(ctx context.Context, productId string) (domains.Offer, error) {
	args := m.Called(ctx, productId)
	return args.Get(0).(domains.Offer), args.Error(1)
}

func (m *MockOfferRepository) GetOfferById(ctx context.Context, offerId string) (domains.Offer, error) {
	args := m.Called(ctx, offerId)
	return args.Get(0).(domains.Offer), args.Error(1)
}

func (m *MockOfferRepository) DeleteOfferByProductId(ctx context.Context, productId string) error {
	args := m.Called(ctx, productId)
	return args.Error(0)
}
