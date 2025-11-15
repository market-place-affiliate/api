package dto

import (
	"github.com/gofrs/uuid"
	"github.com/market-place-affiliate/api/internal/core/domains"
)

type DashboardMetricsResponse struct {
	TopProduct TopProduct    `json:"top_product"`
	Metrics    []MetrictItem `json:"metrics"`
}

type MetrictItem struct {
	Date         string    `json:"date" gorm:"column:date"`
	ClickCount   int       `json:"click_count" gorm:"column:click_count"`
	CampaignId   uuid.UUID `json:"campaign" gorm:"column:campaign_id"`
	CampaignName string    `json:"campaign_name" gorm:"column:name;type:text;not null"`
	Marketplace  string    `json:"marketplace" gorm:"column:marketplace"`
}

type TopProduct struct {
	Product domains.Product `json:"product" `
	Clicks  int64           `json:"clicks"`
}
