package services

import (
	"context"
	"net/http"
	"strconv"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"github.com/market-place-affiliate/api/pkg/customtime"
	"github.com/market-place-affiliate/commonlib/lazada"
	"github.com/market-place-affiliate/commonlib/shopee"
)

type productService struct {
	productRepo    ports.ProductRepository
	offerRepo      ports.OfferRepository
	lazadaRepo     lazada.LazadaRepository
	shopeeRepo     shopee.ShopeeRepository
	marketCredRepo ports.MarketplaceRepository
	linkRepo       ports.LinkRepository
	clickRepo      ports.ClickRepository
}

func NewProductService(productRepo ports.ProductRepository, offerRepo ports.OfferRepository, lazadaRepo lazada.LazadaRepository, shopeeRepo shopee.ShopeeRepository, marketCredRepo ports.MarketplaceRepository, linkRepo ports.LinkRepository, clickRepo ports.ClickRepository) ports.ProductService {
	return &productService{
		productRepo:    productRepo,
		offerRepo:      offerRepo,
		lazadaRepo:     lazadaRepo,
		shopeeRepo:     shopeeRepo,
		marketCredRepo: marketCredRepo,
		linkRepo:       linkRepo,
		clickRepo:      clickRepo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, userId int64, product dto.CreateProductRequest) (dto.Response[[]domains.Product], error) {
	resPProducts := []domains.Product{}
	switch product.Marketplace {
	case "lazada":

		cred, err := s.marketCredRepo.GetByUserIdAndPlatform(ctx, userId, "lazada")
		if err != nil {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     2001,
				Message:  "Failed to fetch marketplace credential",
			}, err
		}

		lazadaResp, err := s.lazadaRepo.GetBatchPromoteLink(lazada.LazadaCredentials{
			AppKey:     cred.AppKey,
			AppSecret:  cred.AppSecret,
			SignMethod: "sha256",
			UserToken:  cred.UserToken,
		}, "url", product.SourceUrl, [6]string{})
		if err != nil {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     2001,
				Message:  "Failed to fetch product from lazada",
			}, err
		}
		if len(lazadaResp.Result.Data.URLBatchGetLinkInfoList) == 0 {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     2001,
				Message:  "Failed to fetch product from lazada",
			}, nil
		}

		for _, promote := range lazadaResp.Result.Data.URLBatchGetLinkInfoList {
			lazadaProductFeed, err := s.lazadaRepo.GetProductFeed(lazada.LazadaCredentials{
				AppKey:     cred.AppKey,
				AppSecret:  cred.AppSecret,
				SignMethod: "sha256",
				UserToken:  cred.UserToken,
			}, promote.ProductID, 1, 1)
			if err != nil {
				return dto.Response[[]domains.Product]{
					HttpCode: http.StatusInternalServerError,
					Success:  false,
					Code:     2001,
					Message:  "Failed to fetch product from lazada",
				}, err
			}
			if len(lazadaProductFeed.Result.Data) == 0 {
				return dto.Response[[]domains.Product]{
					HttpCode: http.StatusInternalServerError,
					Success:  false,
					Code:     2001,
					Message:  "This product is not available for affiliation",
				}, nil
			}
			for _, feed := range lazadaProductFeed.Result.Data {
				prod := domains.Product{
					Title:     feed.ProductName,
					ImageUrl:  feed.Pictures[0],
					UserId:    userId,
					SourceUrl: product.SourceUrl,
				}
				storeName := feed.BrandName
				if storeName == "" {
					storeName = "Lazada Official Store"
				}
				offer := domains.Offer{
					Marketplace:   product.Marketplace,
					StoreName:     storeName,
					Price:         feed.DiscountPrice,
					LastCheckedAt: customtime.Now(),
				}

				createdProd, err := s.productRepo.SaveProduct(ctx, prod)
				if err != nil {
					return dto.Response[[]domains.Product]{
						HttpCode: http.StatusInternalServerError,
						Success:  false,
						Code:     2001,
						Message:  "Failed to save product",
					}, err
				}
				resPProducts = append(resPProducts, createdProd)
				offer.ProductId = createdProd.Id

				err = s.offerRepo.SaveOffer(ctx, offer)
				if err != nil {
					return dto.Response[[]domains.Product]{
						HttpCode: http.StatusInternalServerError,
						Success:  false,
						Code:     2001,
						Message:  "Failed to save offer",
					}, err
				}
			}
		}
	case "shopee":
		shopId, ItemId, err := shopee.ExtractShopIdAndItemIdFromLink(product.SourceUrl)
		if err != nil {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     http.StatusInternalServerError,
				Message:  "Failed to fetch product from shopee",
			}, err
		}
		cred, err := s.marketCredRepo.GetByUserIdAndPlatform(ctx, userId, "shopee")
		if err != nil {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     2001,
				Message:  "Failed to fetch marketplace credential",
			}, err
		}
		shoppeeResp, err := s.shopeeRepo.GetProductOfferListV2(shopee.ShopeeCredentials{
			AppId:     cred.AppId,
			AppSecret: cred.AppSecret,
		}, shopId, ItemId)
		if err != nil {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     http.StatusInternalServerError,
				Message:  "Failed to fetch product from shopee",
			}, err
		}

		if len(shoppeeResp.Data.ProductOfferV2.Nodes) == 0 {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     http.StatusInternalServerError,
				Message:  "Failed to fetch product from shopee",
			}, nil
		}

		prod := domains.Product{
			Title:     shoppeeResp.Data.ProductOfferV2.Nodes[0].ProductName,
			ImageUrl:  shoppeeResp.Data.ProductOfferV2.Nodes[0].ImageURL,
			UserId:    userId,
			SourceUrl: product.SourceUrl,
		}

		createdProd, err := s.productRepo.SaveProduct(ctx, prod)
		if err != nil {
			return dto.Response[[]domains.Product]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     2001,
				Message:  "Failed to save product",
			}, err
		}
		resPProducts = append(resPProducts, createdProd)

		for _, offer := range shoppeeResp.Data.ProductOfferV2.Nodes {
			price, _ := strconv.ParseFloat(offer.Price, 64)
			offer := domains.Offer{
				ProductId:     createdProd.Id,
				Marketplace:   product.Marketplace,
				StoreName:     offer.ShopName,
				Price:         price,
				LastCheckedAt: customtime.Now(),
			}
			err = s.offerRepo.SaveOffer(ctx, offer)
			if err != nil {
				return dto.Response[[]domains.Product]{
					HttpCode: http.StatusInternalServerError,
					Success:  false,
					Code:     2001,
					Message:  "Failed to save offer",
				}, err
			}
		}

	}

	return dto.Response[[]domains.Product]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Product created successfully",
		Data:     resPProducts,
	}, nil
}

