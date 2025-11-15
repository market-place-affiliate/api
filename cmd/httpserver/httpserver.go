package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/market-place-affiliate/api/internal/handlers"
)

func NewHttpServer(
	userHandler *handlers.UserHandler,
	productHandler *handlers.ProductHandler,
	campaignHandler *handlers.CampaignHandler,
	linkHandler *handlers.LinkHandler,
	dashboardHandler *handlers.DashboardHandler,
) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	g.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		// c.Writer.Header().Set("Vary", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	g.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	g.GET("/go/:short_code", linkHandler.RedirectLink)

	api := g.Group("/api")
	apiV1 := api.Group("/v1")
	v1UserGroup := apiV1.Group("user")
	v1UserGroup.POST("/register", userHandler.Register)
	v1UserGroup.POST("/login", userHandler.Login)
	v1UserGroup.POST("/logout", userHandler.Logout)
	v1UserGroup.GET("/me", userHandler.VerifyAndGetUserId, userHandler.GetMe)
	v1UserGroup.POST("/market-credential", userHandler.VerifyAndGetUserId, userHandler.SaveMarketplaceCredential)
	v1UserGroup.GET("/market-credential/:platform", userHandler.VerifyAndGetUserId, userHandler.CheckMarketplaceCredential)
	v1UserGroup.DELETE("/market-credential/:platform", userHandler.VerifyAndGetUserId, userHandler.DeleteMarketplaceCredential)

	v1ProductGroup := apiV1.Group("product")
	v1ProductGroup.GET("/:productId",productHandler.GetProductById)
	v1ProductGroup.Use(userHandler.VerifyAndGetUserId)
	v1ProductGroup.POST("", productHandler.AddProduct)
	v1ProductGroup.GET("", productHandler.GetProducts)
	v1ProductGroup.GET("/:productId/offer", productHandler.GetOffers)
	v1ProductGroup.DELETE("/:productId", productHandler.DeleteProduct)

	v1CampaignGroup := apiV1.Group("campaign")
	v1CampaignGroup.GET("/available",campaignHandler.GetPublicCampaigns)
	v1CampaignGroup.Use(userHandler.VerifyAndGetUserId)
	v1CampaignGroup.POST("", campaignHandler.CreateCampaign)
	v1CampaignGroup.GET("", campaignHandler.GetCampaigns)
	v1CampaignGroup.DELETE("/:campaign_id", campaignHandler.DeleteCampaign)

	v1LinkGroup := apiV1.Group("link")
	v1LinkGroup.POST("", userHandler.VerifyAndGetUserId, linkHandler.CreateLink)
	v1LinkGroup.GET("/campaign/:campaignId", linkHandler.GetLinksByCampaign)
	v1LinkGroup.DELETE("/:link_id", userHandler.VerifyAndGetUserId, linkHandler.DeleteLink)
	v1LinkGroup.GET("/:link_id", linkHandler.GetLinkById)
	v1LinkGroup.GET("/short-code/:short_code", linkHandler.GetLinkByShortCode)
	v1LinkGroup.GET("/redirect/:short_code", linkHandler.RedirectLink)

	v1DashboardGroup := apiV1.Group("dashboard")
	v1DashboardGroup.GET("/metrics", userHandler.VerifyAndGetUserId, dashboardHandler.GetDashboardData)

	return g
}
