package mocks

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/stretchr/testify/mock"
)

type MockClickRepository struct {
	mock.Mock
}

func (m *MockClickRepository) SaveClick(ctx context.Context, click domains.Click) error {
	args := m.Called(ctx, click)
	return args.Error(0)
}

func (m *MockClickRepository) CountClicksByDateRange(ctx context.Context, userId int64, startDate, endDate time.Time) ([]dto.MetrictItem, error) {
	args := m.Called(ctx, userId, startDate, endDate)
	return args.Get(0).([]dto.MetrictItem), args.Error(1)
}

func (m *MockClickRepository) CountTopProductClickByDateRange(ctx context.Context, userId int64, startDate, endDate time.Time) (uuid.UUID, int64, error) {
	args := m.Called(ctx, userId, startDate, endDate)
	return args.Get(0).(uuid.UUID), args.Get(1).(int64), args.Error(2)
}

func (m *MockClickRepository) DeleteClicksByLinkId(ctx context.Context, linkId string) error {
	args := m.Called(ctx, linkId)
	return args.Error(0)
}
