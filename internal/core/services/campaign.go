package services

import (
	"context"
	"net/http"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
)

type campaignService struct {
	campaignRepo ports.CampaignRepository
	linkRepo     ports.LinkRepository
	clickRepo    ports.ClickRepository
}

func NewCampaignService(campaignRepo ports.CampaignRepository, linkRepo ports.LinkRepository, clickRepo ports.ClickRepository) ports.CampaignService {
	return &campaignService{campaignRepo: campaignRepo, linkRepo: linkRepo, clickRepo: clickRepo}
}

func (c *campaignService) CreateCampaign(ctx context.Context, userId int64, campaign dto.CreateCampaignRequest) (dto.Response[domains.Campaign], error) {
	newCampaign, err := c.campaignRepo.SaveCampaign(ctx, domains.Campaign{
		Name:        campaign.Name,
		UtmCampaign: campaign.UtmCampaign,
		StartAt:     campaign.StartAt,
		EndAt:       campaign.EndAt,
		UserId:      userId,
	})
	if err != nil {
		return dto.Response[domains.Campaign]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     3001,
		}, err
	}
	return dto.Response[domains.Campaign]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Data:     newCampaign,
	}, nil
}
func (c *campaignService) GetCampaignByQuery(ctx context.Context, userId int64, query dto.GetCampaignByQueryRequest) (dto.Response[[]domains.Campaign], error) {
	campaigns, err := c.campaignRepo.GetCampaignByQuery(ctx, userId, query)
	if err != nil {
		return dto.Response[[]domains.Campaign]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     3002,
		}, err
	}
	return dto.Response[[]domains.Campaign]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Data:     campaigns,
	}, nil
}
func (c *campaignService) DeleteCampaignById(ctx context.Context, userId int64, campaignId string) (dto.Response[any], error) {

	campaign, err := c.campaignRepo.GetCampaignById(ctx, campaignId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     3004,
			Message:  "Failed to fetch campaign",
		}, err
	}
	if campaign.UserId != userId {
		return dto.Response[any]{
			HttpCode: http.StatusForbidden,
			Success:  false,
			Code:     3005,
			Message:  "You do not have permission to delete this campaign",
		}, nil
	}

	links, err := c.linkRepo.GetLinksByCampaignId(ctx, campaignId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     3006,
			Message:  "Failed to fetch links for the campaign",
		}, err
	}

	for _, link := range links {
		err = c.clickRepo.DeleteClicksByLinkId(ctx, link.Id.String())
		if err != nil {
			return dto.Response[any]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     3007,
				Message:  "Failed to delete clicks for link " + link.Id.String(),
			}, err
		}
		err = c.linkRepo.DeleteLink(ctx, link.Id.String())
		if err != nil {
			return dto.Response[any]{
				HttpCode: http.StatusInternalServerError,
				Success:  false,
				Code:     3008,
				Message:  "Failed to delete link " + link.Id.String(),
			}, err
		}
	}

	err = c.campaignRepo.DeleteCampaign(ctx, campaignId)
	if err != nil {
		return dto.Response[any]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     3003,
		}, err
	}
	return dto.Response[any]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
	}, nil
}
func (c *campaignService) GetPublicCampaigns(ctx context.Context, query dto.GetCampaignByQueryRequest) (dto.Response[[]domains.Campaign], error) {
	campaigns, err := c.campaignRepo.GetAvailableCampaign(ctx)
	if err != nil {
		return dto.Response[[]domains.Campaign]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     3002,
		}, err
	}
	return dto.Response[[]domains.Campaign]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Data:     campaigns,
	}, nil
}
