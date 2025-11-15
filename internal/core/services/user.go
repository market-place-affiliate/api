package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/market-place-affiliate/api/internal/core/domains"
	"github.com/market-place-affiliate/api/internal/core/dto"
	"github.com/market-place-affiliate/api/internal/core/ports"
	jwtPkg "github.com/market-place-affiliate/api/pkg/jwt"
	"github.com/market-place-affiliate/api/pkg/password"
)

type userService struct {
	passwordSalt   string
	jwtSalt        string
	userRepo       ports.UserRepository
	marketCredRepo ports.MarketplaceRepository
}

func NewUserService(passwordSalt, jwtSalt string, userRepo ports.UserRepository, marketCredRepo ports.MarketplaceRepository) ports.UserService {
	return &userService{passwordSalt: passwordSalt, jwtSalt: jwtSalt, userRepo: userRepo, marketCredRepo: marketCredRepo}
}

func (s *userService) Register(ctx context.Context, pwd, email string) (dto.Response[string], error) {
	_, err := s.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return dto.Response[string]{
			HttpCode: http.StatusBadRequest,
			Success:  false,
			Code:     1001,
			Message:  "User already exists",
		}, nil
	}
	hashedPassword := password.HashPassword(pwd, []byte(s.passwordSalt))
	newUser, err := s.userRepo.CreateUser(ctx, domains.User{
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return dto.Response[string]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     1002,
			Message:  "Failed to create user",
		}, err
	}
	token := jwtPkg.GenerateToken(s.jwtSalt, fmt.Sprint(newUser.Id), "", "web", 60*24)
	return dto.Response[string]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "User registered successfully",
		Data:     token,
	}, nil
}

func (s *userService) Login(ctx context.Context, pwd, email string) (dto.Response[string], error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return dto.Response[string]{
			HttpCode: http.StatusBadRequest,
			Success:  false,
			Code:     1003,
			Message:  "User not found",
		}, nil
	}
	if !password.VerifyPassword(pwd, s.passwordSalt, user.Password) {
		return dto.Response[string]{
			HttpCode: http.StatusUnauthorized,
			Success:  false,
			Code:     1004,
			Message:  "Invalid credentials",
		}, nil
	}
	token := jwtPkg.GenerateToken(s.jwtSalt, fmt.Sprint(user.Id), "", "web", 60*24)
	return dto.Response[string]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "User logged in successfully",
		Data:     token,
	}, nil
}

func (s *userService) GetMe(ctx context.Context, userId int64) (dto.Response[domains.User], error) {
	user, err := s.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		return dto.Response[domains.User]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     1008,
			Message:  "Failed to get user",
		}, err
	}
	user.Password = ""
	return dto.Response[domains.User]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "User retrieved successfully",
		Data:     user,
	}, nil
}

func (s *userService) VerifyAndGetUserId(token string) (int64, error) {
	claims, ok := jwtPkg.ValidAndGetClaims(s.jwtSalt, token)
	if !ok {
		return 0, fmt.Errorf("invalid token")
	}
	var userId int64
	_, err := fmt.Sscan(claims.Userid, &userId)
	if err != nil {
		return 0, fmt.Errorf("invalid user id in token")
	}
	return userId, nil
}

func (s *userService) SaveMarketplaceCredential(ctx context.Context, userId int64, cred dto.MarketplaceCredentialRequest) (dto.Response[string], error) {
	_, err := s.marketCredRepo.Save(ctx, domains.MarketplaceCredential{
		UserId:      userId,
		Marketplace: cred.Platform,
		AppId:       cred.AppId,
		AppSecret:   cred.AppSecret,
		AppKey:      cred.AppKey,
		UserToken:   cred.UserToken,
	})
	if err != nil {
		return dto.Response[string]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     1005,
			Message:  "Failed to save marketplace credential",
		}, err
	}
	return dto.Response[string]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Marketplace credential saved successfully",
	}, nil
}

func (s *userService) CheckMarketplaceCredential(ctx context.Context, userId int64, platform string) (dto.Response[bool], error) {
	_, err := s.marketCredRepo.GetByUserIdAndPlatform(ctx, userId, platform)
	if err != nil {
		return dto.Response[bool]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     1007,
			Message:  "Marketplace credential not found",
		}, err
	}
	return dto.Response[bool]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Marketplace credential exists",
		Data:     true,
	}, nil
}

func (s *userService) DeleteMarketplaceCredential(ctx context.Context, userId int64, platform string) (dto.Response[string], error) {
	err := s.marketCredRepo.DeleteByUserIdAndPlatform(ctx, userId, platform)
	if err != nil {
		return dto.Response[string]{
			HttpCode: http.StatusInternalServerError,
			Success:  false,
			Code:     1006,
			Message:  "Failed to delete marketplace credential",
		}, err
	}
	return dto.Response[string]{
		HttpCode: http.StatusOK,
		Success:  true,
		Code:     0,
		Message:  "Marketplace credential deleted successfully",
	}, nil
}
