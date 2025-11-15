package db

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"gorm.io/gorm"
)

type linkRepository struct {
	DB *gorm.DB
}

func NewLinkRepository(db *gorm.DB) ports.LinkRepository {
	return &linkRepository{DB: db}
}

func (r *linkRepository) SaveLink(ctx context.Context, link domains.Link) (domains.Link, error) {
	err := r.DB.Save(&link).Error
	if err != nil {
		return domains.Link{}, err
	}
	return link, nil
}
func (r *linkRepository) DeleteLink(ctx context.Context, linkId string) error {
	err := r.DB.Delete(&domains.Link{}, "id = ?", linkId).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *linkRepository) GetLinksByProductId(ctx context.Context, productId string) ([]domains.Link, error) {
	var links []domains.Link
	err := r.DB.Find(&links, "product_id = ?", productId).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}
func (r *linkRepository) GetLinkById(ctx context.Context, linkId string) (domains.Link, error) {
	var link domains.Link
	err := r.DB.First(&link, "id = ?", linkId).Error
	if err != nil {
		return domains.Link{}, err
	}
	return link, nil
}
func (r *linkRepository) GetLinkByShortCode(ctx context.Context, shortCode string) (domains.Link, error) {
	var link domains.Link
	err := r.DB.First(&link, "short_code = ?", shortCode).Error
	if err != nil {
		return domains.Link{}, err
	}
	return link, nil
}

func (r *linkRepository) GetLinksByCampaignId(ctx context.Context, campaignId string) ([]domains.Link, error) {
	var links []domains.Link
	err := r.DB.Find(&links, "campaign_id = ?", campaignId).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (r *linkRepository) DeleteLinkByProductId(ctx context.Context, productId string) error {
	err := r.DB.Delete(&domains.Link{}, "product_id = ?", productId).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *linkRepository) DeleteLinkByCampaignId(ctx context.Context, campaignId string) error {
	err := r.DB.Delete(&domains.Link{}, "campaign_id = ?", campaignId).Error
	if err != nil {
		return err
	}
	return nil
}
