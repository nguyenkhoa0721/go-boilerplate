package domain

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-boilerplate/pkg/infra"

	"go-boilerplate/internal/user/domain/validation"
)

type UserService struct {
	infra *infra.Infra
}

func NewUserService(infra *infra.Infra) *UserService {
	return &UserService{
		infra: infra,
	}
}

func (s *UserService) CheckHealth(ctx context.Context) validation.CheckHealthResponse {
	log.Info().Msg("CheckHealth")
	return validation.CheckHealthResponse{
		Success: true,
	}
}
