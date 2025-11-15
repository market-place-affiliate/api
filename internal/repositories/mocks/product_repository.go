package mocks

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) SaveProduct(ctx context.Context, product domains.Product) (domains.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(domains.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(ctx context.Context, productId string) error {
	args := m.Called(ctx, productId)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductById(ctx context.Context, productId string) (domains.Product, error) {
	args := m.Called(ctx, productId)
	return args.Get(0).(domains.Product), args.Error(1)
}

func (m *MockProductRepository) GetAllProducts(ctx context.Context, userId int64) ([]domains.Product, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]domains.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProductById(ctx context.Context, productId string) error {
	args := m.Called(ctx, productId)
	return args.Error(0)
}