func (s *productService) GetOffer(ctx context.Context, userId int64, productId string) (dto.Response[domains.Offer], error) {
	product, err := s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		return dto.Response[domains.Offer]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2002,
			Message:  "Failed to fetch product",
		}, err
	}
	if product.UserId != userId {
		return dto.Response[domains.Offer]{
			HttpCode: http.StatusForbidden,
			Success:  false,
			Code:     2003,
			Message:  "You do not have access to this product offers",
		}, nil
	}
	offer, err := s.offerRepo.GetOffersByProductId(ctx, productId)
	if err != nil {
		return dto.Response[domains.Offer]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2004,
			Message:  "Failed to fetch offers",
		}, err
	}
	return dto.Response[domains.Offer]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Offers fetched successfully",
		Data:     offer,
	}, nil
}

func (s *productService) GetProductsByUserId(ctx context.Context, userId int64) (dto.Response[[]domains.Product], error) {
	products, err := s.productRepo.GetAllProducts(ctx, userId)
	if err != nil {
		return dto.Response[[]domains.Product]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2005,
			Message:  "Failed to fetch products",
		}, err
	}
	userProducts := []domains.Product{}
	for _, product := range products {
		if product.UserId == userId {
			userProducts = append(userProducts, product)
		}
	}
	return dto.Response[[]domains.Product]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Products fetched successfully",
		Data:     userProducts,
	}, nil
}

func (s *productService) DeleteProductById(ctx context.Context, userId int64, productId string) (dto.Response[any], error) {
	product, err := s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2002,
			Message:  "Failed to fetch product",
		}, err
	}
	if product.UserId != userId {
		return dto.Response[any]{
			HttpCode: http.StatusForbidden,
			Success:  false,
			Code:     2003,
			Message:  "You do not have access to delete this product",
		}, nil
	}

	links, err := s.linkRepo.GetLinksByProductId(ctx, productId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2006,
			Message:  "Failed to fetch links associated with the product",
		}, err
	}

	for _, link := range links {
		err = s.clickRepo.DeleteClicksByLinkId(ctx, link.Id.String())
		if err != nil {
			return dto.Response[any]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     2006,
				Message:  "Failed to delete clicks associated with the product links",
			}, err
		}
	}

	err = s.linkRepo.DeleteLinkByProductId(ctx, productId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2006,
			Message:  "Failed to delete links associated with the product",
		}, err
	}
	err = s.offerRepo.DeleteOfferByProductId(ctx, productId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2006,
			Message:  "Failed to delete offers associated with the product",
		}, err
	}
	err = s.productRepo.DeleteProductById(ctx, productId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2006,
			Message:  "Failed to delete product",
		}, err
	}
	return dto.Response[any]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Product deleted successfully",
	}, nil
}

func (s *productService) GetProductById(ctx context.Context, productId string) (dto.Response[domains.Product], error) {
	product, err := s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		return dto.Response[domains.Product]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     2002,
			Message:  "Failed to fetch product",
		}, err
	}
	return dto.Response[domains.Product]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Product fetched successfully",
		Data:     product,
	}, nil
}