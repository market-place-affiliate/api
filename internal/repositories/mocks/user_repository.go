package mocks

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user domains.User) (domains.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(domains.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userId int64) (domains.User, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(domains.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (domains.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(domains.User), args.Error(1)
}
