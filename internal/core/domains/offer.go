package domains

import (
	"time"

	"github.com/gofrs/uuid"
)

type Offer struct {
	Id            uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuidv7()"`
	ProductId     uuid.UUID `gorm:"column:product_id;type:uuid REFERENCES products(id)"`
	Marketplace   string    `json:"marketplace" gorm:"column:marketplace;type:text;not null"`
	StoreName     string    `json:"store_name" gorm:"column:store_name;type:text;not null"`
	Price         float64   `json:"price" gorm:"column:price;type:decimal(10,2);not null"`
	LastCheckedAt time.Time `json:"last_checked_at" gorm:"column:last_checked_at;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
}
