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
