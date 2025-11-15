package db

import (
	"context"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ports.ProductRepository {
	return &productRepository{DB: db}
}

func (r *productRepository) SaveProduct(ctx context.Context, product domains.Product) (domains.Product, error) {
	err := r.DB.Save(&product).Error
	if err != nil {
		return domains.Product{}, err
	}
	return product, nil
}
func (r *productRepository) DeleteProduct(ctx context.Context, productId string) error {
	err := r.DB.Delete(&domains.Product{}, "id = ?", productId).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *productRepository) GetProductById(ctx context.Context, productId string) (domains.Product, error) {
	var product domains.Product
	err := r.DB.First(&product, "id = ?", productId).Error
	if err != nil {
		return domains.Product{}, err
	}
	return product, nil
}
func (r *productRepository) GetAllProducts(ctx context.Context, userId int64) ([]domains.Product, error) {
	var products []domains.Product
	err := r.DB.Where("user_id = ?", userId).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) DeleteProductById(ctx context.Context, productId string) error {
	err := r.DB.Delete(&domains.Product{}, "id = ?", productId).Error
	if err != nil {
		return err
	}
	return nil
}