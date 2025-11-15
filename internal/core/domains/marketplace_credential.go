package domains

import "time"

type MarketplaceCredential struct {
	Id          int64  `json:"id" gorm:"primary_key;autoIncrement"`
	UserId      int64  `json:"user_id" gorm:"column:user_id;type:bigint REFERENCES users(id);not null;uniqueIndex:idx_user_marketplace"`
	Marketplace string `json:"marketplace" gorm:"column:marketplace;type:text;not null;uniqueIndex:idx_user_marketplace"`

	AppId     string `json:"app_id" gorm:"column:app_id;type:text;not null"`
	AppKey    string `json:"app_key" gorm:"column:app_key;type:text;not null"`
	AppSecret string `json:"app_secret" gorm:"column:app_secret;type:text;not null"`
	UserToken string `json:"user_token" gorm:"column:user_token;type:text;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}
