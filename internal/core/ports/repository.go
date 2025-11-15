package ports

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domains.User) (domains.User, error)
	GetUserByID(ctx context.Context, userId int64) (domains.User, error)
	GetUserByEmail(ctx context.Context, email string) (domains.User, error)
}

type ProductRepository interface {
	SaveProduct(ctx context.Context, product domains.Product) (domains.Product, error)
	DeleteProduct(ctx context.Context, productId string) error
	GetProductById(ctx context.Context, productId string) (domains.Product, error)
	GetAllProducts(ctx context.Context, userId int64) ([]domains.Product, error)
	DeleteProductById(ctx context.Context, productId string) error
}

type OfferRepository interface {
	SaveOffer(ctx context.Context, offer domains.Offer) error
	DeleteOffer(ctx context.Context, offerId string) error
	GetOffersByProductId(ctx context.Context, productId string) (domains.Offer, error)
	GetOfferById(ctx context.Context, offerId string) (domains.Offer, error)
	DeleteOfferByProductId(ctx context.Context, productId string) error
}

type LinkRepository interface {
	SaveLink(ctx context.Context, link domains.Link) (domains.Link, error)
	DeleteLink(ctx context.Context, linkId string) error
	GetLinksByProductId(ctx context.Context, productId string) ([]domains.Link, error)
	GetLinkById(ctx context.Context, linkId string) (domains.Link, error)
	GetLinkByShortCode(ctx context.Context, shortCode string) (domains.Link, error)
	GetLinksByCampaignId(ctx context.Context, campaignId string) ([]domains.Link, error)
	DeleteLinkByProductId(ctx context.Context, productId string) error
	DeleteLinkByCampaignId(ctx context.Context, campaignId string) error
}

type ClickRepository interface {
	SaveClick(ctx context.Context, click domains.Click) error
	CountClicksByDateRange(ctx context.Context, userId int64, startDate, endDate time.Time) ([]dto.MetrictItem, error)
	CountTopProductClickByDateRange(ctx context.Context, userId int64, startDate, endDate time.Time) (uuid.UUID, int64, error)
	DeleteClicksByLinkId(ctx context.Context, linkId string) error
}

type CampaignRepository interface {
	SaveCampaign(ctx context.Context, campaign domains.Campaign) (domains.Campaign, error)
	DeleteCampaign(ctx context.Context, campaignId string) error
	GetCampaignById(ctx context.Context, campaignId string) (domains.Campaign, error)
	GetCampaignByQuery(ctx context.Context, userId int64, query dto.GetCampaignByQueryRequest) ([]domains.Campaign, error)
}

type MarketplaceRepository interface {
	Save(ctx context.Context, marketplace domains.MarketplaceCredential) (domains.MarketplaceCredential, error)
	GetByUserIdAndPlatform(ctx context.Context, userId int64, platform string) (domains.MarketplaceCredential, error)
	DeleteByUserIdAndPlatform(ctx context.Context, userId int64, platform string) error
}
