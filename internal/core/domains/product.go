package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type Product struct {
	Id       uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuidv7()"`
	Title    string    `json:"title" gorm:"column:title;type:text;not null"`
	ImageUrl string    `json:"image_url" gorm:"column:image_url;type:text;not null"`

	SourceUrl string    `json:"source_url" gorm:"column:source_url;type:text;not null"`
	UserId    int64     `json:"user_id" gorm:"column:user_id;type:bigint REFERENCES users(id);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
}
