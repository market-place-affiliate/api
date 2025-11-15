package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type Click struct {
	Id uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuidv7()"`
	LinkId     uuid.UUID `gorm:"column:link_id;type:uuid REFERENCES links(id)"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
}
