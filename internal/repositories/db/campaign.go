package db

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"gorm.io/gorm"
)

type campaignRepository struct {
	DB *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) ports.CampaignRepository {
	return &campaignRepository{DB: db}
}

func (r *campaignRepository) SaveCampaign(ctx context.Context, campaign domains.Campaign) (domains.Campaign, error) {
	err := r.DB.Save(&campaign).Error
	if err != nil {
		return domains.Campaign{}, err
	}
	return campaign, nil
}
func (r *campaignRepository) DeleteCampaign(ctx context.Context, campaignId string) error {
	err := r.DB.Delete(&domains.Campaign{}, "id = ?", campaignId).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *campaignRepository) GetCampaignById(ctx context.Context, campaignId string) (domains.Campaign, error) {
	var campaign domains.Campaign
	err := r.DB.First(&campaign, "id = ?", campaignId).Error
	if err != nil {
		return domains.Campaign{}, err
	}
	return campaign, nil
}
func (r *campaignRepository) GetCampaignByQuery(ctx context.Context, userId int64, query dto.GetCampaignByQueryRequest) ([]domains.Campaign, error) {
	var campaigns []domains.Campaign
	dbQuery := r.DB.Model(&domains.Campaign{})
	if userId != 0 {
		dbQuery = dbQuery.Where("user_id = ?", userId)
	}
	if query.Name != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if !query.StartAt.IsZero(){
		dbQuery = dbQuery.Where("start_at >= ?", query.StartAt)
	}
	if !query.EndAt.IsZero() {
		dbQuery = dbQuery.Where("end_at <= ?", query.EndAt)
	}
	err := dbQuery.Debug().Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (r *campaignRepository) GetAvailableCampaign(ctx context.Context) ([]domains.Campaign, error) {
	var campaigns []domains.Campaign
	dbQuery := r.DB.Model(&domains.Campaign{})
	dbQuery.Where("date(start_at) <= ? AND date(end_at) >= ?",  gorm.Expr("date(NOW())"), gorm.Expr("date(NOW())"))
	err := dbQuery.Debug().Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}