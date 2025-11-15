package db

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type marketplaceCredentialRepository struct {
	DB *gorm.DB
}

func NewMarketplaceCredentialRepository(db *gorm.DB) ports.MarketplaceRepository {
	return &marketplaceCredentialRepository{DB: db}
}

func (r *marketplaceCredentialRepository) Save(ctx context.Context, marketplace domains.MarketplaceCredential) (domains.MarketplaceCredential, error) {
	err := r.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "marketplace"}},
		DoUpdates: clause.AssignmentColumns([]string{"app_id", "app_key", "app_secret", "user_token", "updated_at"}),
	}).Create(&marketplace).Error
	if err != nil {
		return domains.MarketplaceCredential{}, err
	}
	return marketplace, nil
}

func (r *marketplaceCredentialRepository) GetByUserIdAndPlatform(ctx context.Context, userId int64, platform string) (domains.MarketplaceCredential, error) {
	var marketplace domains.MarketplaceCredential
	err := r.DB.First(&marketplace, "user_id = ? AND marketplace = ?", userId, platform).Error
	if err != nil {
		return domains.MarketplaceCredential{}, err
	}
	return marketplace, nil
}

func (r *marketplaceCredentialRepository) DeleteByUserIdAndPlatform(ctx context.Context, userId int64, platform string) error {
	err := r.DB.Delete(&domains.MarketplaceCredential{}, "user_id = ? AND marketplace = ?", userId, platform).Error
	if err != nil {
		return err
	}
	return nil
}
