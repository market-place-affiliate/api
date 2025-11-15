package services

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/repositories/mocks"
	"github.com/market-place-affiliate/commonlib/lazada"
	"github.com/market-place-affiliate/commonlib/shopee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateLink_Lazada_Success(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	userId := int64(1)
	productId := uuid.Must(uuid.NewV4())
	campaignId := uuid.Must(uuid.NewV4())

	request := dto.CreateLinkRequest{
		ProductId:  productId,
		CampaignId: campaignId,
	}

	product := domains.Product{
		Id:        productId,
		UserId:    userId,
		SourceUrl: "https://lazada.co.th/product",
	}

	campaign := domains.Campaign{
		Id:          campaignId,
		UserId:      userId,
		UtmCampaign: "test_campaign",
	}

	offer := domains.Offer{
		ProductId:   productId,
		Marketplace: "lazada",
		StoreName:   "Test Store",
		Price:       100.0,
	}

	credential := domains.MarketplaceCredential{
		UserId:      userId,
		Marketplace: "lazada",
		AppKey:      "test_key",
		AppSecret:   "test_secret",
		UserToken:   "test_token",
	}

	lazadaResp := lazada.LazadaResponse[lazada.BatchPromoteLinkResponse]{
		Result: struct {
			Data    lazada.BatchPromoteLinkResponse `json:"data"`
			Success bool                            `json:"success"`
		}{
			Data: lazada.BatchPromoteLinkResponse{
				URLBatchGetLinkInfoList: []struct {
					RegularCommission    string `json:"regularCommission"`
					ProductID            string `json:"productId"`
					OriginalURL          string `json:"originalUrl"`
					RegularPromotionLink string `json:"regularPromotionLink"`
					ProductName          string `json:"productName"`
					Class                string `json:"class"`
				}{
					{
						RegularPromotionLink: "https://lazada.co.th/affiliate-link",
					},
				},
			},
		},
	}

	createdLink := domains.Link{
		Id:         uuid.Must(uuid.NewV4()),
		ProductId:  productId,
		CampaignId: campaignId,
		ShortCode:  "abc123",
		TargetURL:  "https://lazada.co.th/affiliate-link",
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)
	mockCampaignRepo.On("GetCampaignById", ctx, campaignId.String()).Return(campaign, nil)
	mockOfferRepo.On("GetOffersByProductId", ctx, productId.String()).Return(offer, nil)
	mockLinkRepo.On("GetLinkByShortCode", ctx, mock.AnythingOfType("string")).Return(domains.Link{}, assert.AnError)
	mockMarketCredRepo.On("GetByUserIdAndPlatform", ctx, userId, "lazada").Return(credential, nil)
	mockLazadaRepo.On("GetBatchPromoteLink", mock.AnythingOfType("lazada.LazadaCredentials"), "url", product.SourceUrl, mock.AnythingOfType("[6]string")).Return(lazadaResp, nil)
	mockLinkRepo.On("SaveLink", ctx, mock.MatchedBy(func(l domains.Link) bool {
		return l.ProductId == productId && l.CampaignId == campaignId
	})).Return(createdLink, nil)

	result, err := service.CreateLink(ctx, userId, request)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, createdLink, result.Data)
	mockProductRepo.AssertExpectations(t)
	mockCampaignRepo.AssertExpectations(t)
	mockOfferRepo.AssertExpectations(t)
	mockLinkRepo.AssertExpectations(t)
	mockMarketCredRepo.AssertExpectations(t)
	mockLazadaRepo.AssertExpectations(t)
}

