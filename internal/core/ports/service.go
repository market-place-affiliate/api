package ports

import (
	"context"
	"time"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
)

type UserService interface {
	Register(ctx context.Context, password, email string) (dto.Response[string], error)
	Login(ctx context.Context, password, email string) (dto.Response[string], error)
	GetMe(ctx context.Context, userId int64) (dto.Response[domains.User], error)
	VerifyAndGetUserId(token string) (int64, error)
	SaveMarketplaceCredential(ctx context.Context, userId int64, cred dto.MarketplaceCredentialRequest) (dto.Response[string], error)
	CheckMarketplaceCredential(ctx context.Context, userId int64, platform string) (dto.Response[bool], error)
	DeleteMarketplaceCredential(ctx context.Context, userId int64, platform string) (dto.Response[string], error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, userId int64, product dto.CreateProductRequest) (dto.Response[[]domains.Product], error)
	GetOffer(ctx context.Context, userId int64, productId string) (dto.Response[domains.Offer], error)
	GetProductsByUserId(ctx context.Context, userId int64) (dto.Response[[]domains.Product], error)
	DeleteProductById(ctx context.Context, userId int64, productId string) (dto.Response[any], error)
}

type CampaignService interface {
	CreateCampaign(ctx context.Context, userId int64, campaign dto.CreateCampaignRequest) (dto.Response[domains.Campaign], error)
	GetCampaignByQuery(ctx context.Context, userId int64, query dto.GetCampaignByQueryRequest) (dto.Response[[]domains.Campaign], error)
}

type LinkService interface {
	CreateLink(ctx context.Context, userId int64, link dto.CreateLinkRequest) (dto.Response[domains.Link], error)
	GetLinkByCampaign(ctx context.Context, campaignId string) (dto.Response[[]domains.Link], error)
	ClickByShortCode(ctx context.Context, shortCode string) (dto.Response[domains.Link], error)
}

type DashboardService interface {
	GetDashboardMetrics(ctx context.Context, userId int64, startDate, endDate time.Time) (dto.Response[dto.DashboardMetricsResponse], error)
}
