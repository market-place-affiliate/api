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

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags user
// @Accept json
// @Produce json
// @Param body body dto.RegisterRequest true "Register request"
// @Success 200 {object} dto.StringResponse
// @Failure 400 {object} dto.EmptyResponse
// @Router /user/register [post]
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

// Login godoc
// @Summary Login user
// @Description Authenticate user and return session token
// @Tags user
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login request"
// @Success 200 {object} dto.StringResponse
// @Failure 400 {object} dto.EmptyResponse
// @Failure 401 {object} dto.EmptyResponse
// @Router /user/login [post]
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

// Logout godoc
// @Summary Logout user
// @Description Clear session cookie
// @Tags user
// @Produce json
// @Success 200 {object} map[string]string
// @Router /user/logout [post]
func (h *UserHandler) Logout(g *gin.Context) {
	g.SetCookie("session", "", -1, "/", "", true, false)
	g.JSON(200, gin.H{"message": "Logged out successfully"})
}

// GetMe godoc
// @Summary Get current user
// @Description Get authenticated user information
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /user/me [get]
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

// SaveMarketplaceCredential godoc
// @Summary Save marketplace credentials
// @Description Save or update marketplace API credentials (Lazada/Shopee)
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.MarketplaceCredentialRequest true "Marketplace credential"
// @Success 200 {object} dto.EmptyResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /user/market-credential [post]
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

// CheckMarketplaceCredential godoc
// @Summary Check marketplace credentials
// @Description Check if marketplace credentials exist for a platform
// @Tags user
// @Produce json
// @Security BearerAuth
// @Param platform path string true "Platform name (lazada/shopee)"
// @Success 200 {object} dto.BoolResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /user/market-credential/{platform} [get]
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

// DeleteMarketplaceCredential godoc
// @Summary Delete marketplace credentials
// @Description Delete stored marketplace credentials for a platform
// @Tags user
// @Produce json
// @Security BearerAuth
// @Param platform path string true "Platform name (lazada/shopee)"
// @Success 200 {object} dto.EmptyResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /user/market-credential/{platform} [delete]
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