func TestCreateLink_Shopee_Success(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	userId := int64(1)
	productId := uuid.Must(uuid.NewV4())
	campaignId := uuid.Must(uuid.NewV4())

	request := dto.CreateLinkRequest{
		ProductId:  productId,
		CampaignId: campaignId,
	}

	product := domains.Product{
		Id:        productId,
		UserId:    userId,
		SourceUrl: "https://shopee.co.th/product",
	}

	campaign := domains.Campaign{
		Id:          campaignId,
		UserId:      userId,
		UtmCampaign: "test_campaign",
	}

	offer := domains.Offer{
		ProductId:   productId,
		Marketplace: "shopee",
		StoreName:   "Test Store",
		Price:       100.0,
	}

	credential := domains.MarketplaceCredential{
		UserId:      userId,
		Marketplace: "shopee",
		AppId:       "test_app_id",
		AppSecret:   "test_secret",
	}

	shopeeResp := shopee.ShopeeGetShortLink{
		Data: struct {
			GenerateShortLink struct {
				ShortLink string `json:"shortLink"`
			} `json:"generateShortLink"`
		}{
			GenerateShortLink: struct {
				ShortLink string `json:"shortLink"`
			}{
				ShortLink: "https://shopee.co.th/short-link",
			},
		},
	}

	createdLink := domains.Link{
		Id:         uuid.Must(uuid.NewV4()),
		ProductId:  productId,
		CampaignId: campaignId,
		ShortCode:  "xyz789",
		TargetURL:  "https://shopee.co.th/short-link",
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)
	mockCampaignRepo.On("GetCampaignById", ctx, campaignId.String()).Return(campaign, nil)
	mockOfferRepo.On("GetOffersByProductId", ctx, productId.String()).Return(offer, nil)
	mockLinkRepo.On("GetLinkByShortCode", ctx, mock.AnythingOfType("string")).Return(domains.Link{}, assert.AnError)
	mockMarketCredRepo.On("GetByUserIdAndPlatform", ctx, userId, "shopee").Return(credential, nil)
	mockShopeeRepo.On("GetShortLink", mock.AnythingOfType("shopee.ShopeeCredentials"), product.SourceUrl, mock.AnythingOfType("[5]string")).Return(shopeeResp, nil)
	mockLinkRepo.On("SaveLink", ctx, mock.MatchedBy(func(l domains.Link) bool {
		return l.ProductId == productId && l.CampaignId == campaignId
	})).Return(createdLink, nil)

	result, err := service.CreateLink(ctx, userId, request)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, createdLink, result.Data)
	mockProductRepo.AssertExpectations(t)
	mockCampaignRepo.AssertExpectations(t)
	mockOfferRepo.AssertExpectations(t)
	mockLinkRepo.AssertExpectations(t)
	mockMarketCredRepo.AssertExpectations(t)
	mockShopeeRepo.AssertExpectations(t)
}

func TestCreateLink_ProductNotOwned(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	userId := int64(1)
	productId := uuid.Must(uuid.NewV4())
	campaignId := uuid.Must(uuid.NewV4())

	request := dto.CreateLinkRequest{
		ProductId:  productId,
		CampaignId: campaignId,
	}

	product := domains.Product{
		Id:     productId,
		UserId: int64(2), // Different user
	}

	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)

	result, err := service.CreateLink(ctx, userId, request)

	assert.NoError(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, 4003, result.Code)
	assert.Equal(t, "You are not allowed to create link for this product", result.Message)
	mockProductRepo.AssertExpectations(t)
}

func TestClickByShortCode_Success(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	shortCode := "abc123"
	linkId := uuid.Must(uuid.NewV4())

	link := domains.Link{
		Id:        linkId,
		ShortCode: shortCode,
		TargetURL: "https://example.com",
	}

	mockLinkRepo.On("GetLinkByShortCode", ctx, shortCode).Return(link, nil)
	mockClickRepo.On("SaveClick", ctx, mock.MatchedBy(func(c domains.Click) bool {
		return c.LinkId == linkId
	})).Return(nil)

	result, err := service.ClickByShortCode(ctx, shortCode)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, link, result.Data)
	mockLinkRepo.AssertExpectations(t)
	mockClickRepo.AssertExpectations(t)
}

