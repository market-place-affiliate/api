package db

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user domains.User) (domains.User, error) {
	err := r.DB.Create(&user).Error
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userId int64) (domains.User, error) {
	var user domains.User
	err := r.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (domains.User, error) {
	var user domains.User
	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}
