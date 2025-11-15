package db

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"gorm.io/gorm"
)

type clickRepository struct {
	DB *gorm.DB
}

func NewClickRepository(db *gorm.DB) ports.ClickRepository {
	return &clickRepository{DB: db}
}

func (r *clickRepository) SaveClick(ctx context.Context, click domains.Click) error {
	err := r.DB.Create(&click).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *clickRepository) CountClicksByDateRange(ctx context.Context, userId int64, startDate, endDate time.Time) ([]dto.MetrictItem, error) {
	var results []dto.MetrictItem
	err := r.DB.Debug().Raw(`
	select 
	date(clicks.created_at) as date, 
	count(*) as click_count,
	links.campaign_id,
	campaigns.name as campaign_name,
	offers.marketplace
	from clicks
	left join links on clicks.link_id = links.id
	left join campaigns on links.campaign_id = campaigns.id
	left join offers on offers.product_id = links.product_id
	where user_id = ? and clicks.created_at >= ? and clicks.created_at <= ?
	group by date(clicks.created_at), links.campaign_id,campaigns.name, offers.marketplace
	order by date(clicks.created_at) asc
	`, userId, startDate, endDate,
	).Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r *clickRepository) CountTopProductClickByDateRange(ctx context.Context, userId int64, startDate, endDate time.Time) (uuid.UUID, int64, error) {
	type results struct {
		ProductId  uuid.UUID 
		ClickCount int64
	}
	var result results
	err := r.DB.Raw(
		`
	select 
		products.id as product_id, 
		count(*) as click_count
	from clicks
	left join links on clicks.link_id = links.id
	left join products on links.product_id = products.id
	where user_id = ? and clicks.created_at >= ? and clicks.created_at <= ?
	group by products.id
	order by count(*) desc
	limit 1
	`, userId, startDate, endDate,
	).Scan(&result).Error
	if err != nil {
		return uuid.Nil, 0, err
	}
	return result.ProductId, result.ClickCount, nil
}

func (r *clickRepository) DeleteClicksByLinkId(ctx context.Context, linkId string) error {
	err := r.DB.Delete(&domains.Click{}, "link_id = ?", linkId).Error
	if err != nil {
		return err
	}
	return nil
}