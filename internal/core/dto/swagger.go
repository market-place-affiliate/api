package dto

import "github.com/market-place-affiliate/api/internal/core/domains"

// Swagger response types for documentation

// StringResponse represents a response with string data
type StringResponse struct {
	Success bool   `json:"success" example:"true"`
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"Operation successful"`
	TxnID   string `json:"txn_id" example:"txn_123456"`
	Data    string `json:"data,omitempty" example:"data_value"`
}

// UserResponse represents a response with user data
type UserResponse struct {
	Success bool         `json:"success" example:"true"`
	Code    int          `json:"code" example:"0"`
	Message string       `json:"message" example:"User retrieved successfully"`
	TxnID   string       `json:"txn_id" example:"txn_123456"`
	Data    domains.User `json:"data,omitempty"`
}

// ProductResponse represents a response with product data
type ProductResponse struct {
	Success bool            `json:"success" example:"true"`
	Code    int             `json:"code" example:"0"`
	Message string          `json:"message" example:"Product retrieved successfully"`
	TxnID   string          `json:"txn_id" example:"txn_123456"`
	Data    domains.Product `json:"data,omitempty"`
}

// ProductsResponse represents a response with product array
type ProductsResponse struct {
	Success bool              `json:"success" example:"true"`
	Code    int               `json:"code" example:"0"`
	Message string            `json:"message" example:"Products retrieved successfully"`
	TxnID   string            `json:"txn_id" example:"txn_123456"`
	Data    []domains.Product `json:"data,omitempty"`
}

// CampaignResponse represents a response with campaign data
type CampaignResponse struct {
	Success bool             `json:"success" example:"true"`
	Code    int              `json:"code" example:"0"`
	Message string           `json:"message" example:"Campaign created successfully"`
	TxnID   string           `json:"txn_id" example:"txn_123456"`
	Data    domains.Campaign `json:"data,omitempty"`
}

// CampaignsResponse represents a response with campaign array
type CampaignsResponse struct {
	Success bool               `json:"success" example:"true"`
	Code    int                `json:"code" example:"0"`
	Message string             `json:"message" example:"Campaigns retrieved successfully"`
	TxnID   string             `json:"txn_id" example:"txn_123456"`
	Data    []domains.Campaign `json:"data,omitempty"`
}

// LinkResponse represents a response with link data
type LinkResponse struct {
	Success bool         `json:"success" example:"true"`
	Code    int          `json:"code" example:"0"`
	Message string       `json:"message" example:"Link created successfully"`
	TxnID   string       `json:"txn_id" example:"txn_123456"`
	Data    domains.Link `json:"data,omitempty"`
}

// LinksResponse represents a response with link array
type LinksResponse struct {
	Success bool           `json:"success" example:"true"`
	Code    int            `json:"code" example:"0"`
	Message string         `json:"message" example:"Links retrieved successfully"`
	TxnID   string         `json:"txn_id" example:"txn_123456"`
	Data    []domains.Link `json:"data,omitempty"`
}

// OfferResponse represents a response with offer data
type OfferResponse struct {
	Success bool          `json:"success" example:"true"`
	Code    int           `json:"code" example:"0"`
	Message string        `json:"message" example:"Offer retrieved successfully"`
	TxnID   string        `json:"txn_id" example:"txn_123456"`
	Data    domains.Offer `json:"data,omitempty"`
}

// BoolResponse represents a response with boolean data
type BoolResponse struct {
	Success bool   `json:"success" example:"true"`
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"Operation successful"`
	TxnID   string `json:"txn_id" example:"txn_123456"`
	Data    bool   `json:"data,omitempty" example:"true"`
}

// EmptyResponse represents a response with no data
type EmptyResponse struct {
	Success bool   `json:"success" example:"true"`
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"Operation successful"`
	TxnID   string `json:"txn_id" example:"txn_123456"`
}

// DashboardResponse represents a response with dashboard data
type DashboardResponse struct {
	Success bool                     `json:"success" example:"true"`
	Code    int                      `json:"code" example:"0"`
	Message string                   `json:"message" example:"Metrics retrieved successfully"`
	TxnID   string                   `json:"txn_id" example:"txn_123456"`
	Data    DashboardMetricsResponse `json:"data,omitempty"`
}
