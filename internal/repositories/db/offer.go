package db

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"gorm.io/gorm"
)

type offerRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(db *gorm.DB) ports.OfferRepository {
	return &offerRepository{DB: db}
}

func (r *offerRepository) SaveOffer(ctx context.Context, offer domains.Offer) error {
	err := r.DB.Save(&offer).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *offerRepository) DeleteOffer(ctx context.Context, offerId string) error {
	err := r.DB.Delete(&domains.Offer{}, "id = ?", offerId).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *offerRepository) GetOffersByProductId(ctx context.Context, productId string) (domains.Offer, error) {
	var offer domains.Offer
	err := r.DB.First(&offer, "product_id = ?", productId).Error
	if err != nil {
		return domains.Offer{}, err
	}
	return offer, nil
}
func (r *offerRepository) GetOfferById(ctx context.Context, offerId string) (domains.Offer, error) {
	var offer domains.Offer
	err := r.DB.First(&offer, "id = ?", offerId).Error
	if err != nil {
		return domains.Offer{}, err
	}
	return offer, nil
}
func (r *offerRepository) DeleteOfferByProductId(ctx context.Context, productId string) error {
	err := r.DB.Delete(&domains.Offer{}, "product_id = ?", productId).Error
	if err != nil {
		return err
	}
	return nil
}