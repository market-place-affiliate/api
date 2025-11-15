package db

import (
	"github.com/market-place-affiliate/api/internal/core/domains"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(&domains.User{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&domains.Product{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&domains.Campaign{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&domains.Offer{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&domains.Link{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&domains.MarketplaceCredential{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&domains.Click{})
	if err != nil {
		return err
	}
	return nil
}
