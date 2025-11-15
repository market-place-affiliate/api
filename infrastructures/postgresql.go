package infrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(host string, port int, username string, password string, dbname string) *gorm.DB {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d`, host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetMaxIdleConns(10)
	if err != nil {
		panic(err)
	}
	return db
}
