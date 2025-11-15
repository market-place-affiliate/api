package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
)

type ProductHandler struct {
	productService ports.ProductService
}

func NewProductHandler(productService ports.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// AddProduct godoc
// @Summary Add a new product
// @Description Create a new affiliate product from marketplace URL
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.CreateProductRequest true "Product request"
// @Success 200 {object} dto.ProductsResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /product [post]
func (h *ProductHandler) AddProduct(g *gin.Context) {
	ctx := g.Request.Context()
	body := dto.CreateProductRequest{}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userId := g.GetInt64("userId")
	res, err := h.productService.CreateProduct(ctx, userId, body)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

// GetOffers godoc
// @Summary Get product offers
// @Description Get marketplace offers for a specific product
// @Tags product
// @Produce json
// @Security BearerAuth
// @Param productId path string true "Product ID"
// @Success 200 {object} dto.OfferResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {object} dto.EmptyResponse "Forbidden"
// @Router /product/{productId}/offer [get]
func (h *ProductHandler) GetOffers(g *gin.Context) {
	ctx := g.Request.Context()
	productId := g.Param("productId")
	userId := g.GetInt64("userId")
	res, err := h.productService.GetOffer(ctx, userId, productId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

// GetProducts godoc
// @Summary Get user products
// @Description Get all products for the authenticated user
// @Tags product
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.ProductsResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /product [get]
func (h *ProductHandler) GetProducts(g *gin.Context) {
	ctx := g.Request.Context()
	userId := g.GetInt64("userId")
	res, err := h.productService.GetProductsByUserId(ctx, userId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product and all associated links and clicks
// @Tags product
// @Produce json
// @Security BearerAuth
// @Param productId path string true "Product ID"
// @Success 200 {object} dto.EmptyResponse
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {object} dto.EmptyResponse "Forbidden"
// @Router /product/{productId} [delete]
func (h *ProductHandler) DeleteProduct(g *gin.Context) {
	ctx := g.Request.Context()
	productId := g.Param("productId")
	userId := g.GetInt64("userId")
	res, err := h.productService.DeleteProductById(ctx, userId, productId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}

// GetProductById godoc
// @Summary Get product by ID
// @Description Get a specific product by its ID
// @Tags product
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 500 {object} dto.EmptyResponse
// @Router /product/{productId} [get]
func (h *ProductHandler) GetProductById(g *gin.Context) {
	ctx := g.Request.Context()
	productId := g.Param("productId")
	res, err := h.productService.GetProductById(ctx, productId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(http.StatusOK, res)
}
