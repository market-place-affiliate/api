package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(g *gin.Context) {
	ctx := g.Request.Context()
	body := dto.RegisterRequest{}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.AbortWithStatus(400)
		return
	}
	res, err := h.userService.Register(ctx, body.Password, body.Email)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.SetCookie("session", res.Data, 60*60*24, "/", "", true, false)
	g.JSON(200, res)
}

func (h *UserHandler) Login(g *gin.Context) {
	ctx := g.Request.Context()
	body := dto.LoginRequest{}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.AbortWithStatus(400)
		return
	}
	res, err := h.userService.Login(ctx, body.Password, body.Email)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.SetCookie("session", res.Data, 60*60*24, "/", "", true, false)
	g.JSON(200, res)
}

func (h *UserHandler) Logout(g *gin.Context) {
	g.SetCookie("session", "", -1, "/", "", true, false)
	g.JSON(200, gin.H{"message": "Logged out successfully"})
}

func (h *UserHandler) GetMe(g *gin.Context) {
	ctx := g.Request.Context()
	userId := g.GetInt64("userId")
	res, err := h.userService.GetMe(ctx, userId)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(200, res)
}

func (h *UserHandler) VerifyAndGetUserId(g *gin.Context) {
	token, err := g.Cookie("session")
	if err != nil {
		g.AbortWithStatus(401)
		return
	}
	userId, err := h.userService.VerifyAndGetUserId(token)
	if err != nil {
		g.AbortWithStatus(401)
		return
	}
	g.Set("userId", userId)
	g.Next()
}

func (h *UserHandler) SaveMarketplaceCredential(g *gin.Context) {
	ctx := g.Request.Context()
	body := dto.MarketplaceCredentialRequest{}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.AbortWithStatus(400)
		return
	}
	userId := g.GetInt64("userId")
	res, err := h.userService.SaveMarketplaceCredential(ctx, userId, body)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(200, res)
}

func (h *UserHandler) CheckMarketplaceCredential(g *gin.Context) {
	ctx := g.Request.Context()
	platform := g.Param("platform")
	userId := g.GetInt64("userId")
	res, err := h.userService.CheckMarketplaceCredential(ctx, userId, platform)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(200, res)
}

func (h *UserHandler) DeleteMarketplaceCredential(g *gin.Context) {
	ctx := g.Request.Context()
	platform := g.Param("platform")
	userId := g.GetInt64("userId")
	res, err := h.userService.DeleteMarketplaceCredential(ctx, userId, platform)
	if err != nil {
		g.JSON(res.HttpCode, res)
		return
	}
	g.JSON(200, res)
}
