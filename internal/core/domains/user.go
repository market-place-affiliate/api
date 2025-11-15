package domains

import "time"

type User struct {
	Id       int64  `json:"id" gorm:"primary_key;autoIncrement"`
	Email    string `json:"email" gorm:"column:email;type:text;not null;unique"`
	Password string `json:"password" gorm:"column:password;type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}
