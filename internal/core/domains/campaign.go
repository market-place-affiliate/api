package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type Campaign struct {
	Id          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuidv7()"`
	Name        string    `json:"name" gorm:"column:name;type:text;not null"`
	UtmCampaign string    `json:"utm_campaign" gorm:"column:utm_campaign;type:text;not null"`
	StartAt     time.Time `json:"start_at" gorm:"column:start_at;not null"`
	EndAt       time.Time `json:"end_at" gorm:"column:end_at;not null"`

	UserId    int64     `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
}
