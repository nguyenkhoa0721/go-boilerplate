package domain

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-boilerplate/internal/auth/domain/interfaces"
	"go-boilerplate/internal/auth/domain/validation"
	"go-boilerplate/internal/user/data"
	"go-boilerplate/pkg/constant"
	"go-boilerplate/pkg/entity"
	"go-boilerplate/pkg/infra"
	"go-boilerplate/pkg/utils"
)

type AuthService struct {
	common   *infra.Infra
	userRepo *data.UserRepo
}

func NewAuthService(common *infra.Infra, userRepo *data.UserRepo) *AuthService {
	return &AuthService{
		common:   common,
		userRepo: userRepo,
	}
}

func (s *AuthService) Signup(dto *validation.SignupRequest) error {
	if err := s.userRepo.FindOne(&entity.User{}, &entity.User{Email: dto.Email}); err == nil {
		return &fiber.Error{
			Code:    fiber.StatusConflict,
			Message: "Email already exists",
		}
	}

	hashPassword := utils.Hash(dto.Password)
	if err := s.userRepo.Create(&entity.User{
		ID:       fmt.Sprint(s.common.Uuid.GenerateUuid(constant.UUID_USER)),
		Email:    dto.Email,
		Name:     dto.Name,
		Phone:    dto.Phone,
		Password: hashPassword,
	}); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(dto *validation.LoginRequest) (*interfaces.LoginResponse, error) {
	var existedUser *entity.User
	if err := s.userRepo.FindOne(&existedUser, &entity.User{Email: dto.Email}); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	hashPassword := utils.Hash(dto.Password)
	if hashPassword != existedUser.Password {
		return nil, fiber.ErrUnauthorized
	}

	return &interfaces.LoginResponse{
		Token:        "",
		RefreshToken: "",
	}, nil
}
