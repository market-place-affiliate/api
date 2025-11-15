package services

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetOffer_Success(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewProductService(mockProductRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	productId := uuid.Must(uuid.NewV4())

	product := domains.Product{
		Id:     productId,
		UserId: userId,
		Title:  "Test Product",
	}

	offer := domains.Offer{
		ProductId:   productId,
		Marketplace: "lazada",
		StoreName:   "Test Store",
		Price:       100.0,
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)
	mockOfferRepo.On("GetOffersByProductId", ctx, productId.String()).Return(offer, nil)

	result, err := service.GetOffer(ctx, userId, productId.String())

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, offer, result.Data)
	mockProductRepo.AssertExpectations(t)
	mockOfferRepo.AssertExpectations(t)
}

func TestGetOffer_Forbidden(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewProductService(mockProductRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	productId := uuid.Must(uuid.NewV4())

	product := domains.Product{
		Id:     productId,
		UserId: int64(2), // Different user
		Title:  "Test Product",
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)

	result, err := service.GetOffer(ctx, userId, productId.String())

	assert.NoError(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, 2003, result.Code)
	assert.Equal(t, "You do not have access to this product offers", result.Message)
	mockProductRepo.AssertExpectations(t)
}

func TestGetProductsByUserId_Success(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewProductService(mockProductRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)

	products := []domains.Product{
		{
			Id:     uuid.Must(uuid.NewV4()),
			UserId: userId,
			Title:  "Product 1",
		},
		{
			Id:     uuid.Must(uuid.NewV4()),
			UserId: userId,
			Title:  "Product 2",
		},
	}

	mockProductRepo.On("GetAllProducts", ctx, userId).Return(products, nil)

	result, err := service.GetProductsByUserId(ctx, userId)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, 2, len(result.Data))
	mockProductRepo.AssertExpectations(t)
}

func TestDeleteProductById_Success(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewProductService(mockProductRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	productId := uuid.Must(uuid.NewV4())
	linkId := uuid.Must(uuid.NewV4())

	product := domains.Product{
		Id:     productId,
		UserId: userId,
		Title:  "Test Product",
	}

	links := []domains.Link{
		{
			Id:        linkId,
			ProductId: productId,
		},
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)
	mockLinkRepo.On("GetLinksByProductId", ctx, productId.String()).Return(links, nil)
	mockClickRepo.On("DeleteClicksByLinkId", ctx, linkId.String()).Return(nil)
	mockLinkRepo.On("DeleteLinkByProductId", ctx, productId.String()).Return(nil)
	mockOfferRepo.On("DeleteOfferByProductId", ctx, productId.String()).Return(nil)
	mockProductRepo.On("DeleteProductById", ctx, productId.String()).Return(nil)

	result, err := service.DeleteProductById(ctx, userId, productId.String())

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	mockProductRepo.AssertExpectations(t)
	mockLinkRepo.AssertExpectations(t)
	mockClickRepo.AssertExpectations(t)
	mockOfferRepo.AssertExpectations(t)
}

func TestDeleteProductById_Forbidden(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewProductService(mockProductRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	userId := int64(1)
	productId := uuid.Must(uuid.NewV4())

	product := domains.Product{
		Id:     productId,
		UserId: int64(2), // Different user
		Title:  "Test Product",
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)

	result, err := service.DeleteProductById(ctx, userId, productId.String())

	assert.NoError(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, 2003, result.Code)
	assert.Equal(t, "You do not have access to delete this product", result.Message)
	mockProductRepo.AssertExpectations(t)
}

func TestGetProductById_Success(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)

	service := NewProductService(mockProductRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo, mockLinkRepo, mockClickRepo)

	ctx := context.Background()
	productId := uuid.Must(uuid.NewV4())

	product := domains.Product{
		Id:     productId,
		UserId: int64(1),
		Title:  "Test Product",
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)

	result, err := service.GetProductById(ctx, productId.String())

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, product, result.Data)
	mockProductRepo.AssertExpectations(t)
}
