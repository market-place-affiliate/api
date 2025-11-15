package dto

import (
	"time"

	"github.com/gofrs/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type CreateProductRequest struct {
	SourceUrl   string `json:"source_url" binding:"required,url"`
	Marketplace string `json:"marketplace" binding:"required,oneof=shopee lazada"`
}

type CreateCampaignRequest struct {
	Name        string    `json:"name" binding:"required,min=3,max=100"`
	UtmCampaign string    `json:"utm_campaign" binding:"required,min=3,max=100"`
	StartAt     time.Time `json:"start_at" binding:"required"`
	EndAt       time.Time `json:"end_at" binding:"required,gtefield=StartAt"`
}

type CreateLinkRequest struct {
	ProductId  uuid.UUID `json:"product_id" binding:"required,uuid"`
	CampaignId uuid.UUID `json:"campaign_id" binding:"required,uuid"`
}

type GetCampaignByQueryRequest struct {
	Name    string    `form:"name" binding:"omitempty,min=3,max=100"`
	StartAt time.Time `form:"start_at" binding:"omitempty"`
	EndAt   time.Time `form:"end_at" binding:"omitempty,gtfield=StartAt"`

	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

type MarketplaceCredentialRequest struct {
	Platform   string `json:"platform" binding:"required,oneof=shopee lazada"`
	AppKey     string `json:"app_key"`
	SignMethod string `json:"sign_method"`
	UserToken  string `json:"user_token"`
	AppId      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
}
