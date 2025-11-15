package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type Link struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuidv7()"`
	ProductId uuid.UUID `gorm:"column:product_id;type:uuid REFERENCES products(id)"`
	CampaignId uuid.UUID `gorm:"column:campaign_id;type:uuid REFERENCES campaigns(id)"`
	ShortCode string    `json:"short_code" gorm:"column:short_code;type:text;not null;unique"`
	TargetURL string    `json:"target_url" gorm:"column:target_url;type:text;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
}
