package services

import (
	"context"
	"net/http"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
	"github.com/market-place-affiliate/api/pkg/uniqe"
	"github.com/market-place-affiliate/commonlib/lazada"
	"github.com/market-place-affiliate/commonlib/shopee"
)

type linkService struct {
	linkRepo       ports.LinkRepository
	clickRepo      ports.ClickRepository
	productRepo    ports.ProductRepository
	campaignRepo   ports.CampaignRepository
	offerRepo      ports.OfferRepository
	lazadaRepo     lazada.LazadaRepository
	shopeeRepo     shopee.ShopeeRepository
	marketCredRepo ports.MarketplaceRepository
}

func NewLinkService(linkRepo ports.LinkRepository, clickRepo ports.ClickRepository, productRepo ports.ProductRepository, campaignRepo ports.CampaignRepository, offerRepo ports.OfferRepository, lazadaRepo lazada.LazadaRepository, shopeeRepo shopee.ShopeeRepository, marketCredRepo ports.MarketplaceRepository) ports.LinkService {
	return &linkService{linkRepo: linkRepo, clickRepo: clickRepo, productRepo: productRepo, campaignRepo: campaignRepo, offerRepo: offerRepo, lazadaRepo: lazadaRepo, shopeeRepo: shopeeRepo, marketCredRepo: marketCredRepo}
}

func (s *linkService) CreateLink(ctx context.Context, userId int64, link dto.CreateLinkRequest) (dto.Response[domains.Link], error) {

	product, err := s.productRepo.GetProductById(ctx, link.ProductId.String())
	if err != nil {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4002,
			Message:  "Product not found",
		}, err
	}

	if product.UserId != userId {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusForbidden,
			Success:  false,
			Code:     4003,
			Message:  "You are not allowed to create link for this product",
		}, nil
	}

	campaign, err := s.campaignRepo.GetCampaignById(ctx, link.CampaignId.String())
	if err != nil {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4004,
			Message:  "Campaign not found",
		}, err
	}
	if campaign.UserId != userId {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusForbidden,
			Success:  false,
			Code:     4005,
			Message:  "You are not allowed to create link for this campaign",
		}, nil
	}

	offer, err := s.offerRepo.GetOffersByProductId(ctx, product.Id.String())
	if err != nil {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4006,
			Message:  "Offer not found for this product",
		}, err
	}

	newLink := domains.Link{
		ProductId: link.ProductId,
		ShortCode: uniqe.UUID(),
	}

	switch offer.Marketplace {
	case "lazada":
		cred, err := s.marketCredRepo.GetByUserIdAndPlatform(ctx, userId, "lazada")
		if err != nil {
			return dto.Response[domains.Link]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     4009,
				Message:  "Lazada marketplace credentials not found",
			}, err
		}
		lazadaResp, err := s.lazadaRepo.GetBatchPromoteLink(lazada.LazadaCredentials{
			AppKey:     cred.AppKey,
			AppSecret:  cred.AppSecret,
			SignMethod: "sha256",
			UserToken:  cred.UserToken,
		}, "url", product.SourceUrl, [6]string{campaign.UtmCampaign})
		if err != nil || len(lazadaResp.Result.Data.URLBatchGetLinkInfoList) == 0 {
			return dto.Response[domains.Link]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     4007,
				Message:  "Failed to generate lazada affiliate link",
			}, err
		}
		newLink.TargetURL = lazadaResp.Result.Data.URLBatchGetLinkInfoList[0].RegularPromotionLink
	case "shopee":

		cred, err := s.marketCredRepo.GetByUserIdAndPlatform(ctx, userId, "shopee")
		if err != nil {
			return dto.Response[domains.Link]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     4009,
				Message:  "Shopee marketplace credentials not found",
			}, err
		}

		shopeeResp, err := s.shopeeRepo.GetShortLink(shopee.ShopeeCredentials{
			AppId:     cred.AppId,
			AppSecret: cred.AppSecret,
		}, product.SourceUrl, [5]string{campaign.UtmCampaign})
		if err != nil || shopeeResp.Data.GenerateShortLink.ShortLink == "" {
			return dto.Response[domains.Link]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     4008,
				Message:  "Failed to generate shopee affiliate link",
			}, err
		}
		newLink.TargetURL = shopeeResp.Data.GenerateShortLink.ShortLink
	}

	createdLink, err := s.linkRepo.SaveLink(ctx, newLink)
	if err != nil {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4001,
			Message:  "Failed to create link",
		}, err
	}

	return dto.Response[domains.Link]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Data:     createdLink,
		Message:  "Link created successfully",
	}, nil
}

func (s *linkService) ClickByShortCode(ctx context.Context, shortCode string) (dto.Response[domains.Link], error) {
	link, err := s.linkRepo.GetLinkByShortCode(ctx, shortCode)
	if err != nil {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4001,
			Message:  "Failed to get link by short code",
		}, err
	}
	err = s.clickRepo.SaveClick(ctx, domains.Click{
		LinkId: link.Id,
	})
	if err != nil {
		return dto.Response[domains.Link]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4002,
			Message:  "Failed to record click",
		}, err
	}
	return dto.Response[domains.Link]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Data:     link,
	}, nil
}

func (s *linkService) GetLinkByCampaign(ctx context.Context, campaignId string) (dto.Response[[]domains.Link], error) {
	links, err := s.linkRepo.GetLinksByCampaignId(ctx, campaignId)
	if err != nil {
		return dto.Response[[]domains.Link]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     4001,
			Message:  "Failed to get links by campaign id",
		}, err
	}
	return dto.Response[[]domains.Link]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Data:     links,
	}, nil
}