package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/market-place-affiliate/api/cmd/httpserver"
	"github.com/market-place-affiliate/api/config"
	infrastructure "github.com/market-place-affiliate/api/infrastructures"
	"github.com/market-place-affiliate/api/internal/core/services"
	"github.com/market-place-affiliate/api/internal/handlers"
	"github.com/market-place-affiliate/api/internal/repositories/db"
	"github.com/market-place-affiliate/commonlib/lazada"
	"github.com/market-place-affiliate/commonlib/shopee"
)

func main() {
	cfg := config.Init()
	postgresClient := infrastructure.NewPostgresDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.DbName)
	err := db.Migrate(postgresClient)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	userRepository := db.NewUserRepository(postgresClient)
	productRepository := db.NewProductRepository(postgresClient)
	campaignRepository := db.NewCampaignRepository(postgresClient)
	linkRepository := db.NewLinkRepository(postgresClient)
	offerRepository := db.NewOfferRepository(postgresClient)
	marketplaceCredentialRepository := db.NewMarketplaceCredentialRepository(postgresClient)
	clickRepository := db.NewClickRepository(postgresClient)

	lazadaRepository := lazada.NewLazadaRepository(lazada.ApiGatewayTH,true)
	shopeeRepository := shopee.NewShopeeRepository(true)

	userService := services.NewUserService(string(cfg.Secret.PasswordSecret), string(cfg.Secret.JWTSecret), userRepository, marketplaceCredentialRepository)
	productService := services.NewProductService(productRepository, offerRepository, lazadaRepository, shopeeRepository, marketplaceCredentialRepository,linkRepository,clickRepository)
	campaignService := services.NewCampaignService(campaignRepository,linkRepository,clickRepository)
	linkService := services.NewLinkService(linkRepository, clickRepository, productRepository, campaignRepository, offerRepository, lazadaRepository, shopeeRepository, marketplaceCredentialRepository)
	dashboardService := services.NewDashboardService(clickRepository, productRepository)

	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	campaignHandler := handlers.NewCampaignHandler(campaignService)
	linkHandler := handlers.NewLinkHandler(linkService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	httpServer := httpserver.NewHttpServer(
		userHandler,
		productHandler,
		campaignHandler,
		linkHandler,
		dashboardHandler,
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
		Handler: httpServer,
	}
	log.Printf("Starting server at %s:%d\n", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
}