func TestGetLinkByCampaign_Success(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	campaignId := uuid.Must(uuid.NewV4())

	expectedLinks := []domains.Link{
		{
			Id:         uuid.Must(uuid.NewV4()),
			CampaignId: campaignId,
			ShortCode:  "abc123",
		},
		{
			Id:         uuid.Must(uuid.NewV4()),
			CampaignId: campaignId,
			ShortCode:  "xyz789",
		},
	}

	mockLinkRepo.On("GetLinksByCampaignId", ctx, campaignId.String()).Return(expectedLinks, nil)

	result, err := service.GetLinkByCampaign(ctx, campaignId.String())

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, expectedLinks, result.Data)
	mockLinkRepo.AssertExpectations(t)
}

func TestDeleteLinkById_Success(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	userId := int64(1)
	linkId := uuid.Must(uuid.NewV4())
	productId := uuid.Must(uuid.NewV4())

	link := domains.Link{
		Id:        linkId,
		ProductId: productId,
	}

	product := domains.Product{
		Id:     productId,
		UserId: userId,
	}

	mockLinkRepo.On("GetLinkById", ctx, linkId.String()).Return(link, nil)
	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)
	mockClickRepo.On("DeleteClicksByLinkId", ctx, linkId.String()).Return(nil)
	mockLinkRepo.On("DeleteLink", ctx, linkId.String()).Return(nil)

	result, err := service.DeleteLinkById(ctx, userId, linkId.String())

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	mockLinkRepo.AssertExpectations(t)
	mockProductRepo.AssertExpectations(t)
	mockClickRepo.AssertExpectations(t)
}

func TestDeleteLinkById_Forbidden(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	userId := int64(1)
	linkId := uuid.Must(uuid.NewV4())
	productId := uuid.Must(uuid.NewV4())

	link := domains.Link{
		Id:        linkId,
		ProductId: productId,
	}

	product := domains.Product{
		Id:     productId,
		UserId: int64(2), // Different user
	}

	mockLinkRepo.On("GetLinkById", ctx, linkId.String()).Return(link, nil)
	mockProductRepo.On("GetProductById", ctx, productId.String()).Return(product, nil)

	result, err := service.DeleteLinkById(ctx, userId, linkId.String())

	assert.NoError(t, err)
	assert.False(t, result.Success)
	assert.Equal(t, 5003, result.Code)
	assert.Equal(t, "You do not have permission to delete this link", result.Message)
	mockLinkRepo.AssertExpectations(t)
	mockProductRepo.AssertExpectations(t)
}

func TestGetLinkById_Success(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	linkId := uuid.Must(uuid.NewV4())

	expectedLink := domains.Link{
		Id:        linkId,
		ShortCode: "abc123",
		TargetURL: "https://example.com",
	}

	mockLinkRepo.On("GetLinkById", ctx, linkId.String()).Return(expectedLink, nil)

	result, err := service.GetLinkById(ctx, linkId.String())

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, expectedLink, result.Data)
	mockLinkRepo.AssertExpectations(t)
}

func TestGetLinkByShortCode_Success(t *testing.T) {
	mockLinkRepo := new(mocks.MockLinkRepository)
	mockClickRepo := new(mocks.MockClickRepository)
	mockProductRepo := new(mocks.MockProductRepository)
	mockCampaignRepo := new(mocks.MockCampaignRepository)
	mockOfferRepo := new(mocks.MockOfferRepository)
	mockLazadaRepo := new(mocks.MockLazadaRepository)
	mockShopeeRepo := new(mocks.MockShopeeRepository)
	mockMarketCredRepo := new(mocks.MockMarketplaceRepository)

	service := NewLinkService(mockLinkRepo, mockClickRepo, mockProductRepo, mockCampaignRepo, mockOfferRepo, mockLazadaRepo, mockShopeeRepo, mockMarketCredRepo)

	ctx := context.Background()
	shortCode := "abc123"

	expectedLink := domains.Link{
		Id:        uuid.Must(uuid.NewV4()),
		ShortCode: shortCode,
		TargetURL: "https://example.com",
	}

	mockLinkRepo.On("GetLinkByShortCode", ctx, shortCode).Return(expectedLink, nil)

	result, err := service.GetLinkByShortCode(ctx, shortCode)

	assert.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 0, result.Code)
	assert.Equal(t, expectedLink, result.Data)
	mockLinkRepo.AssertExpectations(t)
}
